package staffbiz

import (
	"DemoProject/modules/staff/staffmodel"
	"context"
	"errors"
)

type DeleteStaffStore interface {
	FindDataByCondition(ctx context.Context,
		condition map[string]interface{},
		morekeys ...string) (*staffmodel.Staff, error)

	DeleteData(ctx context.Context, id int) error
}

type deleteStaffBiz struct {
	store DeleteStaffStore
}

func NewDeleteStaffBiz(store DeleteStaffStore) *deleteStaffBiz {
	return &deleteStaffBiz{store: store}
}

func (biz *deleteStaffBiz) DeleteStaffBiz(ctx context.Context, id int) error {

	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return errors.New("Data deleted")
	}

	if err := biz.store.DeleteData(ctx, id); err != nil {
		return err
	}

	return err
}
