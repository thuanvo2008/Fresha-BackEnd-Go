package staffworkingbiz

import (
	"DemoProject/modules/staff_working/staffworkingmodel"
	"context"
	"errors"
	"time"
)

type CreateStaffWorkingStorage interface {
	Create(ctx context.Context, data *staffworkingmodel.StaffWorking) error
	Find(ctx context.Context, staffID int, startTime,
		endTime *time.Time) (*staffworkingmodel.StaffWorking, error)
}

type createStaffWorkingBiz struct {
	store CreateStaffWorkingStorage
}

func NewCreateStaffWorkingBiz(store CreateStaffWorkingStorage) *createStaffWorkingBiz {
	return &createStaffWorkingBiz{store: store}
}

func (biz *createStaffWorkingBiz) CreateStaffWorking(ctx context.Context, data *staffworkingmodel.StaffWorking) error {
	oldData, err := biz.store.Find(ctx, data.StaffID, data.StartTime, data.EndTime)
	if err != nil {
		return err
	}

	if oldData != nil {
		return errors.New("Staff working is duplicated")
	}

	if err := biz.store.Create(ctx, data); err != nil {
		return err
	}

	return nil
}
