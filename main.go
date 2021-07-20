package main

import (
	"DemoProject/component/appctx"
	"DemoProject/modules/staff/stafftransport/ginstaff"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/Fresha?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Loi cmnr")
	}

	if err := runService(db); err != nil {
		panic("loi roi")
	}

}

func runService(db *gorm.DB) error {
	r := gin.Default()

	appCtx := appctx.NewAppContext(db)

	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	staff := r.Group("/staff")
	{
		staff.POST("", ginstaff.CreateStaff(appCtx))
		staff.GET("", ginstaff.ListStaff(appCtx))
		staff.GET("/:id", ginstaff.GetStaff(appCtx))
		staff.PATCH("/:id", ginstaff.UpdateStaff(appCtx))
		staff.DELETE("/:id", ginstaff.DeleteStaff(appCtx))
	}

	return r.Run()
}
