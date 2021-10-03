package common

const EntityName = "staffs"

type SimpleStaff struct {
	SQLModel  `json:",inline"`
	Firstname string `json:"first_name" gorm:"column:first_name"`
	Lastname  string `json:"last_name" gorm:"column:last_name"`
	Avatar    *Image `json:"avatar" gorm:"column:avatar"`
	Role      string `json:"role" gorm:"column:role"`
}

func (SimpleStaff) TableName() string {
	return "staffs"
}

func (s *SimpleStaff) Mask(isAdminOrOwner bool) {
	s.GenUID(DBTypeStaff)
}
