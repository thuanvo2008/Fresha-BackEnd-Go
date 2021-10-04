package ginappointment

import (
	"DemoProject/common"
	"DemoProject/component/appctx"
	"DemoProject/modules/booking/appointmentbiz"
	"DemoProject/modules/booking/appointmentmodel"
	"DemoProject/modules/booking/appointmentstorage"
	"DemoProject/modules/business_time/businesstimestorage"
	"DemoProject/modules/closed_date/closeddatestorage"
	"DemoProject/modules/staff_working/staffworkingstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateAppointment(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var appointment appointmentmodel.AppointmentCreate

		if err := c.ShouldBind(&appointment); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		businessTimeStore := businesstimestorage.NewSQLStore(ctx.GetMainDBConnection())
		closedDateStore := closeddatestorage.NewSQLStore(ctx.GetMainDBConnection())
		staffWorkingStore := staffworkingstorage.NewSQLStore(ctx.GetMainDBConnection())
		store := appointmentstorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := appointmentbiz.NewCreateAppointmentBiz(store, businessTimeStore, closedDateStore, staffWorkingStore)
		if err := biz.CreateAppointment(c.Request.Context(), &appointment); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
