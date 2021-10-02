package appointmentstorage

import (
	"DemoProject/common"
	"DemoProject/modules/booking/appointmentmodel"
	"context"
	"time"
)

func (s *sqlStore) FindDataByDuration(ctx context.Context, staffId int, startTime *time.Time,
	endTime *time.Time) (*appointmentmodel.Appointment, error) {
	db := s.db
	var result appointmentmodel.Appointment

	if err := db.Where("staff_id = ?", staffId).
		Where("start_time >= ? And ? <= end_time", startTime, startTime).
		Or("start_time >= ? And ? <= end_time", endTime, endTime).
		Or("? >= start_time And end_time <= ?", startTime, endTime).
		Where("status = 1").First(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return &result, nil
}
