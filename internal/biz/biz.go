package biz

import (
	"github.com/google/wire"
	wechat "kratos-mall/internal/biz/pay/wechat/impl"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(wechat.NewWechatUsecase())
