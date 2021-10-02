package appointmentstorage

import "context"

func (s *sqlStore) DeleteMiddleTable(ctx context.Context, appointmentId int) error {
	db := s.db

	if err := db.Exec("DELETE FROM appointment_service where appointment_id = ? ", appointmentId).Error; err != nil {
		return err
	}

	return nil
}
