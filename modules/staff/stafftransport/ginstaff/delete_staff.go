package ginstaff

import (
	"DemoProject/common"
	"DemoProject/component/appctx"
	"DemoProject/modules/staff/staffbiz"
	"DemoProject/modules/staff/staffstorage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteStaff(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := staffstorage.NewSQLStore(ctx.GetMainDBConnect())
		biz := staffbiz.NewDeleteStaffBiz(store)

		if err := biz.DeleteStaffBiz(c.Request.Context(), id); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
