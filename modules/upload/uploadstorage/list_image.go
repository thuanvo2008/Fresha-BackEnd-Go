package uploadstorage

import (
	"DemoProject/common"
	"context"
)

func (s *sqlStore) ListImage(ctx context.Context, ids []int, moreKeys ...string) ([]common.Image, error) {
	var imgs []common.Image

	db := s.db

	db = db.Table(common.Image{}.TableName())
	if err := db.Where("id in (?)", ids).Find(&imgs).Error; err != nil {
		return nil, err
	}

	return imgs, nil
}
