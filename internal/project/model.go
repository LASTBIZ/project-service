package project

import "time"

type Image struct {
	RoadMapImgUrl string
	MainImageUrl  string
}

type Date struct {
	StartDate time.Time
	EndDate   time.Time
}

type Budget struct {
	Current uint64
	Need    uint64
}

type Project struct {
	ID             uint64 `gorm:"primaryKey"`
	Images         Image  `gorm:"embedded"`
	Dates          Date   `gorm:"embedded"`
	Budgets        Budget `gorm:"embedded"`
	Location       string
	Name           string
	Description    string `gorm:"type:text"`
	Industry       string
	Investors      []Investor
	Roadmaps       []Roadmap
	OrganizationID uint32
}

type Roadmap struct {
	ID          uint64 `gorm:"primaryKey"`
	Dates       Date   `gorm:"embedded"`
	Name        string
	Description string `gorm:"type:text"`
	Job         string `gorm:"type:text"`
	Target      string `gorm:"type:text"`
	ProjectID   uint64
}

type Investor struct {
	ID             uint64 `gorm:"primaryKey"`
	Money          uint64
	FullName       string
	UserID         uint32
	OrganizationID uint32
	ProjectID      uint64
}
