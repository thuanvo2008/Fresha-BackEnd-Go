package appointmentstorage

import (
	"DemoProject/common"
	"context"
)

func (s *sqlStore) CancelAppointment(ctx context.Context, id int, cancelReason string) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(map[string]interface{}{
		"status":        0,
		"cancel_reason": cancelReason,
	}).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
