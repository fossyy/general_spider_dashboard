package db

import (
	"fmt"
	"general_spider_controll_panel/types"
	"general_spider_controll_panel/types/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

type postgresDB struct {
	*gorm.DB
}

type SSLMode string

const (
	DisableSSL SSLMode = "disable"
	EnableSSL  SSLMode = "enable"
)

func NewPostgresDB(username, password, host, port, dbName string, mode SSLMode) types.Database {
	var err error
	var count int64

	connection := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=%s TimeZone=Asia/Jakarta", host, username, password, port, mode)
	initDB, err := gorm.Open(postgres.New(postgres.Config{
		DSN: connection,
	}), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Silent),
	})

	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	initDB.Raw("SELECT count(*) FROM pg_database WHERE datname = ?", dbName).Scan(&count)
	if count <= 0 {
		if err := initDB.Exec("CREATE DATABASE " + dbName).Error; err != nil {
			panic("Error creating database: " + err.Error())
		}
	}

	connection = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Jakarta", host, username, password, dbName, port, mode)
	DB, err := gorm.Open(postgres.New(postgres.Config{
		DSN: connection,
	}), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Silent),
	})

	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	err = DB.AutoMigrate(&models.Config{})
	if err != nil {
		panic(err.Error())
		return nil
	}
	err = DB.AutoMigrate(&models.Proxy{})
	if err != nil {
		panic(err.Error())
		return nil
	}
	return &postgresDB{DB}
}

func (db *postgresDB) CreateConfig(config *models.Config) error {
	return db.Create(config).Error
}

func (db *postgresDB) GetConfigByID(id string) (*models.Config, error) {
	var config models.Config
	if err := db.Where("id = ?", id).First(&config).Error; err != nil {
		return nil, err
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
	if err := db.Where("domain = ?", domain).Find(&configs).Error; err != nil {
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

func (db *postgresDB) GetProxyByID(id string) (*models.Proxy, error) {
	var proxy models.Proxy
	if err := db.Where("id = ?", id).First(&proxy).Error; err != nil {
		return nil, err
	}
	return &proxy, nil
}

func (db *postgresDB) CreateProxy(proxy *models.Proxy) (*models.Proxy, error) {
	return proxy, db.Create(proxy).Error
}

func (db *postgresDB) UpdateProxyStatus(addr string, status models.ProxyStatus) error {
	fmt.Println("got : ", addr, "and :", status)
	var proxy models.Proxy
	if err := db.Where("address = ?", addr).First(&proxy).Error; err != nil {
		return err
	}
	proxy.Status = status
	return db.Save(&proxy).Error
}

func (db *postgresDB) RemoveProxy(addr string) error {
	var proxy models.Proxy
	if err := db.Where("address = ?", addr).First(&proxy).Error; err != nil {
		return err
	}
	return db.Delete(&proxy).Error
}
