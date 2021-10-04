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

func ListStaff(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter staffmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := staffstorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := staffbiz.NewListStaffBiz(store)

		data, err := biz.ListStaff(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		//for i := range data {
		//	data[i].Mask(false)
		//
		//	if i == (len(data) - 1) {
		//		paging.NextCursor = data[i].FakeID.String()
		//	}
		//}

		c.JSON(http.StatusOK, common.NewSuccessResponse(data, paging, filter))
	}
}
