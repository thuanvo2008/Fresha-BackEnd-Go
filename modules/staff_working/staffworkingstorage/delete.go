package staffworkingstorage

import (
	"DemoProject/common"
	"DemoProject/modules/staff_working/staffworkingmodel"
	"context"
)

func (s *sqlStore) DeleteData(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(staffworkingmodel.StaffWorking{}.TableName()).
		Where("id = ?", id).Updates(map[string]interface{}{
		"status": 0,
	}).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
