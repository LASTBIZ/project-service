package data

type Investor struct {
	ID        uint64 `gorm:"primaryKey"`
	Money     uint64
	FullName  string
	UserID    uint32
	ProjectID *uint64
}
