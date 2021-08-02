package staffbiz

import (
	"DemoProject/common"
	"DemoProject/component/appctx"
	"DemoProject/component/tokenprovider"
	"DemoProject/modules/staff/staffmodel"
	"context"
)

type LoginStorage interface {
	FindDataByCondition(ctx context.Context,
		condition map[string]interface{},
		monkeys ...string) (*staffmodel.Staff, error)
}

type loginBusiness struct {
	appCtx        appctx.AppContext
	storeUser     LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginBusiness(storeUser LoginStorage, tokenProvider tokenprovider.Provider,
	hasher Hasher, expiry int) *loginBusiness {
	return &loginBusiness{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

func (business *loginBusiness) Login(ctx context.Context, data *staffmodel.Login) (*tokenprovider.Token, error) {
	staff, err := business.storeUser.FindDataByCondition(ctx, map[string]interface{}{"email": data.Email})

	if err != nil {
		return nil, staffmodel.ErrUsernameOrPasswordInvalid
	}

	passHashed := business.hasher.Hash(data.Password + staff.Salt)

	if staff.Password != passHashed {
		return nil, staffmodel.ErrUsernameOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		StaffId: staff.Id,
		Role:    staff.Role,
	}

	accessToken, err := business.tokenProvider.Generate(payload, business.expiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	//refreshToken, err := business.tokenProvider.Generate(payload, business.tkCfg.GetRtExp())
	//if err != nil {
	//	return nil, common.ErrInternal(err)
	//}

	//account := usermodel.NewAccount(accessToken, refreshToken)

	return accessToken, nil
}
