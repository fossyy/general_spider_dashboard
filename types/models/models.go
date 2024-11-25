package models

import (
	"encoding/json"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Config struct {
	ID          uuid.UUID       `gorm:"type:uuid;primaryKey" json:"id"`
	Domain      string          `gorm:"type:varchar(255);not null" json:"domain"`
	Type        string          `gorm:"type:varchar(12);not null" json:"type"`
	Name        string          `gorm:"type:varchar(255);not null;unique" json:"name"`
	Description string          `gorm:"type:text" json:"description"`
	Data        json.RawMessage `gorm:"type:bytea" json:"data"`
}

type ProxyStatus string

const (
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
}
