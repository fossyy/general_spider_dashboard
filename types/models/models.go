package models

import (
	"database/sql"
	"encoding/json"
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type Config struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      sql.NullTime    `gorm:"index"`
	Domain         string          `gorm:"type:varchar(255);not null" json:"domain"`
	DomainProtocol string          `gorm:"type:varchar(10);default:'http';not null" json:"domain_protocol"`
	Type           string          `gorm:"type:varchar(12);not null" json:"type"`
	Name           string          `gorm:"type:varchar(255);not null" json:"name"`
	Description    string          `gorm:"type:text" json:"description"`
	ConfigVersion  int             `gorm:"default:1" json:"config_version"`
	Data           json.RawMessage `gorm:"type:bytea" json:"data"`
}

type DashboardConfig struct {
	ID               uint   `gorm:"primaryKey"`
	ConfigName       string `gorm:"type:text;not null;unique" json:"config_name"`
	DashboardVersion int    `gorm:"default:1" json:"dashboard_version"`
}

type ConfigHistory struct {
	ID        uuid.UUID       `gorm:"type:uuid;primaryKey" json:"id"`
	ConfigsID uuid.UUID       `gorm:"type:uuid;not null" json:"configs_id"`
	BaseURL   string          `gorm:"type:varchar(255);not null" json:"base_url"`
	Version   string          `gorm:"type:varchar(50);not null" json:"version"`
	Data      json.RawMessage `gorm:"type:bytea" json:"data"`
	Config    Config          `gorm:"foreignKey:ConfigsID;constraint:OnDelete:CASCADE;" json:"config"` // Foreign key to Config
}

type TempVersion struct {
	ID          uuid.UUID       `gorm:"type:uuid;primaryKey;unique" json:"id"`
	BaseURL     string          `gorm:"type:varchar(255);not null" json:"base_url"`
	DashboardID uint            `gorm:"not null" json:"dashboard_id"`
	ConfigsID   uuid.UUID       `gorm:"type:uuid;not null" json:"configs_id"`
	Version     string          `gorm:"type:varchar(50);not null" json:"version"`
	Dashboard   DashboardConfig `gorm:"foreignKey:DashboardID;constraint:OnDelete:CASCADE;" json:"dashboard"` // Foreign key to DashboardConfig
	Config      Config          `gorm:"foreignKey:ConfigsID;constraint:OnDelete:CASCADE;" json:"config"`      // Foreign key to Config
}

type ProxyStatus string

const (
	Used      ProxyStatus = "Used"
	Online    ProxyStatus = "Online"
	Offline   ProxyStatus = "Offline"
	Checking  ProxyStatus = "Checking"
	Unchecked ProxyStatus = "Unchecked"
)

type Proxy struct {
	gorm.Model
	Address  string      `gorm:"size:255;not null" json:"address"`
	Port     string      `gorm:"size:10;not null" json:"port"`
	Protocol string      `gorm:"size:50;not null" json:"protocol"`
	Status   ProxyStatus `gorm:"type:varchar(50);not null;default:'Unchecked'" json:"status"`
	Usage    string      `gorm:"size:255" json:"usage"`
}

type Schedule struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      sql.NullTime   `gorm:"index"`
	Schedule       string         `gorm:"size:255;not null" json:"schedule"`
	Project        string         `gorm:"size:255;not null" json:"project"`
	Spider         string         `gorm:"size:255;not null" json:"spider"`
	ConfigID       string         `gorm:"size:255;not null" json:"config_id"`
	OutputDST      string         `gorm:"size:255;not null" json:"output_dst"`
	JobID          string         `gorm:"size:255;not null" json:"job_id"`
	AdditionalArgs datatypes.JSON `gorm:"type:text" json:"additional_args"`
	ProxyAddresses datatypes.JSON `gorm:"type:json" json:"proxy_addresses"`
	Proxies        []*Proxy       `gorm:"many2many:schedule_proxies;" json:"proxies"`
}

type TimelineStatus string

const (
	Success TimelineStatus = "Success"
	Failed  TimelineStatus = "Failed"
)

type KafkaBroker struct {
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	BrokerID         uuid.UUID `gorm:"type:uuid;primaryKey" json:"broker_id"`
	BrokerGroup      string    `gorm:"size:255;not null" json:"broker_group"`
	Host             string    `gorm:"size:255;not null" json:"host"`
	Port             string    `gorm:"size:10;not null" json:"port"`
	SecurityProtocol string    `gorm:"size:50;not null;default:'PLAINTEXT'" json:"security_protocol"`
}

type Timeline struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	Title     string         `gorm:"size:255;not null" json:"title"`
	Message   string         `gorm:"type:text;not null" json:"message"`
	Context   string         `gorm:"size:255;not null" json:"context"`
	Status    TimelineStatus `gorm:"type:varchar(10);not null" json:"status"`
}

type CombinedVersion struct {
	DashboardID      string `gorm:"primaryKey" json:"dashboard_id"`
	ConfigName       string `gorm:"size:255;not null" json:"config_name"`
	DashboardVersion string `gorm:"size:255;not null" json:"dashboard_version"`
	ConfigVersion    string `gorm:"size:255;not null" json:"config_version"`
	FullVersion      string `gorm:"size:255;not null" json:"full_version"`
}
