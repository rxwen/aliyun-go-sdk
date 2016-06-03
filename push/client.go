package push

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/nu7hatch/gouuid"
	"github.com/rxwen/aliyun-go-sdk/parameters"
)

const baseURL = "https://cloudpush.aliyuncs.com/"

// Client is used to send push request.
type Client struct {
	Region          string
	AccessKeyID     string
	AccessKeySecret string
}

// NewClient function construct a new Client.
func NewClient(region, accessKeyID, accessKeySecret string) *Client {
	c := new(Client)
	c.Region = region
	c.AccessKeyID = accessKeyID
	c.AccessKeySecret = accessKeySecret

	return c
}

// SendRequest method sends out a PushRequest.
func (c Client) SendRequest(r *PushRequest) (string, error) {
	// https://help.aliyun.com/document_detail/mobilepush/api-reference/api-list/Push.html?spm=5176.docmobilepush/api-reference/call-method/common-parameters.6.124.OBDIvc
	ps := parameters.NewParameterSet()
	ps.Add("Format", "JSON")
	ps.Add("RegionId", c.Region)
	ps.Add("Version", "2015-08-27")
	ps.Add("AccessKeyId", c.AccessKeyID)
	ps.Add("SignatureMethod", "HMAC-SHA1")
	ps.Add("Timestamp", time.Now().UTC().Format("2006-01-02T03:04:05Z"))
	ps.Add("SignatureVersion", "1.0")
	id, _ := uuid.NewV4()
	ps.Add("SignatureNonce", id.String())
	// basic
	ps.Add("Action", "Push")
	ps.Add("AppKey", r.AppKey)
	// destination
	ps.Add("Target", r.Target)
	ps.Add("TargetValue", r.TargetValue)
	ps.Add("DeviceType", strconv.Itoa(r.DeviceType))
	// push config
	ps.Add("Type", strconv.Itoa(r.Type))
	ps.Add("Title", r.Title)
	ps.Add("Summary", r.Summary)
	ps.Add("Body", r.Body)
	// ios specific
	ps.Add("Remind", strconv.FormatBool(r.Remind))
	if len(r.IOSMusic) > 0 {
		ps.Add("iOSMusic", r.IOSMusic)
	}
	if len(r.IOSBadge) > 0 {
		ps.Add("iOSBadge", r.IOSBadge)
	}
	if len(r.IOSExtParameters) > 0 {
		ps.Add("iOSExtParameters", r.IOSExtParameters)
	}
	if len(r.ApnsEnv) > 0 {
		ps.Add("ApnsEnv", r.ApnsEnv)
	}
	// android specific
	ps.Add("AndroidOpenType", strconv.Itoa(r.AndroidOpenType))
	if len(r.AndroidActivity) > 0 {
		ps.Add("AndroidActivity", r.AndroidActivity)
	}
	if len(r.AndroidMusic) > 0 {
		ps.Add("AndroidMusic", r.AndroidMusic)
	}
	if len(r.AndroidOpenURL) > 0 {
		ps.Add("AndroidOpenUrl", r.AndroidOpenURL)
	}
	if len(r.AndroidExtParameters) > 0 {
		ps.Add("AndroidExtParameters", r.AndroidExtParameters)
	}
	// push control
	if len(r.PushTime) > 0 {
		ps.Add("PushTime", r.PushTime)
	}
	ps.Add("StoreOffline", strconv.FormatBool(r.StoreOffline))
	if r.StoreOffline && len(r.ExpireTime) > 0 {
		ps.Add("ExpireTime", r.ExpireTime)
	}
	// trace
	if len(r.BatchNumber) > 0 {
		ps.Add("BatchNumber", r.BatchNumber)
	}

	ps.Add("Signature", ps.Sign(ps.GetStringToSign(), c.AccessKeySecret))
	resp, err := http.Get(baseURL + "?" + ps.Concatenate())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
