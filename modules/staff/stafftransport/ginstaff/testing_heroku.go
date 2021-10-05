package ginstaff

import (
	"DemoProject/common"
	"DemoProject/component/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestingHeroku(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, common.SimpleSuccessResponse("hello world"))
	}
}
