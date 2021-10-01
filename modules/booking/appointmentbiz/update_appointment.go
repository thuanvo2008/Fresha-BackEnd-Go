package appointmentbiz

import (
	"DemoProject/modules/booking/appointmentmodel"
	"context"
	"errors"
)

type UpdateAppointmentStorage interface {
	UpdateData(ctx context.Context, id int, data *appointmentmodel.AppointmentUpdate) error
	FindDataByCondition(ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string) (*appointmentmodel.Appointment, error)
	DeleteMiddleTable(ctx context.Context, appointmentId int) error
	InsertMiddleTable(ctx context.Context, ids []*int, appointmentId int) error
}

type updateAppointmentBiz struct {
	store UpdateAppointmentStorage
}

func NewUpdateAppointmentBiz(store UpdateAppointmentStorage) *updateAppointmentBiz {
	return &updateAppointmentBiz{store: store}
}

func (biz *updateAppointmentBiz) UpdateAppointment(ctx context.Context, id int, data *appointmentmodel.AppointmentUpdate) error {
	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return errors.New("Data deleted")
	}

	if err := biz.store.DeleteMiddleTable(ctx, id); err != nil {
		return err
	}

	ids := make([]*int, len(data.Service))
	for i, item := range data.Service {
		ids[i] = item
	}

	if err := biz.store.InsertMiddleTable(ctx, ids, id); err != nil {
		return err
	}

	if err := biz.store.UpdateData(ctx, id, data); err != nil {
		return err
	}

	return nil
}
