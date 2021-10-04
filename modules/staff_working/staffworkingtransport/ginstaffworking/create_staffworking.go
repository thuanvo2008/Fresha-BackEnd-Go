package ginstaffworking

import (
	"DemoProject/common"
	"DemoProject/component/appctx"
	"DemoProject/modules/staff_working/staffworkingbiz"
	"DemoProject/modules/staff_working/staffworkingmodel"
	"DemoProject/modules/staff_working/staffworkingstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Create(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data staffworkingmodel.StaffWorking

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := staffworkingstorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := staffworkingbiz.NewCreateStaffWorkingBiz(store)

		if err := biz.CreateStaffWorking(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
