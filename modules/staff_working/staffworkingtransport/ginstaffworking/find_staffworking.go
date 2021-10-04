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

func Find(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data staffworkingmodel.StaffWorking

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := staffworkingstorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := staffworkingbiz.NewFindStaffWorkingBiz(store)

		result, err := biz.FindStaffWorking(c.Request.Context(), data.StaffID,
			data.StartTime, data.EndTime)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
