package staffworkingstorage

import (
	"DemoProject/common"
	"DemoProject/modules/staff_working/staffworkingmodel"
	"context"
)

func (s *sqlStore) Create(ctx context.Context, data *staffworkingmodel.StaffWorking) error {
	db := s.db

	if err := db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
