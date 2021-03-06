package staffbiz

import (
	"DemoProject/modules/staff/staffmodel"
	"context"
)

type CreateStaffStore interface {
	Create(ctx context.Context, data *staffmodel.StaffCrete) error
}

type createStaffBiz struct {
	store CreateStaffStore
}

func NewCreateStaffBiz(store CreateStaffStore) *createStaffBiz {
	return &createStaffBiz{store: store}
}

func (biz *createStaffBiz) CreateStaff(ctx context.Context, data *staffmodel.StaffCrete) error {
	if err := data.Validation(); err != nil {
		return err
	}

	if err := biz.store.Create(ctx, data); err != nil {
		return err
	}

	return nil
}
