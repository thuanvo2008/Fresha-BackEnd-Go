package staffstorage

import (
	"DemoProject/common"
	"DemoProject/modules/staff/staffmodel"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) FindDataByCondition(ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string) (*staffmodel.Staff, error) {

	var result staffmodel.Staff

	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Where(condition).First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &result, nil
}
