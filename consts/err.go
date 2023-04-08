package consts

import (
	"errors"
	"github.com/noahlsl/noah/tools/strx"
)

var (
	ErrTokenExpired         = strx.S2b("Token has expired")                      // Token 过期
	ErrIdempotence          = strx.S2b("Requests limit, please try again later") // 请求限制，请稍后再试
	ErrAuthorizationExpired = strx.S2b("Authorization has expired")
)

var (
	ErrSysBadRequest   = errors.New("网络异常,请稍后重试")
	ErrSysTokenExpired = errors.New("token失效")
	ErrSysAuthFailed   = errors.New("权限校验失败")
	ErrRequestLimit    = errors.New("请求频率限制")
	ErrTimeout         = errors.New("请求超时")
	ErrIPLimit         = errors.New("IP限制")
	ErrImageSizeLimit  = errors.New("图片大小限制")
	ErrImageSuffix     = errors.New("图片后缀错误")
)

// 2、定义Code
const (
	ErrCodeSysBadRequest   = 1100 // 请求错误
	ErrCodeSysTokenExpired = 1101 // Token失效
	ErrCodeSysAuthFailed   = 1102 // 权限校验失败
	ErrCodeRequestLimit    = 1103 // 请求频率限制
	ErrCodeTimeout         = 1104 // 请求超时
	ErrCodeIPLimit         = 1105 // IP限制
	ErrCodeImageSizeLimit  = 1106 // 图片大小限制
	ErrCodeImageSuffix     = 1107 // 图片后缀错误
)
