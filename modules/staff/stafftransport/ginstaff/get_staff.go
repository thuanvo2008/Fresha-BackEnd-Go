package ginstaff

import (
	"DemoProject/common"
	"DemoProject/component/appctx"
	"DemoProject/modules/staff/staffbiz"
	"DemoProject/modules/staff/staffstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetStaff(ctx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		uid, err := common.FromBase58(context.Param("id"))

		//id, err := strconv.Atoi(context.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := staffstorage.NewSQLStore(ctx.GetMainDBConnect())
		biz := staffbiz.NewGetStaffBiz(store)

		data, err := biz.GetStaffBiz(context.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}

		data.Mask(false)

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
