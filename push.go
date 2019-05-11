package gista

import (
	"github.com/aliforever/gista/constants"
	"github.com/aliforever/gista/errors"
	"github.com/aliforever/gista/responses"
)

type push struct {
	Ig *instagram
}

func newPush(i *instagram) *push {
	return &push{Ig: i}
}

func (p *push) Register(pushChannel, token string) (res *responses.PushRegister, err error) {
	if pushChannel != "mqtt" && pushChannel != "gcm" {
		err = errors.BadPushChannel(pushChannel)
		return
	}
	res = &responses.PushRegister{}
	dt := "android_gcm"
	if pushChannel == "mqtt" {
		dt = "android_mqtt"
	}
	mainPushChannel := "false"
	if pushChannel == "mqtt" {
		mainPushChannel = "true"
	}
	err = p.Ig.Client.Request(constants.PushRegister).
		SetSignedPost(false).
		AddPost("device_type", dt).
		AddPost("is_main_push_channel", mainPushChannel).
		AddPhoneIdPost().
		AddPost("device_token", token).
		AddCSRFPost().
		AddGuIdPost().
		AddUuIdPost().
		AddPost("users", *p.Ig.AccountId).
		GetResponse(res)
	return
}
