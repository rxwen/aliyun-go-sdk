package main

import (
	//"net/http"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/nu7hatch/gouuid"
	"github.com/rxwen/aliyun-go-sdk-push/parameters"
)

const (
	baseURL = "https://cloudpush.aliyuncs.com/"

	accessKeyID     = "xxx"
	accessKeySecret = "xxx"
	appKey          = "xxx"
)

func main() {
	log.Info("start main")

	// https://help.aliyun.com/document_detail/mobilepush/api-reference/api-list/Push.html?spm=5176.docmobilepush/api-reference/call-method/common-parameters.6.124.OBDIvc
	ps := parameters.NewParameterSet()
	ps.Add("Format", "JSON")
	ps.Add("RegionId", "cn-hangzhou")
	ps.Add("Version", "2015-08-27")
	ps.Add("AccessKeyId", accessKeyID)
	ps.Add("SignatureMethod", "HMAC-SHA1")
	ps.Add("Timestamp", time.Now().UTC().Format("2006-01-02T03:04:05Z"))
	ps.Add("SignatureVersion", "1.0")
	id, _ := uuid.NewV4()
	ps.Add("SignatureNonce", id.String())
	ps.Add("Action", "Push")
	ps.Add("AppKey", appKey)
	ps.Add("Target", "all")
	ps.Add("TargetValue", "all")
	ps.Add("DeviceType", "1")
	ps.Add("Type", "1")
	ps.Add("Title", "hello")
	ps.Add("Summary", "hw")
	ps.Add("Body", id.String())
	ps.Add("StoreOffline", "true")
	ps.Add("Remind", "true")       // ios specific
	ps.Add("AndroidOpenType", "1") // android specific

	ps.Add("Signature", ps.Sign(ps.GetStringToSign(), accessKeySecret))
	resp, _ := http.Get(baseURL + "?" + ps.Concatenate())
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	log.Info(string(body))
}
