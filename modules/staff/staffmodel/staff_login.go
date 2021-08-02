package staffmodel

type Login struct {
	Email    string `json:"email" form:"email" gorm:"column:email"`
	Password string `json:"password" form:"password" gorm:"column:password"`
}

func (l *Login) TableName() string {
	return "staffs"
}
