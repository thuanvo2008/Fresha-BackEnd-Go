package businesstimebiz

import (
	"DemoProject/modules/business_time/businesstimemodel"
	"context"
	"errors"
	"time"
)

type CreateBusinessTimeStorage interface {
	Create(ctx context.Context, data *businesstimemodel.BusinessTime) error
	Find(ctx context.Context, storeID int, appointmentStartTime,
		appointmentEndTime *time.Time, dayOfWeek string) (*businesstimemodel.BusinessTime, error)
}

type createBusinessTimeBiz struct {
	store CreateBusinessTimeStorage
}

func NewCreateBusinessTimeBiz(store CreateBusinessTimeStorage) *createBusinessTimeBiz {
	return &createBusinessTimeBiz{store: store}
}

func (biz *createBusinessTimeBiz) CreateBusinessTime(ctx context.Context, data *businesstimemodel.BusinessTime) error {
	oldData, err := biz.store.Find(ctx, data.StoreID, data.StartTime, data.EndTime, data.DayOfWeek.String())
	if err != nil {
		return err
	}

	if oldData != nil {
		return errors.New("Business Time is duplicated")
	}

	if err := biz.store.Create(ctx, data); err != nil {
		return err
	}

	return nil
}
