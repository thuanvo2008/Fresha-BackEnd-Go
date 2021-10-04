package staffworkingbiz

import (
	"DemoProject/modules/staff_working/staffworkingmodel"
	"context"
	"time"
)

type FindStaffWorkingStorage interface {
	Find(ctx context.Context, staffID int, startTime,
		endTime *time.Time) (*staffworkingmodel.StaffWorking, error)
}

type findStaffWorkingBiz struct {
	store FindStaffWorkingStorage
}

func NewFindStaffWorkingBiz(store FindStaffWorkingStorage) *findStaffWorkingBiz {
	return &findStaffWorkingBiz{store: store}
}

func (biz *findStaffWorkingBiz) FindStaffWorking(ctx context.Context, storeID int, startTime,
	endTime *time.Time) (*staffworkingmodel.StaffWorking, error) {

	data, err := biz.store.Find(ctx, storeID, startTime, endTime)
	if err != nil {
		return nil, err
	}

	return data, nil
}
