package staffmodel

import (
	"DemoProject/common"
	"strings"
)

const EntityName = "staffs"

type Staff struct {
	common.SQLModel `json:",inline"`
	Firstname       string `json:"first_name" gorm:"column:first_name"`
	Lastname        string `json:"last_name" gorm:"column:last_name"`
}

func (Staff) TableName() string {
	return "staffs"
}

type StaffCrete struct {
	common.SQLModel `json:",inline"`
	Firstname       string `json:"first_name" gorm:"column:first_name"`
	Lastname        string `json:"last_name" gorm:"column:last_name"`
}

func (StaffCrete) TableName() string {
	return Staff{}.TableName()
}

func (s *StaffCrete) Validation() error {
	s.Firstname = strings.TrimSpace(s.Firstname)
	s.Lastname = strings.TrimSpace(s.Lastname)

	if len(s.Firstname) == 0 || len(s.Lastname) == 0 {
		panic("Loi cmnr")
	}

	return nil
}

type StaffUpdate struct {
	Firstname *string `json:"first_name" gorm:"column:first_name"`
	Lastname  *string `json:"last_name" gorm:"column:last_name"`
	Status    int     `json:"status" gorm:"column:status"`
}

func (StaffUpdate) TableName() string {
	return Staff{}.TableName()
}
