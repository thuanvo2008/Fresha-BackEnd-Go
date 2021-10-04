package servicebiz

import (
	"DemoProject/modules/service/servicemodel"
	"context"
)

type CreateServiceStorage interface {
	CreateService(ctx context.Context, data *servicemodel.ServiceCreate) error
}

type createServiceBiz struct {
	store CreateServiceStorage
}

func NewCreateServiceBiz(store CreateServiceStorage) *createServiceBiz {
	return &createServiceBiz{store: store}
}

func (biz *createServiceBiz) CreateService(ctx context.Context, data *servicemodel.ServiceCreate) error {
	if err := biz.store.CreateService(ctx, data); err != nil {
		return err
	}
	return nil
}
