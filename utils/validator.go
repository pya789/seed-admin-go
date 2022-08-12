package utils

import (
	"github.com/gookit/validate"
	"github.com/gookit/validate/locales/zhcn"
)

// 验证数据
func ParamsVerify(st any) error {
	// 中文输出注册到全局 如有需要可按照RegisterGlobal自定义msg信息
	zhcn.RegisterGlobal()
	v := validate.New(st)
	if !v.Validate() {
		return v.Errors.OneError()
	}
	return nil
}
