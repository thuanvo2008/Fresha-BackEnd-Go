package ginbusinesstime

import (
	"DemoProject/common"
	"DemoProject/component/appctx"
	"DemoProject/modules/business_time/businesstimebiz"
	"DemoProject/modules/business_time/businesstimemodel"
	"DemoProject/modules/business_time/businesstimestorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Create(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data businesstimemodel.BusinessTime

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := businesstimestorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := businesstimebiz.NewCreateBusinessTimeBiz(store)

		if err := biz.CreateBusinessTime(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
