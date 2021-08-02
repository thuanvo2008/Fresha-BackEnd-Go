package main

import (
	"DemoProject/component/appctx"
	"DemoProject/component/uploadprovider"
	"DemoProject/middleware"
	"DemoProject/modules/staff/stafftransport/ginstaff"
	"DemoProject/modules/upload/uploadtransport/ginupload"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/Fresha?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Loi cmnr")
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

	staff := v1.Group("/staff", middleware.RequiredAuth(appCtx))
	{
		staff.POST("", ginstaff.CreateStaff(appCtx))
		staff.GET("", ginstaff.ListStaff(appCtx))
		staff.GET("/:id", ginstaff.GetStaff(appCtx))
		staff.PATCH("/:id", ginstaff.UpdateStaff(appCtx))
		staff.DELETE("/:id", ginstaff.DeleteStaff(appCtx))
	}

	return r.Run()
}
