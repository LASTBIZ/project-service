package data

import "time"

type RoadMap struct {
	ID          uint64 `gorm:"primaryKey"`
	StartDate   time.Time
	EndDate     time.Time
	Description string `gorm:"type:text"`
	Job         string `gorm:"type:text"`
	ProjectID   uint64
}
