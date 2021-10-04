package ginservice

import (
	"DemoProject/common"
	"DemoProject/component/appctx"
	"DemoProject/modules/service/servicebiz"
	"DemoProject/modules/service/servicemodel"
	"DemoProject/modules/service/servicestorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateNew(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data servicemodel.ServiceCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := servicestorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := servicebiz.NewCreateServiceBiz(store)

		if err := biz.CreateService(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
