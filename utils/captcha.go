package utils

import (
	"bytes"
	"fmt"
	"seed-admin/common"
	"sync"

	"github.com/dchest/captcha"
	"go.uber.org/zap"
)

type Captcha struct {
}

var once sync.Once
var cap *Captcha

// 单例
func NewCaptcha() *Captcha {
	once.Do(func() {
		cap = &Captcha{}
	})
	return cap
}

// 验证
func CaptchaVerify(captchaId string, value string) error {
	if captcha.VerifyString(captchaId, value) {
		return nil
	}
	return fmt.Errorf("验证码错误")
}

// 创建一个图片验证码
func (c *Captcha) CreateImage() string {
	captchaId := captcha.NewLen(common.CONFIG.Int("captcha.len"))
	return captchaId
}

// 重载验证码
func (c *Captcha) Reload(captchaId string) bool {
	return captcha.Reload(captchaId)
}

// 获取二进制图片
func (c *Captcha) ImageByte(captchaId string) ([]byte, error) {
	var content bytes.Buffer
	if err := captcha.WriteImage(&content, captchaId, common.CONFIG.Int("captcha.width"), common.CONFIG.Int("captcha.height")); err != nil {
		common.LOG.Error("验证码获取失败", zap.String("err", err.Error()))
		return nil, err
	}
	return content.Bytes(), nil
}
