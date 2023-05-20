package project

import (
	"errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	"lastbiz/project-service/pkg/project"
	"sort"
	"strings"
	"time"
)

type Image struct {
	RoadMapImgUrl string
	MainImageUrl  string
}

func (i Image) ToRPC() *project.Images {
	return &project.Images{
		RoadmapImgUrl: i.RoadMapImgUrl,
		MainImageUrl:  i.MainImageUrl,
	}
}

type Date struct {
	StartDate time.Time
	EndDate   time.Time
}

func (d Date) ToRPC() *project.Date {
	return &project.Date{
		StartDate: timestamppb.New(d.StartDate),
		EndDate:   timestamppb.New(d.EndDate),
	}
}

type Budget struct {
	Current uint64
	Need    uint64
}

func (b Budget) ToRPC() *project.Budget {
	return &project.Budget{
		Current: b.Current,
		Need:    b.Need,
	}
}

type RoadMapSlice []Roadmap

func (p RoadMapSlice) Len() int {
	return len(p)
}

func (p RoadMapSlice) Less(i, j int) bool {
	return p[i].Dates.StartDate.Before(p[j].Dates.StartDate)
}

func (p RoadMapSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

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
}

func (p Project) ToRPC() *project.Project {
	investors := make([]*project.Investor, 0)
	for _, investor := range p.Investors {
		investors = append(investors, investor.ToRPC())
	}
	r1 := make(RoadMapSlice, 0)
	for _, roadmap := range p.Roadmaps {
		r1 = append(r1, roadmap)
	}
	sort.Sort(r1)
	roadmaps := make([]*project.RoadMap, 0)

	for _, roadmap := range r1 {
		roadmaps = append(roadmaps, roadmap.ToRPC())
	}

	return &project.Project{
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

type Roadmap struct {
	ID          uint64 `gorm:"primaryKey"`
	Dates       Date   `gorm:"embedded"`
	Name        string
	Description string `gorm:"type:text"`
	Job         string `gorm:"type:text"`
	Target      string `gorm:"type:text"`
	ProjectID   uint64
}

func (r Roadmap) Validate() error {
	if r.Dates.StartDate.IsZero() {
		return errors.New("start date is required")
	}
	if r.Dates.EndDate.IsZero() {
		return errors.New("end date is required")
	}
	if strings.TrimSpace(r.Name) == "" {
		return errors.New("name is required")
	}
	if strings.TrimSpace(r.Description) == "" {
		return errors.New("description is required")
	}
	if strings.TrimSpace(r.Job) == "" {
		return errors.New("job is required")
	}
	if strings.TrimSpace(r.Target) == "" {
		return errors.New("target is required")
	}
	return nil
}

func (r Roadmap) ToRPC() *project.RoadMap {
	return &project.RoadMap{
		Id:          r.ID,
		Dates:       r.Dates.ToRPC(),
		Name:        r.Name,
		Description: r.Description,
		Job:         r.Job,
		Target:      r.Target,
	}
}

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

type GRPCRoadmap struct {
	*project.RoadMap
}

func (r *GRPCRoadmap) ToRoadmap() Roadmap {
	return Roadmap{
		ID: r.GetId(),
		Dates: Date{
			StartDate: r.Dates.GetStartDate().AsTime(),
			EndDate:   r.Dates.GetEndDate().AsTime(),
		},
		Name:        r.Name,
		Description: r.Description,
		Job:         r.Job,
		Target:      r.Target,
	}
}
