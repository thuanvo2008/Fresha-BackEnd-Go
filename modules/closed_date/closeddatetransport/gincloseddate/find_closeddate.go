package gincloseddate

import (
	"DemoProject/common"
	"DemoProject/component/appctx"
	"DemoProject/modules/closed_date/closeddatebiz"
	"DemoProject/modules/closed_date/closeddatemodel"
	"DemoProject/modules/closed_date/closeddatestorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Find(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data closeddatemodel.ClosedDate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := closeddatestorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := closeddatebiz.NewFindClosedDateBiz(store)

		result, err := biz.FindClosedDate(c.Request.Context(), data.StoreID,
			data.StartDateTime, data.EndDateTime)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
