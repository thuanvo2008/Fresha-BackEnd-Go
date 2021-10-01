package appointmentbiz

import (
	"DemoProject/common"
	"DemoProject/modules/booking/appointmentmodel"
	"DemoProject/modules/staff/staffmodel"
	"context"
)

type FindAppointmentStorage interface {
	FindDataByCondition(ctx context.Context,
		condition map[string]interface{},
		morekeys ...string) (*appointmentmodel.Appointment, error)
}

type findAppointmentBiz struct {
	store FindAppointmentStorage
}

func FindAppointmentBiz(store FindAppointmentStorage) *findAppointmentBiz {
	return &findAppointmentBiz{store: store}
}

func (biz *findAppointmentBiz) FindAppointment(ctx context.Context, id int) (*appointmentmodel.Appointment, error) {
	data, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id}, "Service", "Staff")
	if err != nil {
		if err != common.RecordNotFound {
			return nil, common.ErrCannotGetEntity(staffmodel.EntityName, err)
		}

		return nil, common.ErrCannotGetEntity(appointmentmodel.EntityName, err)
	}
	return data, nil
}
