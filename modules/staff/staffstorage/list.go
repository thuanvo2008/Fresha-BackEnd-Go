package staffstorage

import (
	"DemoProject/common"
	"DemoProject/modules/staff/staffmodel"
	"context"
)

func (s *sqlStore) ListDataByCondition(ctx context.Context,
	condition map[string]interface{},
	filter *staffmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]staffmodel.Staff, error) {
	db := s.db

	var result []staffmodel.Staff

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	db = db.Table(staffmodel.EntityName).Where(condition).Where("status in (1)")

	//if v := filter; v != nil {
	//	if v.CityId > 0 {
	//		db = db.Where("city_id = ?", v.CityId)
	//	}
	//}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if err := db.Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
