package appointmentmodel

import (
	"DemoProject/common"
	"DemoProject/modules/service/servicemodel"
	"time"
)

const EntityName = "appointments"

type Appointment struct {
	common.SQLModel `json:",inline"`
	Service         []*servicemodel.Service `json:"service" gorm:"many2many:appointment_service;"`
	UserID          int                     `json:"user_id" gorm:"column:user_id"`
	StaffID         int                     `json:"staff_id" gorm:"column:staff_id"`
	Staff           *common.SimpleStaff     `json:"staff" gorm:"Preload:false"`
	StartTime       *time.Time              `json:"start_time" gorm:"column:start_time"`
	EndTime         *time.Time              `json:"end_time" gorm:"column:end_time"`
	TotalPrice      float32                 `json:"total_price" gorm:"column:total_price"`
}

func (Appointment) TableName() string {
	return "appointments"
}

type AppointmentCreate struct {
	common.SQLModel `json:",inline"`
	Service         []*int     `json:"service" gorm:"-"`
	UserID          int        `json:"user_id" gorm:"column:user_id"`
	StaffID         int        `json:"staff_id" gorm:"column:staff_id"`
	StartTime       *time.Time `json:"start_time" gorm:"column:start_time"`
	EndTime         *time.Time `json:"end_time" gorm:"column:end_time"`
	TotalPrice      float32    `json:"total_price" gorm:"column:total_price"`
}

func (AppointmentCreate) TableName() string {
	return Appointment{}.TableName()
}

func (*AppointmentCreate) Validation() error {

	return nil
}

type AppointmentUpdate struct {
	Service    []*servicemodel.Service `json:"service" gorm:"many2many:appointment_service;"`
	StaffID    int                     `json:"staff_id" gorm:"column:staff_id"`
	StartTime  *time.Time              `json:"start_time" gorm:"column:start_time"`
	EndTime    *time.Time              `json:"end_time" gorm:"column:end_time"`
	TotalPrice float32                 `json:"total_price" gorm:"column:total_price"`
	Status     int                     `json:"status" gorm:"column:status;"`
}

func (AppointmentUpdate) TableName() string {
	return Appointment{}.TableName()
}
