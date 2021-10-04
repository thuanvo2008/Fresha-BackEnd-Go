package staffworkingmodel

import "time"

type StaffWorking struct {
	ID        int        `json:"id" gorm:"column:id"`
	StaffID   int        `json:"staff_id" gorm:"column:staff_id"`
	StartTime *time.Time `json:"start_time" gorm:"column:start_time"`
	EndTime   *time.Time `json:"end_time" gorm:"end_time"`
	CreatedAt *time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"updated_at"`
}

func (StaffWorking) TableName() string {
	return "staff_working"
}
