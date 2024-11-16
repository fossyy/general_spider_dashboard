package models

type Config struct {
	ID          uint   `gorm:"primaryKey"`
	Domain      string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:text"`
	Data        []byte `gorm:"type:bytea"`
}
