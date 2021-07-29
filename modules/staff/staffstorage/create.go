package staffstorage

import (
	"DemoProject/common"
	"DemoProject/modules/staff/staffmodel"
	"context"
)

func (s *sqlStore) Create(ctx context.Context, data *staffmodel.StaffCrete) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
