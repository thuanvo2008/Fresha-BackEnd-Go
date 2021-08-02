package common

const (
	DBTypeStaff  = 1
	DBTypeUser   = 2
	DBTypeClient = 3
)

const CurrentStaff = "staff"

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}
