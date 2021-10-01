package appointmentstorage

import (
	"DemoProject/common"
	"DemoProject/modules/booking/appointmentmodel"
	"context"
)

func (s *sqlStore) UpdateData(ctx context.Context, id int, data *appointmentmodel.AppointmentUpdate) error {
	db := s.db

	if err := db.Table(data.TableName()).Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
