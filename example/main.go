package main

import (
	"fmt"

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
