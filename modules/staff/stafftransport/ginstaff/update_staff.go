package ginstaff

import (
	"DemoProject/common"
	"DemoProject/component/appctx"
	"DemoProject/modules/staff/staffbiz"
	"DemoProject/modules/staff/staffmodel"
	"DemoProject/modules/staff/staffstorage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UpdateStaff(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data staffmodel.StaffUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := staffstorage.NewSQLStore(ctx.GetMainDBConnect())
		biz := staffbiz.NewUpdateStaffBiz(store)

		if err := biz.UpdateStaffBiz(c.Request.Context(), id, &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
