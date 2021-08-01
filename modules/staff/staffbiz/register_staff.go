package staffbiz

import (
	"DemoProject/common"
	"DemoProject/modules/staff/staffmodel"
	"context"
)

type RegisterStaffStore interface {
	FindDataByCondition(ctx context.Context, condition map[string]interface{}, monkeys ...string) (*staffmodel.Staff, error)
	Create(ctx context.Context, data *staffmodel.StaffCrete) error
}

type Hasher interface {
	Hash(data string) string
}

type registerBusiness struct {
	registerStaffStore RegisterStaffStore
	hasher             Hasher
}

func NewRegisterBusiness(store RegisterStaffStore, hasher Hasher) (resBiz *registerBusiness) {
	return &registerBusiness{registerStaffStore: store, hasher: hasher}
}

func (resBiz *registerBusiness) Register(ctx context.Context, staff *staffmodel.StaffCrete) error {
	data, err := resBiz.registerStaffStore.FindDataByCondition(ctx, map[string]interface{}{"email": staff.Email})
	if err == common.RecordNotFound {
		return err
	}

	if data != nil {
		return staffmodel.ErrEmailExisted
	}

	salt := common.GenSlat(50)
	staff.Password = resBiz.hasher.Hash(staff.Password + salt)
	staff.Salt = salt
	staff.Role = "staff"

	if err := resBiz.registerStaffStore.Create(ctx, staff); err != nil {
		return common.ErrCannotCreateEntity(staff.TableName(), err)
	}

	return nil
}
