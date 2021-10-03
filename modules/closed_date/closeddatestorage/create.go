package closeddatestorage

import (
	"DemoProject/modules/closed_date/closeddatemodel"
	"context"
)

func (s *sqlStore) Create(ctx context.Context, data *closeddatemodel.ClosedDate) error {
	db := s.db

	if err := db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
