package ginupload

import (
	"DemoProject/common"
	"DemoProject/component/appctx"
	"DemoProject/modules/upload/uploadbiz"
	"DemoProject/modules/upload/uploadstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadFile(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		folder := c.DefaultPostForm("folder", "img")

		file, err := fileHeader.Open()

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		defer file.Close()

		myBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(myBytes); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := uploadstorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := uploadbiz.NewUploadBiz(ctx.UploadProvider(), store)
		img, err := biz.Upload(c.Request.Context(), myBytes, folder, fileHeader.Filename)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(img))
	}
}
