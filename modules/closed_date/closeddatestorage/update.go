package closeddatestorage

import (
	"DemoProject/common"
	"DemoProject/modules/closed_date/closeddatemodel"
	"context"
)

func (s *sqlStore) UpdateData(ctx context.Context, id int, data *closeddatemodel.ClosedDate) error {
	db := s.db

	if err := db.Table(data.TableName()).Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
