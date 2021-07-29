package uploadprovider

import (
	"DemoProject/common"
	"context"
)

type UploadProvider interface {
	SaveFileProvider(ctx context.Context, data []byte, dst string) (*common.Image, error)
}
