package parameters_test

import (
	"github.com/rxwen/aliyun-go-sdk-push/parameters"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParameters(t *testing.T) {
	assert := assert.New(t)
	ps := parameters.NewParameterSet()
	ps.Add("Format", "XML")
	ps.Add("AccessKeyId", "testid")
	ps.Add("Action", "GetDeviceInfos")
	ps.Add("SignatureMethod", "HMAC-SHA1")
	ps.Add("RegionId", "cn-hangzhou")
	ps.Add("Devices", "e2ba19de97604f55b165576736477b74,92a1da34bdfd4c9692714917ce22d53d")
	ps.Add("SignatureNonce", "c4f5f0de-b3ff-4528-8a89-fa478bda8d80")
	ps.Add("SignatureVersion", "1.0")
	ps.Add("Version", "2015-08-27")
	ps.Add("AppKey", "23267207")
	ps.Add("Timestamp", "2016-03-29T03:59:24Z")
	assert.Equal("GET&%2F&AccessKeyId%3Dtestid%26Action%3DGetDeviceInfos%26AppKey%3D23267207%26Devices%3De2ba19de97604f55b165576736477b74%252C92a1da34bdfd4c9692714917ce22d53d%26Format%3DXML%26RegionId%3Dcn-hangzhou%26SignatureMethod%3DHMAC-SHA1%26SignatureNonce%3Dc4f5f0de-b3ff-4528-8a89-fa478bda8d80%26SignatureVersion%3D1.0%26Timestamp%3D2016-03-29T03%253A59%253A24Z%26Version%3D2015-08-27", ps.GetStringToSign())
	assert.Equal("Q4jj5vC+NRtz294V+oIW7gfaJ6U=", ps.Sign(ps.GetStringToSign(), "testsecret"))
}
