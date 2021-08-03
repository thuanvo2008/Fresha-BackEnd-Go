package ginstaff

import (
	"DemoProject/common"
	"DemoProject/component/appctx"
	"DemoProject/modules/staff/staffbiz"
	"DemoProject/modules/staff/staffmodel"
	"DemoProject/modules/staff/staffstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateStaff(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data staffmodel.StaffCrete

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := staffstorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := staffbiz.NewCreateStaffBiz(store)

		if err := biz.CreateStaffBiz(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.GenUID(common.DBTypeStaff)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeID.String()))
	}
}
