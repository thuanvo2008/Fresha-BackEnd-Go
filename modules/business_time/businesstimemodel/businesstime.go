package businesstimemodel

import (
	"time"
)

type BusinessTime struct {
	ID        int        `json:"id" gorm:"column:id"`
	StoreID   int        `json:"store_id" gorm:"column:store_id"`
	StartTime *time.Time `json:"start_time" gorm:"column:start_time"`
	EndTime   *time.Time `json:"end_time" gorm:"column:end_time"`
	DayOfWeek string     `json:"day_of_week" gorm:"column:day_of_week;"`
	CreatedAt *time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"updated_at"`
}

func (BusinessTime) TableName() string {
	return "business_time"
}
