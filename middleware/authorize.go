package middleware

import (
	"DemoProject/common"
	"DemoProject/component/appctx"
	"DemoProject/component/tokenprovider/jwt"
	"DemoProject/modules/staff/staffstorage"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("wrong authen header"),
		fmt.Sprintf("ErrWrongAuthHeader"),
	)
}

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")

	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}

	return parts[1], nil
}

func RequiredAuth(appCtx appctx.AppContext) func(c *gin.Context) {
	tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))

		if err != nil {
			panic(err)
		}

		db := appCtx.GetMainDBConnect()
		store := staffstorage.NewSQLStore(db)

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			panic(err)
		}

		staff, err := store.FindDataByCondition(c.Request.Context(), map[string]interface{}{"id": payload.StaffId})

		if err != nil {
			panic(err)
		}

		if staff.Status == 0 {
			panic(common.ErrNoPermission(errors.New("staff has been deleted or banned")))
		}

		staff.Mask(false)

		c.Set(common.CurrentStaff, staff)
		c.Next()
	}
}
