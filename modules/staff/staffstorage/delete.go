package staffstorage

import (
	"DemoProject/common"
	"context"
)

func (s *sqlStore) DeleteData(ctx context.Context, id int) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(map[string]interface{}{
		"status": 0,
	}).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
