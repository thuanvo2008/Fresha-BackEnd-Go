package businesstimebiz

import (
	"DemoProject/modules/business_time/businesstimemodel"
	"context"
)

type FindBusinessTimeStorage interface {
	Find(ctx context.Context, storeID int, dayOfWeek string) (*businesstimemodel.BusinessTime, error)
}

type findBusinessTimeBiz struct {
	store FindBusinessTimeStorage
}

func NewFindBusinessTimeBiz(store FindBusinessTimeStorage) *findBusinessTimeBiz {
	return &findBusinessTimeBiz{store: store}
}

func (biz *findBusinessTimeBiz) FindBusinessTime(ctx context.Context,
	storeID int, dayOfWeek string) (*businesstimemodel.BusinessTime, error) {

	data, err := biz.store.Find(ctx, storeID, dayOfWeek)
	if err != nil {
		return nil, err
	}

	return data, nil
}
