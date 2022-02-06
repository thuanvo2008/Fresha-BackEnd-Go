package ginappointment

import (
	"DemoProject/common"
	"DemoProject/component/appctx"
	"DemoProject/modules/booking/appointmentbiz"
	"DemoProject/modules/booking/appointmentstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CancelAppointment(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		type cancelAppointment struct {
			ID int `json:"id" gorm:"column:id"`
		}
		var cancel *cancelAppointment

		if err := c.ShouldBind(&cancel); err != nil {
			panic(err)
		}

		store := appointmentstorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := appointmentbiz.NewCancelAppointmentBiz(store)
		if err := biz.CancelAppointment(c.Request.Context(), cancel.ID); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
