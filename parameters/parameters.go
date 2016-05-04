package parameters

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"net/url"
	"strings"
)

type ParameterSet struct {
	set map[string]string
}

func NewParameterSet() *ParameterSet {
	ps := new(ParameterSet)
	ps.set = make(map[string]string)
	return ps
}

func (s ParameterSet) Add(key, value string) {
	s.set[key] = value
}

func urlEncode(v string) string {
	values := url.Values{}
	values.Add("", v)
	return values.Encode()[1:]
}

func (s ParameterSet) Concatenate() string {
	values := url.Values{}
	for k, v := range s.set {
		values.Add(k, v)
	}

	result := values.Encode()
	result = strings.Replace(result, "+", "%20", -1)
	result = strings.Replace(result, "*", "%2A", -1)
	result = strings.Replace(result, "%7E", "~", -1)
	return result
}

func (s ParameterSet) GetStringToSign() string {
	cstr := s.Concatenate()
	result := "GET&%2F&" + urlEncode(cstr)
	return result
}

func (s ParameterSet) Sign(content, accessKeySecret string) string {
	// https://help.aliyun.com/document_detail/mobilepush/api-reference/call-method/signature.html?spm=5176.docmobilepush/api-reference/call-method/common-parameters.2.1.WGHWYh
	mac := hmac.New(sha1.New, []byte(accessKeySecret+"&"))
	mac.Write([]byte(content))
	sum := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(sum)
}
