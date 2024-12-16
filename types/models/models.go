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
	ID               uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        sql.NullTime    `gorm:"index"`
	Domain           string          `gorm:"type:varchar(255);not null" json:"domain"`
	DomainProtocol   string          `gorm:"type:varchar(10);default:'http';not null" json:"domain_protocol"`
	Type             string          `gorm:"type:varchar(12);not null" json:"type"`
	Name             string          `gorm:"type:varchar(255);not null;unique" json:"name"`
	Description      string          `gorm:"type:text" json:"description"`
	Data             json.RawMessage `gorm:"type:bytea" json:"data"`
	DashboardVersion string          `gorm:"type:varchar(10);not null;default:'unknown'" json:"dashboard_version"`
	FullVersion      string          `gorm:"type:varchar(255)" json:"full_version"`
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

type Timeline struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	Title     string         `gorm:"size:255;not null" json:"title"`
	Message   string         `gorm:"type:text;not null" json:"message"`
	Context   string         `gorm:"size:255;not null" json:"context"`
	Status    TimelineStatus `gorm:"type:varchar(10);not null" json:"status"`
}

//type KafkaTopic struct {
//	CreatedAt         time.Time `json:"created_at"`
//	UpdatedAt         time.Time `json:"updated_at"`
//	TopicID           uuid.UUID `gorm:"type:uuid;primaryKey" json:"topic_id"`
//	TopicName         string    `gorm:"size:255;not null" json:"topic_name"`
//	PartitionCount    int       `gorm:"not null" json:"partition_count"`
//	ReplicationFactor int       `gorm:"not null" json:"replication_factor"`
//	BrokerID          uuid.UUID `gorm:"type:uuid;not null" json:"broker_id"`
//}
//
//type KafkaBroker struct {
//	CreatedAt        time.Time    `json:"created_at"`
//	UpdatedAt        time.Time    `json:"updated_at"`
//	BrokerID         uuid.UUID    `gorm:"type:uuid;primaryKey" json:"broker_id"`
//	BrokerName       string       `gorm:"size:255;not null" json:"broker_name"`
//	Host             string       `gorm:"size:255;not null" json:"host"`
//	Port             string       `gorm:"size:10;not null" json:"port"`
//	SecurityProtocol string       `gorm:"size:50;not null;default:'PLAINTEXT'" json:"security_protocol"`
//	Topics           []KafkaTopic `gorm:"foreignKey:BrokerID;constraint:OnDelete:CASCADE;" json:"topics"`
//}

type KafkaTopic struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	TopicID   uuid.UUID `gorm:"type:uuid;primaryKey" json:"topic_id"`
	TopicName string    `gorm:"size:255;not null" json:"topic_name"`
}

type KafkaBroker struct {
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	BrokerID    uuid.UUID `gorm:"type:uuid;primaryKey" json:"broker_id"`
	BrokerGroup string    `gorm:"size:255;not null" json:"broker_group"`
	Host        string    `gorm:"size:255;not null" json:"host"`
	Port        string    `gorm:"size:10;not null" json:"port"`
}
