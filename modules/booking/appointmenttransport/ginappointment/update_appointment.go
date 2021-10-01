package ginappointment

import (
	"DemoProject/common"
	"DemoProject/component/appctx"
	"DemoProject/modules/booking/appointmentbiz"
	"DemoProject/modules/booking/appointmentmodel"
	"DemoProject/modules/booking/appointmentstorage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UpdateAppointment(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data appointmentmodel.AppointmentUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(err)
		}

		store := appointmentstorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := appointmentbiz.NewUpdateAppointmentBiz(store)
		if err := biz.UpdateAppointment(c.Request.Context(), id, &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
