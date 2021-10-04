package staffbiz

import (
	"DemoProject/modules/staff/staffmodel"
	"context"
	"errors"
)

type UpdateStaffStore interface {
	FindDataByCondition(ctx context.Context,
		condition map[string]interface{},
		morekeys ...string) (*staffmodel.Staff, error)

	UpdateData(ctx context.Context,
		id int,
		data *staffmodel.StaffUpdate,
	) error
}

type updateStaffBiz struct {
	store UpdateStaffStore
}

func NewUpdateStaffBiz(store UpdateStaffStore) *updateStaffBiz {
	return &updateStaffBiz{store: store}
}

func (biz *updateStaffBiz) UpdateStaff(ctx context.Context,
	id int,
	data *staffmodel.StaffUpdate) error {

	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return errors.New("Data deleted")
	}

	if err := biz.store.UpdateData(ctx, id, data); err != nil {
		return err
	}

	return err
}
