package appointmentbiz

import (
	"DemoProject/modules/booking/appointmentmodel"
	"context"
	"errors"
	"gorm.io/gorm"
	"time"
)

type CancelAppointmentStorage interface {
	FindDataByCondition(ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string) (*appointmentmodel.Appointment, error)
	CancelAppointment(ctx context.Context, id int, cancelReason string) error
}

type cancelAppointmentBiz struct {
	store CancelAppointmentStorage
}

func NewCancelAppointmentBiz(store CancelAppointmentStorage) *cancelAppointmentBiz {
	return &cancelAppointmentBiz{store: store}
}

func (biz *cancelAppointmentBiz) CancelAppointment(ctx context.Context, id int, cancelReason string) error {
	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return gorm.ErrRecordNotFound
	}

	if oldData.Status == 0 {
		return errors.New("Data deleted")
	}

	now := time.Now().Add(2 * time.Hour)
	if now.After(*oldData.StartTime) {
		return errors.New("Cann't cancel this appointment")
	}

	if err := biz.store.CancelAppointment(ctx, id, cancelReason); err != nil {
		return err
	}

	return nil
}
