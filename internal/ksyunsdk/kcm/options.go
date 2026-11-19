package kcm

import (
	common "github.com/certimate-go/plugin-provider-ksyun-slb-deployer/internal/ksyunsdk/zz-shared-common"
)

func WithAkSk(ak, sk string) common.OptionsFunc {
	return common.WithAkSk(ak, sk)
}
