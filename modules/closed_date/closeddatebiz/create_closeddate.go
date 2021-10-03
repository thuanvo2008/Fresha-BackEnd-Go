package closeddatebiz

import (
	"DemoProject/modules/closed_date/closeddatemodel"
	"context"
	"errors"
	"time"
)

type CreateClosedDateStorage interface {
	Create(ctx context.Context, data *closeddatemodel.ClosedDate) error
	Find(ctx context.Context, storeID int, startDateTime,
		endDateTime *time.Time) (*closeddatemodel.ClosedDate, error)
}

type createClosedDateBiz struct {
	store CreateClosedDateStorage
}

func NewCreateClosedDateBiz(store CreateClosedDateStorage) *createClosedDateBiz {
	return &createClosedDateBiz{store: store}
}

func (biz *createClosedDateBiz) CreateClosedDate(ctx context.Context, data *closeddatemodel.ClosedDate) error {
	oldData, err := biz.store.Find(ctx, data.StoreID, data.StartDateTime, data.EndDateTime)
	if err != nil {
		return err
	}

	if oldData != nil {
		return errors.New("Closed Date is duplicated")
	}

	if err := biz.store.Create(ctx, data); err != nil {
		return err
	}

	return nil
}
