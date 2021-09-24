package ginappointment

import (
	"DemoProject/common"
	"DemoProject/component/appctx"
	"DemoProject/modules/booking/appointmentbiz"
	"DemoProject/modules/booking/appointmentmodel"
	"DemoProject/modules/booking/appointmentstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateAppointment(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var appointment appointmentmodel.AppointmentCreate

		if err := c.ShouldBind(&appointment); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := appointmentstorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := appointmentbiz.NewCreateAppointmentBiz(store)
		if err := biz.CreateAppointmentBiz(c.Request.Context(), &appointment); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
