package appointmentstorage

import (
	"DemoProject/common"
	"context"
)

func (s *sqlStore) CancelAppointment(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table("appointments").Where("id = ?", id).Updates(map[string]interface{}{
		"status": 0,
	}).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
