package parameters

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"net/url"
	"sort"
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

func (s ParameterSet) Concatenate() string {
	var keys []string
	for k := range s.set {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	values := url.Values{}
	for _, k := range keys {
		values.Add(k, s.set[k])
	}

	return values.Encode()
}

func (s ParameterSet) GetStringToSign() string {
	cstr := s.Concatenate()
	values := url.Values{}
	values.Add("", cstr)
	result := "GET&%2F&" + values.Encode()[1:]
	return result
}

func (s ParameterSet) Sign(content, accessKeySecret string) string {
	// https://help.aliyun.com/document_detail/mobilepush/api-reference/call-method/signature.html?spm=5176.docmobilepush/api-reference/call-method/common-parameters.2.1.WGHWYh
	mac := hmac.New(sha1.New, []byte(accessKeySecret+"&"))
	mac.Write([]byte(content))
	sum := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(sum)
}
