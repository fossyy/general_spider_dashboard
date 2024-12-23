package db

import (
	"context"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"general_spider_controll_panel/types"
	"general_spider_controll_panel/types/models"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"net/url"
	"strings"
)

type postgresDB struct {
	*gorm.DB
}

type SSLMode string

const (
	DisableSSL SSLMode = "disable"
	EnableSSL  SSLMode = "enable"
)

//go:embed trigger.sql
var trigger []byte

func NewPostgresDB(ctx context.Context, username, password, host, port, dbName string, mode SSLMode) (types.Database, error) {
	var err error
	var count int64

	connection := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=%s TimeZone=Asia/Jakarta", host, username, password, port, mode)
	initDB, err := gorm.Open(postgres.New(postgres.Config{
		DSN: connection,
	}), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Silent),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %s", err.Error())
	}

	initDB.WithContext(ctx).Raw("SELECT count(*) FROM pg_database WHERE datname = ?", dbName).Scan(&count)
	if count <= 0 {
		if err := initDB.WithContext(ctx).Exec("CREATE DATABASE " + dbName).Error; err != nil {
			return nil, fmt.Errorf("error creating database: %s", err.Error())
		}
	}

	connection = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Jakarta", host, username, password, dbName, port, mode)
	DB, err := gorm.Open(postgres.New(postgres.Config{
		DSN: connection,
	}), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Silent),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %s", err.Error())
	}

	err = DB.WithContext(ctx).AutoMigrate(&models.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %s", err.Error())
	}
	err = DB.WithContext(ctx).AutoMigrate(&models.Proxy{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %s", err.Error())
	}
	err = DB.WithContext(ctx).AutoMigrate(&models.Schedule{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %s", err.Error())
	}
	err = DB.WithContext(ctx).AutoMigrate(&models.Timeline{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %s", err.Error())
	}
	err = DB.WithContext(ctx).AutoMigrate(&models.KafkaBroker{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %s", err.Error())
	}
	err = DB.WithContext(ctx).AutoMigrate(&models.ConfigHistory{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %s", err.Error())
	}
	err = DB.WithContext(ctx).AutoMigrate(&models.TempVersion{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %s", err.Error())
	}
	err = DB.WithContext(ctx).AutoMigrate(&models.DashboardConfig{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %s", err.Error())
	}

	err = DB.Exec(string(trigger)).Error
	return &postgresDB{DB}, nil
}

func (db *postgresDB) CreateConfig(config *models.Config) error {
	return db.Create(config).Error
}

func (db *postgresDB) IsConfigExists(configName string) (bool, error) {
	var count int64
	err := db.Model(&models.Config{}).Where(&models.Config{Name: configName}).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (db *postgresDB) UpdateConfigByName(configName string, config *models.Config) error {
	return db.Model(&models.Config{}).Where(&models.Config{Name: configName}).Updates(config).Error
}

func (db *postgresDB) GetConfigByName(configName string) (*models.Config, error) {
	var config models.Config
	err := db.Model(&models.Config{}).Where(&models.Config{Name: configName}).First(&config).Error
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func (db *postgresDB) GetConfigByID(id string) (*models.Config, error) {
	var config models.Config
	if err := db.Where("id = ?", id).First(&config).Error; err != nil {
		return nil, err
	}
	return &config, nil
}

func (db *postgresDB) GetCombinedVersion(id string) (*models.CombinedVersion, error) {
	var name string

	if err := db.Table("configs").Select("name").Where("id = ?", id).Take(&name).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("config with ID %s not found: %w", id, err)
		}
		return nil, fmt.Errorf("failed to fetch config name: %w", err)
	}

	var config models.CombinedVersion

	if err := db.Table("combined_version").Where("config_name = ?", name).Take(&config).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("combined version for config %s not found: %w", name, err)
		}
		return nil, fmt.Errorf("failed to fetch combined version: %w", err)
	}

	return &config, nil
}

func (db *postgresDB) GetDomains() ([]string, error) {
	var domains []string
	if err := db.Model(&models.Config{}).Distinct("domain").Pluck("domain", &domains).Error; err != nil {
		return nil, err
	}
	return domains, nil
}

func (db *postgresDB) GetDomainsWithSchema() ([]string, error) {
	var results []struct {
		DomainProtocol string
		Domain         string
	}
	var combinedResults []string

	err := db.Model(&models.Config{}).
		Select("DISTINCT domain_protocol, domain").
		Scan(&results).Error
	if err != nil {
		return nil, err
	}

	for _, row := range results {
		combined := row.DomainProtocol + "://" + row.Domain
		combinedResults = append(combinedResults, combined)
	}

	return combinedResults, nil
}

func (db *postgresDB) GetConfigsIDByDomain(domain string) ([]string, error) {
	var configs []models.Config
	if err := db.Where("domain = ?", domain).Find(&configs).Error; err != nil {
		return nil, err
	}
	configID := make([]string, 0, len(configs))
	for _, config := range configs {
		configID = append(configID, config.ID.String())
	}
	return configID, nil
}

func (db *postgresDB) GetConfigNameAndIDByDomain(domain string) ([]*types.ConfigDetail, error) {
	var configs []models.Config
	if err := db.Where("domain = ?", domain).Order("created_at DESC").Find(&configs).Error; err != nil {
		return nil, err
	}
	data := make([]*types.ConfigDetail, 0, len(configs))
	for _, config := range configs {
		data = append(data, &types.ConfigDetail{
			ID:   config.ID.String(),
			Name: fmt.Sprintf("%s - %s", config.Name, config.ID.String()),
		})
	}
	return data, nil
}

func (db *postgresDB) GetConfigs() ([]*models.Config, error) {
	var configs []*models.Config
	if err := db.Find(&configs).Error; err != nil {
		return nil, err
	}
	return configs, nil
}

func (db *postgresDB) GetProxies() ([]*models.Proxy, error) {
	var proxies []*models.Proxy
	if err := db.Find(&proxies).Error; err != nil {
		return nil, err
	}
	return proxies, nil
}

func (db *postgresDB) GetActiveProxies() ([]*models.Proxy, error) {
	var proxies []*models.Proxy
	if err := db.Where("status = ?", models.Online).Find(&proxies).Error; err != nil {
		return nil, err
	}
	return proxies, nil
}

func (db *postgresDB) GetProxyByID(id string) (*models.Proxy, error) {
	var proxy models.Proxy
	if err := db.Where("id = ?", id).First(&proxy).Error; err != nil {
		return nil, err
	}
	return &proxy, nil
}

func (db *postgresDB) GetProxiesByJobID(id string) ([]*models.Proxy, error) {
	var proxies []*models.Proxy
	if err := db.Where("usage = ?", id).Find(&proxies).Error; err != nil {
		return nil, err
	}
	return proxies, nil
}

func (db *postgresDB) CreateProxy(proxy *models.Proxy) (*models.Proxy, error) {
	proxy.Address = strings.ReplaceAll(proxy.Address, " ", "")
	return proxy, db.Create(proxy).Error
}

func (db *postgresDB) UpdateProxyStatus(addr string, status models.ProxyStatus) error {
	var proxy models.Proxy
	if err := db.Where("address = ?", addr).First(&proxy).Error; err != nil {
		return err
	}
	proxy.Status = status
	return db.Save(&proxy).Error
}

func (db *postgresDB) UpdateProxyAsUsed(addr string, jobid string) error {
	var proxy models.Proxy
	if err := db.Where("address = ?", addr).First(&proxy).Error; err != nil {
		return err
	}
	proxy.Status = models.Used
	proxy.Usage = jobid
	return db.Save(&proxy).Error
}

func (db *postgresDB) RemoveProxyUsedStatus(addr string) error {
	var proxy models.Proxy
	if err := db.Where("address = ?", addr).First(&proxy).Error; err != nil {
		return err
	}
	proxy.Usage = ""
	proxy.Status = models.Unchecked
	return db.Save(&proxy).Error
}

func (db *postgresDB) RemoveProxyUsedStatusByJobID(jobid string) error {
	var proxies []*models.Proxy
	if err := db.Where("usage = ? ", jobid).Find(&proxies).Error; err != nil {
		return err
	}
	for _, proxy := range proxies {
		proxy.Usage = ""
		proxy.Status = models.Online
	}
	return db.Save(&proxies).Error
}

func (db *postgresDB) RemoveProxy(addr string) error {
	var proxy models.Proxy
	if err := db.Where("address = ?", addr).First(&proxy).Error; err != nil {
		return err
	}
	return db.Delete(&proxy).Error
}

func (db *postgresDB) CreateCron(cron *models.Schedule) error {
	var proxies []*models.Proxy

	var proxyAddresses []string
	if err := json.Unmarshal(cron.ProxyAddresses, &proxyAddresses); err != nil {
		return fmt.Errorf("failed to unmarshal ProxyAddresses: %w", err)
	}
	for _, address := range proxyAddresses {
		var proxy *models.Proxy
		parse, err := url.Parse(address)
		if err != nil {
			return err
		}
		if err := db.Where("address = ?", strings.Split(parse.Host, ":")[0]).First(&proxy).Error; err != nil {
			return fmt.Errorf("proxy with address %s not found", strings.Split(parse.Host, ":")[0])
		}
		proxies = append(proxies, proxy)
	}

	cron.Proxies = proxies

	return db.Create(cron).Error
}

func (db *postgresDB) GetCrons() ([]*models.Schedule, error) {
	var crons []*models.Schedule
	if err := db.Preload("Proxies").Find(&crons).Error; err != nil {
		return nil, err
	}
	return crons, nil
}

func (db *postgresDB) GetCronByID(id string) (*models.Schedule, error) {
	var cron models.Schedule
	if err := db.Preload("Proxies").Where("id = ?", id).First(&cron).Error; err != nil {
		return nil, err
	}
	return &cron, nil
}

func (db *postgresDB) ChangeCronID(id string, newID uuid.UUID) error {
	tx := db.Begin()

	var cron *models.Schedule
	if err := tx.Preload("Proxies").Where("id = ?", id).First(&cron).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to find schedule with ID %s: %w", id, err)
	}

	cron.ID = newID

	if err := tx.Create(&cron).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to create new schedule with ID %s: %s", newID, err.Error)
	}

	if err := tx.Exec("DELETE FROM schedule_proxies WHERE schedule_id = ?", id).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete old references from schedule_proxies: %w", err)
	}

	if err := tx.Where("id = ?", id).Delete(&models.Schedule{}).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete old schedule with ID %s: %w", id, err)
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (db *postgresDB) RemoveScheduleByID(id string) error {
	tx := db.Begin()
	var cron *models.Schedule
	if err := tx.Preload("Proxies").Where("id = ?", id).First(&cron).Error; err != nil {
		return err
	}
	if err := tx.Exec("DELETE FROM schedule_proxies WHERE schedule_id = ?", id).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete schedule_proxies: %w", err)
	}
	if err := tx.Delete(&cron).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (db *postgresDB) CountScheduledSpiders(project string) (int64, error) {
	var count int64
	if err := db.Table("schedules").Where("project = ?", project).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (db *postgresDB) CreateTimeline(timeline *models.Timeline) error {
	return db.Create(timeline).Error
}

func (db *postgresDB) GetTimelineByContext(context string) ([]*models.Timeline, error) {
	var timeline []*models.Timeline
	if err := db.Where("context = ?", context).Find(&timeline).Error; err != nil {
		return nil, err
	}
	return timeline, nil
}

func (db *postgresDB) RemoveTimelineByContext(context string) error {
	tx := db.Begin()
	var timeline []*models.Timeline
	if err := tx.Where("context = ?", context).Delete(&timeline).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (db *postgresDB) GetKafkaBrokers() ([]*models.KafkaBroker, error) {
	var kafkaBrokers []*models.KafkaBroker
	if err := db.Find(&kafkaBrokers).Error; err != nil {
		return nil, err
	}
	return kafkaBrokers, nil
}

func (db *postgresDB) GetKafkaBrokersById(id string) (*models.KafkaBroker, error) {
	var kafkaBroker models.KafkaBroker
	if err := db.Where("broker_id = ?", id).First(&kafkaBroker).Error; err != nil {
		return nil, err
	}
	return &kafkaBroker, nil
}

func (db *postgresDB) GetKafkaBrokersByName(name string) ([]*models.KafkaBroker, error) {
	var kafkaBrokers []*models.KafkaBroker

	if err := db.Where("broker_name LIKE ?", "%"+name+"%").Find(&kafkaBrokers).Error; err != nil {
		return nil, err
	}
	return kafkaBrokers, nil
}

func (db *postgresDB) CreateKafkaBroker(broker *models.KafkaBroker) error {
	return db.Create(broker).Error
}

//func (db *postgresDB) GetKafkaTopics() ([]*models.KafkaTopic, error) {
//	var topics []*models.KafkaTopic
//	if err := db.Find(&topics).Error; err != nil {
//		return nil, err
//	}
//	return topics, nil
//}
//
//func (db *postgresDB) IsTopicPresent(name string) bool {
//	var count int64
//	err := db.DB.Model(&models.KafkaTopic{}).Where("topic_name = ?", name).Count(&count).Error
//	if err != nil {
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//			return false
//		}
//		return false
//	}
//	return count > 0
//}
//
//func (db *postgresDB) CreateKafkaTopic(kafkaTopic *models.KafkaTopic) error {
//	return db.Create(kafkaTopic).Error
//}
