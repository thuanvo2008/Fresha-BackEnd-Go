package appointmentstorage

import (
	"DemoProject/common"
	"DemoProject/modules/booking/appointmentmodel"
	"context"
)

func (s *sqlStore) Create(ctx context.Context, data *appointmentmodel.AppointmentCreate) error {
	db := s.db

	if err := db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
