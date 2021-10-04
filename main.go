package main

import (
	"DemoProject/component/appctx"
	"DemoProject/component/uploadprovider"
	"DemoProject/middleware"
	"DemoProject/modules/booking/appointmenttransport/ginappointment"
	"DemoProject/modules/business_time/businesstimetransport/ginbusinesstime"
	"DemoProject/modules/service/servicetransport/ginservice"
	"DemoProject/modules/staff/stafftransport/ginstaff"
	"DemoProject/modules/staff_working/staffworkingtransport/ginstaffworking"
	"DemoProject/modules/upload/uploadtransport/ginupload"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/Fresha?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db = db.Debug()

	if err != nil {
		panic(gorm.ErrInvalidDB)
	}
	var provider uploadprovider.UploadProvider
	secretKey := "Fresha_CopDeal"
	if err := runService(db, provider, secretKey); err != nil {
		panic("loi roi")
	}

}

func runService(db *gorm.DB, provider uploadprovider.UploadProvider, secretKey string) error {
	r := gin.Default()

	appCtx := appctx.NewAppContext(db, provider, secretKey)
	r.Use(middleware.Recover(appCtx))

	v1 := r.Group("/v1")
	v1.POST("/uploadFile", ginupload.UploadFile(appCtx))
	v1.POST("/login", ginstaff.Login(appCtx))
	v1.POST("/register", ginstaff.RegisterStaff(appCtx))

	staff := v1.Group("/staff") //, middleware.RequiredAuth(appCtx)
	{
		staff.POST("", ginstaff.CreateStaff(appCtx))
		staff.GET("", ginstaff.ListStaff(appCtx))
		staff.GET("/:id", ginstaff.GetStaff(appCtx))
		staff.PATCH("/:id", ginstaff.UpdateStaff(appCtx))
		staff.DELETE("/:id", ginstaff.DeleteStaff(appCtx))
	}

	appointment := v1.Group("/appointment")
	{
		appointment.POST("", ginappointment.CreateAppointment(appCtx))
		appointment.GET("/:id", ginappointment.GetAppointment(appCtx))
		appointment.PATCH("/:id", ginappointment.UpdateAppointment(appCtx))
		appointment.DELETE("/:id", ginappointment.CancelAppointment(appCtx))
	}

	service := v1.Group("/service")
	{
		service.POST("", ginservice.CreateNew(appCtx))
	}

	business_time := v1.Group("/business_time")
	{
		business_time.POST("", ginbusinesstime.Create(appCtx))
		//business_time.GET()
	}

	staff_working := v1.Group("/staff_working")
	{
		staff_working.POST("", ginstaffworking.Create(appCtx))
	}

	return r.Run()
}
