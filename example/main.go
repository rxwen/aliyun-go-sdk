package main

import (
	"fmt"

	"github.com/rxwen/aliyun-go-sdk/directmail"
	"github.com/rxwen/aliyun-go-sdk/push"
)

const (
	accessKeyID     = "xxx"
	accessKeySecret = "xxx"
	appKey          = "xxx"
)

func main() {
	req := push.PushRequest{
		// push target
		AppKey:     appKey,
		Target:     "all",
		DeviceType: 1,

		Type:    1,
		Title:   "hello",
		Body:    "world",
		Summary: "sum",

		AndroidOpenType: 1,

		// Message control
		StoreOffline: true,
	}
	c := push.NewClient("cn-hangzhou", accessKeyID, accessKeySecret)
	body, _ := c.SendRequest(&req)
	fmt.Println(body)
}

func sendMail() {
	req := directmail.MailRequest{
		Action:         "SingleSendMail",
		AccountName:    "notice@push.liveliy.com",
		ReplyToAddress: true,
		ToAddress:      "xxx@gmail.com",
		FromAlias:      "xxx",
		Subject:        "please activate your account",
		HtmlBody:       "please click <a href='google.com'>here</a> to activate your account",
	}

	c := directmail.NewClient(accessKeyID, accessKeySecret)
	body, _ := c.SendRequest(&req)
	fmt.Println(body)
}

func main() {
	sendMail()
}
