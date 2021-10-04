package staffworkingstorage

import (
	"DemoProject/common"
	"DemoProject/modules/staff_working/staffworkingmodel"
	"context"
)

func (s *sqlStore) UpdateData(ctx context.Context, id int, data *staffworkingmodel.StaffWorking) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
