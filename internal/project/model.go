package project

import "time"

type Project struct {
	ID             uint64
	Location       string
	OrganizationID uint32
	RoadMapImgUrl  string
	Name           string
	MainImage      string
	StartDate      time.Time
	EndDate        time.Time
	Industry       string
	Budget         uint64
	EndBudget      uint64
}

type Investor struct {
	ID     uint64
	Money  uint64
	UserID uint32
}
