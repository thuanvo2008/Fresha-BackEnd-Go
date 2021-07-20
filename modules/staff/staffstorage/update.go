package staffstorage

import (
	"DemoProject/common"
	"DemoProject/modules/staff/staffmodel"
	"context"
)

func (s *sqlStore) UpdateData(ctx context.Context,
	id int,
	data *staffmodel.StaffUpdate,
) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
