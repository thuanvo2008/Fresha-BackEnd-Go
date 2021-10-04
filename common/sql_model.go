package common

import (
	"time"
)

type SQLModel struct {
	Id        int        `json:"id" gorm:"column:id;"`
	FakeID    *UID       `json:"-" gorm:"-"`
	Status    int        `json:"status" gorm:"column:status;default:1;"`
	CreatedAt *time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"updated_at"`
}

func (s *SQLModel) GenUID(DBType int) {
	uid := NewUID(uint32(s.Id), DBType, 1)
	s.FakeID = &uid
}
