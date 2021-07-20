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

func GetStaff(ctx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := staffstorage.NewSQLStore(ctx.GetMainDBConnect())
		biz := staffbiz.NewGetStaffBiz(store)

		data, err := biz.GetStaffBiz(context.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
