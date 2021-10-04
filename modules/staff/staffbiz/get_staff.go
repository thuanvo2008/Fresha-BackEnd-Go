package staffbiz

import (
	"DemoProject/common"
	"DemoProject/modules/staff/staffmodel"
	"context"
)

type GetStaffStore interface {
	FindDataByCondition(ctx context.Context,
		condition map[string]interface{},
		morekeys ...string) (*staffmodel.Staff, error)
}

type getStaffBiz struct {
	store GetStaffStore
}

func NewGetStaffBiz(store GetStaffStore) *getStaffBiz {
	return &getStaffBiz{store: store}
}

func (biz *getStaffBiz) GetStaff(ctx context.Context,
	id int) (*staffmodel.Staff, error) {

	data, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err != common.RecordNotFound {
			return nil, common.ErrCannotGetEntity(staffmodel.EntityName, err)
		}

		return nil, common.ErrCannotGetEntity(staffmodel.EntityName, err)
	}

	if data.Status == 0 {
		return nil, common.ErrEntityDeleted(staffmodel.EntityName, nil)
	}

	return data, err
}
