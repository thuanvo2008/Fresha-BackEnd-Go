package businesstimestorage

import (
	"DemoProject/common"
	"DemoProject/modules/business_time/businesstimemodel"
	"context"
)

func (s *sqlStore) UpdateData(ctx context.Context, id int, data *businesstimemodel.BusinessTime) error {
	db := s.db

	if err := db.Table(data.TableName()).Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
