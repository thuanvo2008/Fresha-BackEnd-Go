package closeddatemodel

import "time"

type ClosedDate struct {
	ID            int        `json:"id" gorm:"column:id"`
	StoreID       int        `json:"store_id" gorm:"column:store_id"`
	StartDateTime *time.Time `json:"start_date_time" gorm:"column:start_date_time"`
	EndDateTime   *time.Time `json:"end_date_time" gorm:"end_date_time"`
	CreatedAt     *time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at" gorm:"updated_at"`
}

func (ClosedDate) TableName() string {
	return "closed_date"
}
