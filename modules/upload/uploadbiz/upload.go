package uploadbiz

import (
	"DemoProject/common"
	"context"
)

type CreateUploadStorage interface {
	CreateImage(context context.Context, data *common.Image) error
}

type uploadBiz struct {
	provider uploadprovider.UploadProvider
	imgStore CreateUploadStorage
}
