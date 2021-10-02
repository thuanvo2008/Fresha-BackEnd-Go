package appointmentstorage

import (
	"context"
)

func (s *sqlStore) InsertMiddleTable(ctx context.Context, ids []*int, appointmentId int) error {
	db := s.db

	for _, item := range ids {
		if err := db.Exec("Insert Into appointment_service(service_id,appointment_id) "+
			"value(?,?)", item, appointmentId).Error; err != nil {
			return err
		}
	}

	return nil
}
