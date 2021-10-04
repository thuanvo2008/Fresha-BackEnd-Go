package staffbiz

import (
	"DemoProject/common"
	"DemoProject/modules/staff/staffmodel"
	"context"
)

type ListStaffStore interface {
	ListDataByCondition(ctx context.Context,
		condition map[string]interface{},
		filter *staffmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]staffmodel.Staff, error)
}

type listStaffBiz struct {
	store ListStaffStore
}

func NewListStaffBiz(store ListStaffStore) *listStaffBiz {
	return &listStaffBiz{store: store}
}

func (biz *listStaffBiz) ListStaff(ctx context.Context,
	filter *staffmodel.Filter,
	paging *common.Paging,
) ([]staffmodel.Staff, error) {

	result, err := biz.store.ListDataByCondition(ctx, nil, filter, paging)

	return result, err
}
