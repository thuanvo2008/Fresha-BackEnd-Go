package staffmodel

import (
	"DemoProject/common"
	"errors"
	"strings"
)

const EntityName = "staffs"

type Staff struct {
	common.SQLModel `json:",inline"`
	StoreId         int           `json:"store_id" gorm:"column:store_id"`
	Firstname       string        `json:"first_name" gorm:"column:first_name"`
	Lastname        string        `json:"last_name" gorm:"column:last_name"`
	Password        string        `json:"password" gorm:"column:password;"`
	Salt            string        `json:"-" gorm:"column:salt;"`
	Avatar          *common.Image `json:"avatar" gorm:"column:avatar"`
	Role            string        `json:"role" gorm:"column:role"`
	Phone           string        `json:"phone" gorm:"column:phone"`
	Email           string        `json:"email" gorm:"column:email"`
}

func (Staff) TableName() string {
	return "staffs"
}

func (s *Staff) Mask(isAdminOrOwner bool) {
	s.GenUID(common.DBTypeStaff)
}

type StaffCrete struct {
	common.SQLModel `json:",inline"`
	StoreId         int           `json:"store_id" gorm:"column:store_id"`
	Firstname       string        `json:"first_name" gorm:"column:first_name"`
	Lastname        string        `json:"last_name" gorm:"column:last_name"`
	Password        string        `json:"password" gorm:"column:password;"`
	Salt            string        `json:"-" gorm:"column:salt;"`
	Avatar          *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
	Role            string        `json:"-" gorm:"column:role"`
	Phone           string        `json:"-" gorm:"column:phone"`
	Email           string        `json:"email" gorm:"column:email"`
}

func (StaffCrete) TableName() string {
	return Staff{}.TableName()
}

func (s *StaffCrete) Mask(isAdminOrOwner bool) {
	s.GenUID(common.DBTypeStaff)
}

func (s *StaffCrete) Validation() error {
	s.Firstname = strings.TrimSpace(s.Firstname)
	s.Lastname = strings.TrimSpace(s.Lastname)

	if len(s.Firstname) == 0 || len(s.Lastname) == 0 {
		panic("ErrFillName")
	}

	return nil
}

type StaffUpdate struct {
	Firstname *string       `json:"first_name" gorm:"column:first_name"`
	Lastname  *string       `json:"last_name" gorm:"column:last_name"`
	Password  string        `json:"password" gorm:"column:password;"`
	Status    int           `json:"status" gorm:"column:status"`
	Avatar    *common.Image `json:"avatar" gorm:"column:avatar"`
	Role      string        `json:"role" gorm:"column:role"`
	Phone     string        `json:"phone" gorm:"column:phone"`
	Email     string        `json:"email" gorm:"column:email"`
}

func (StaffUpdate) TableName() string {
	return Staff{}.TableName()
}

var (
	ErrUsernameOrPasswordInvalid = common.NewCustomError(
		errors.New("username or password invalid"),
		"username or password invalid",
		"ErrUsernameOrPasswordInvalid",
	)

	ErrEmailExisted = common.NewCustomError(
		errors.New("email has already existed"),
		"email has already existed",
		"ErrEmailExisted",
	)
)
