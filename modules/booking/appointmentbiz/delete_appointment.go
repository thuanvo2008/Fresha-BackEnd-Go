package appointmentbiz

import (
	"context"
)

type DeleteAppointmentStorage interface {
	Delete(ctx context.Context, appointmentId int) error
}

type deleteAppointmentBiz struct {
	store DeleteAppointmentStorage
}

func NewDeleteAppointmentBiz(store DeleteAppointmentStorage) *deleteAppointmentBiz {
	return &deleteAppointmentBiz{store: store}
}

func (biz *deleteAppointmentBiz) DeleteAppointment(ctx context.Context, id int) error {
	if err := biz.store.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
