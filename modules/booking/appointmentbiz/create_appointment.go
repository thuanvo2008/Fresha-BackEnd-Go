package appointmentbiz

import (
	"DemoProject/modules/booking/appointmentmodel"
	"context"
)

type CreateAppointmentStorage interface {
	Create(ctx context.Context, data *appointmentmodel.AppointmentCreate) error
	InsertMiddleTable(ctx context.Context, ids []*int, appointmentId int) error
}

type createAppointmentBiz struct {
	store CreateAppointmentStorage
}

func NewCreateAppointmentBiz(store CreateAppointmentStorage) *createAppointmentBiz {
	return &createAppointmentBiz{store: store}
}

func (biz *createAppointmentBiz) CreateAppointmentBiz(ctx context.Context, data *appointmentmodel.AppointmentCreate) error {
	if err := biz.store.Create(ctx, data); err != nil {
		return err
	}

	ids := make([]*int, len(data.Service))
	for i, item := range data.Service {
		ids[i] = item
	}

	if err := biz.store.InsertMiddleTable(ctx, ids, data.Id); err != nil {
		return err
	}

	return nil
}
