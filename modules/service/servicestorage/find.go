package servicestorage

import (
	"DemoProject/common"
	"DemoProject/modules/service/servicemodel"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) FindService(ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string) (*servicemodel.Service, error) {

	var data servicemodel.Service
	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
