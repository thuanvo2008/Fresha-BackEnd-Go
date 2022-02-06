package appointmentstorage

import (
	"DemoProject/common"
	"DemoProject/modules/booking/appointmentmodel"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) FindDataByCondition(ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string) (*appointmentmodel.Appointment, error) {
	var result appointmentmodel.Appointment

	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Where(condition).Where("status = 1").First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &result, nil
}
