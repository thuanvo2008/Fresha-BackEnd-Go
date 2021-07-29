package uploadstorage

import (
	"DemoProject/common"
	"DemoProject/component/appctx"
)

func (s *sqlStore) DeleteImage(ctx appctx.AppContext, ids []int) error {
	db := s.db

	if err := db.Table(common.Image{}.TableName()).
		Where("id in (?)", ids).
		Delete(nil).Error; err != nil {
		return err
	}

	return nil
}
