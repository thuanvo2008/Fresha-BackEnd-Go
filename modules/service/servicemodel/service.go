package servicemodel

import "time"

type Service struct {
	ID          int        `json:"id" gorm:"id"`
	StoreId     int        `json:"store_id" gorm:"store_id"`
	CategoryId  int        `json:"category_id" gorm:"category_id"`
	Title       string     `json:"title" gorm:"title"`
	Description string     `json:"description" gorm:"description"`
	Status      int        `json:"status" gorm:"status"`
	CreateAt    *time.Time `json:"create_at" gorm:"create_at"`
	UpdateAt    *time.Time `json:"update_at" gorm:"update_at"`
}

func (Service) TableName() string {
	return "services"
}

type ServiceCreate struct {
	ID          int        `json:"id" gorm:"id"`
	StoreId     int        `json:"store_id" gorm:"store_id"`
	CategoryId  int        `json:"category_id" gorm:"category_id"`
	Title       string     `json:"title" gorm:"title"`
	Description string     `json:"description" gorm:"description"`
	Status      int        `json:"status" gorm:"status"`
	CreateAt    *time.Time `json:"create_at" gorm:"create_at"`
	UpdateAt    *time.Time `json:"update_at" gorm:"update_at"`
}

func (ServiceCreate) TableName() string {
	return Service{}.TableName()
}

type ServiceUpdate struct {
	CategoryId  int    `json:"category_id" gorm:"category_id"`
	Title       string `json:"title" gorm:"title"`
	Description string `json:"description" gorm:"description"`
	Status      int    `json:"status" gorm:"status"`
}

func (ServiceUpdate) TableName() string {
	return Service{}.TableName()
}
