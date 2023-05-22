package models

import (
	"errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	"lastbiz/project-service/pkg/project/pb"
	"sort"
	"strings"
	"time"
)

type Project struct {
	ID          uint64 `gorm:"primaryKey"`
	Images      Image  `gorm:"embedded"`
	Dates       Date   `gorm:"embedded"`
	Budgets     Budget `gorm:"embedded"`
	Location    string
	Name        string
	Description string `gorm:"type:text"`
	Industry    string
	Investors   []Investor
	Roadmaps    []Roadmap
	CategoryID  string `gorm:"not null; unique_index"`
}

func (p Project) ToRPC() *pb.Project {
	investors := make([]*pb.Investor, 0)
	for _, investor := range p.Investors {
		investors = append(investors, investor.ToRPC())
	}
	r1 := make(RoadMapSlice, 0)
	for _, roadmap := range p.Roadmaps {
		r1 = append(r1, roadmap)
	}
	sort.Sort(r1)
	roadmaps := make([]*pb.RoadMap, 0)

	for _, roadmap := range r1 {
		roadmaps = append(roadmaps, roadmap.ToRPC())
	}

	return &pb.Project{
		Id:          p.ID,
		Images:      p.Images.ToRPC(),
		Dates:       p.Dates.ToRPC(),
		Budgets:     p.Budgets.ToRPC(),
		Location:    p.Location,
		Name:        p.Name,
		Description: p.Description,
		Industry:    p.Industry,
		Investors:   investors,
		Roadmaps:    roadmaps,
	}
}

func (p Project) Validate() error {
	//if p.ID == 0 {
	//	return errors.New("id is required")
	//}
	if strings.TrimSpace(p.Name) == "" {
		return errors.New("name is required")
	}
	if strings.TrimSpace(p.Industry) == "" {
		return errors.New("industry is required")
	}
	if strings.TrimSpace(p.Description) == "" {
		return errors.New("description is required")
	}
	if strings.TrimSpace(p.Location) == "" {
		return errors.New("location is required")
	}
	if p.Dates.StartDate.IsZero() {
		return errors.New("start date is required")
	}
	if p.Dates.EndDate.IsZero() {
		return errors.New("end date is required")
	}
	if strings.TrimSpace(p.Images.MainImageUrl) == "" {
		return errors.New("main image is required")
	}
	if strings.TrimSpace(p.Images.RoadMapImgUrl) == "" {
		return errors.New("road map image is required")
	}
	if p.Budgets.Need == 0 {
		return errors.New("budget is required")
	}
	return nil
}

type Image struct {
	RoadMapImgUrl string
	MainImageUrl  string
}

func (i Image) ToRPC() *pb.Images {
	return &pb.Images{
		RoadmapImgUrl: i.RoadMapImgUrl,
		MainImageUrl:  i.MainImageUrl,
	}
}

type Date struct {
	StartDate time.Time
	EndDate   time.Time
}

func (d Date) ToRPC() *pb.Date {
	return &pb.Date{
		StartDate: timestamppb.New(d.StartDate),
		EndDate:   timestamppb.New(d.EndDate),
	}
}

type Budget struct {
	Current uint64
	Need    uint64
}

func (b Budget) ToRPC() *pb.Budget {
	return &pb.Budget{
		Current: b.Current,
		Need:    b.Need,
	}
}
