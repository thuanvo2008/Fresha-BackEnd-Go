package ginstaff

import (
	"DemoProject/common"
	"DemoProject/component/appctx"
	"DemoProject/component/haser"
	"DemoProject/modules/staff/staffbiz"
	"DemoProject/modules/staff/staffmodel"
	"DemoProject/modules/staff/staffstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterStaff(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var staffCreate staffmodel.StaffCrete

		if err := c.ShouldBind(&staffCreate); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := staffstorage.NewSQLStore(ctx.GetMainDBConnect())
		md5 := haser.NewMd5Hash()
		biz := staffbiz.NewRegisterBusiness(store, md5)

		if err := biz.Register(c.Request.Context(), &staffCreate); err != nil {
			panic(err)
		}

		staffCreate.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(staffCreate.FakeID.String()))
	}
}
