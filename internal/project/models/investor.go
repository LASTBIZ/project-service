package models

import "lastbiz/project-service/pkg/project"

type Investor struct {
	ID        uint64 `gorm:"primaryKey"`
	Money     uint64
	FullName  string
	UserID    uint32
	ProjectID *uint64
}

func (r Investor) ToRPC() *project.Investor {
	return &project.Investor{
		Id:       r.ID,
		Money:    r.Money,
		FullName: r.FullName,
		UserId:   r.UserID,
	}
}
