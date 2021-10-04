package closeddatebiz

import (
	"DemoProject/modules/closed_date/closeddatemodel"
	"context"
	"time"
)

type FindClosedDateStorage interface {
	Find(ctx context.Context, storeID int, startDateTime,
		endDateTime *time.Time) (*closeddatemodel.ClosedDate, error)
}

type findClosedDateBiz struct {
	store FindClosedDateStorage
}

func NewFindClosedDateBiz(store FindClosedDateStorage) *findClosedDateBiz {
	return &findClosedDateBiz{store: store}
}

func (biz *findClosedDateBiz) FindClosedDate(ctx context.Context, storeID int, startTime,
	endTime *time.Time) (*closeddatemodel.ClosedDate, error) {

	data, err := biz.store.Find(ctx, storeID, startTime, endTime)
	if err != nil {
		return nil, err
	}

	return data, nil
}
