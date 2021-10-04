package appointmentbiz

import (
	"DemoProject/modules/booking/appointmentmodel"
	"DemoProject/modules/business_time/businesstimemodel"
	"DemoProject/modules/closed_date/closeddatemodel"
	"DemoProject/modules/staff_working/staffworkingmodel"
	"context"
	errors "errors"
	"gorm.io/gorm"
	"time"
)

type CreateAppointmentStorage interface {
	Create(ctx context.Context, data *appointmentmodel.AppointmentCreate) error
	InsertMiddleTable(ctx context.Context, ids []*int, appointmentId int) error
	FindDataByDuration(ctx context.Context,
		staffId int,
		startTime *time.Time,
		endTime *time.Time) (*appointmentmodel.Appointment, error)
}

type FindBusinessTimeStorage interface {
	Find(ctx context.Context, storeID int, dayOfWeek string) (*businesstimemodel.BusinessTime, error)
}

type FindClosedDateStorage interface {
	Find(ctx context.Context, storeID int, startDateTime,
		endDateTime *time.Time) (*closeddatemodel.ClosedDate, error)
}

type FindStaffWorkingStorage interface {
	Find(ctx context.Context, staffID int, startTime,
		endTime *time.Time) (*staffworkingmodel.StaffWorking, error)
}

type createAppointmentBiz struct {
	store             CreateAppointmentStorage
	businessTimeStore FindBusinessTimeStorage
	closedDateStore   FindClosedDateStorage
	staffWorkingStore FindStaffWorkingStorage
}

func NewCreateAppointmentBiz(store CreateAppointmentStorage, businessTimeStore FindBusinessTimeStorage,
	closedDateStore FindClosedDateStorage, staffWorkingStore FindStaffWorkingStorage) *createAppointmentBiz {
	return &createAppointmentBiz{store: store, businessTimeStore: businessTimeStore,
		closedDateStore: closedDateStore, staffWorkingStore: staffWorkingStore}
}

func (biz *createAppointmentBiz) CreateAppointment(ctx context.Context, data *appointmentmodel.AppointmentCreate) error {
	oldData, err := biz.store.FindDataByDuration(ctx, data.StaffID, data.StartTime, data.EndTime)

	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	if oldData != nil {
		return errors.New("Appointment has duplicated")
	}

	bizTime, errBT := biz.businessTimeStore.Find(ctx, data.StoreID, data.StartTime.Weekday().String())
	if errBT != nil {
		if errBT == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return errBT
	}
	if bizTime != nil {
		if bizTime.StartTime.Hour() > data.StartTime.Hour() ||
			bizTime.EndTime.Hour() < data.StartTime.Hour() ||
			bizTime.EndTime.Hour() < data.EndTime.Hour() {
			return errors.New("The store is not working at that time")
		} else if bizTime.StartTime.Hour() == data.StartTime.Hour() &&
			bizTime.StartTime.Minute() > data.StartTime.Minute() {
			return errors.New("The store is not working at that time")
		} else if bizTime.EndTime.Hour() == data.StartTime.Hour() &&
			bizTime.EndTime.Minute() < data.EndTime.Minute() {
			return errors.New("The store is not working at that time")
		}
	}

	closedDate, errCD := biz.closedDateStore.Find(ctx, data.StoreID, data.StartTime, data.EndTime)
	if errCD != nil && errCD != gorm.ErrRecordNotFound {
		return errCD
	}
	if closedDate != nil {
		return errors.New("The store is not working at that time")
	}

	staffWorking, errSW := biz.staffWorkingStore.Find(ctx, data.StaffID, data.StartTime, data.EndTime)
	if errSW != nil && errSW != gorm.ErrRecordNotFound {
		return errSW
	}
	if staffWorking != nil {
		return errors.New("Staff is not working at that time")
	}

	if errStore := biz.store.Create(ctx, data); errStore != nil {
		return errStore
	}

	ids := make([]*int, len(data.Service))
	for i, item := range data.Service {
		ids[i] = item
	}

	if errInsert := biz.store.InsertMiddleTable(ctx, ids, data.Id); errInsert != nil {
		return errInsert
	}

	return nil
}
