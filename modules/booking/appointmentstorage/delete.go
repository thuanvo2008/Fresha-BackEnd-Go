package appointmentstorage

import "context"

func (s *sqlStore) Delete(ctx context.Context, appointmentId int) error {
	db := s.db

	if err := db.Exec("DELETE FROM appointments where id = ? ", appointmentId).Error; err != nil {
		return err
	}

	return nil
}
