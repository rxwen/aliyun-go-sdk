package directmail

// https://help.aliyun.com/document_detail/29430.html?spm=5176.doc29428.6.108.pG6UlQ
import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/nu7hatch/gouuid"
	"github.com/rxwen/aliyun-go-sdk/parameters"
)

const baseURL = "https://dm.aliyuncs.com/"

// Client is used to send mail request.
type Client struct {
	AccessKeyID     string
	AccessKeySecret string
}

// NewClient function construct a new Client.
func NewClient(accessKeyID, accessKeySecret string) *Client {
	c := new(Client)
	c.AccessKeyID = accessKeyID
	c.AccessKeySecret = accessKeySecret

	return c
}

// SendRequest method sends out a MailRequest.
func (c Client) SendRequest(r *MailRequest) (string, error) {
	// https://help.aliyun.com/document_detail/29440.html?spm=5176.doc29444.2.1.hYchmR
	ps := parameters.NewParameterSet()
	ps.Add("Format", "JSON")
	ps.Add("Version", "2015-11-23")
	ps.Add("AccessKeyId", c.AccessKeyID)
	ps.Add("SignatureMethod", "HMAC-SHA1")
	ps.Add("Timestamp", time.Now().UTC().Format("2006-01-02T03:04:05Z"))
	ps.Add("SignatureVersion", "1.0")
	id, _ := uuid.NewV4()
	ps.Add("SignatureNonce", id.String())

	ps.Add("Action", r.Action)
	ps.Add("AccountName", r.AccountName)
	if r.ReplyToAddress {
		ps.Add("ReplyToAddress", "true")
	} else {
		ps.Add("ReplyToAddress", "false")
	}
	ps.Add("AddressType", strconv.Itoa(r.AddressType))
	ps.Add("ToAddress", r.ToAddress)
	if len(r.FromAlias) > 0 {
		ps.Add("FromAlias", r.FromAlias)
	}
	if len(r.Subject) > 0 {
		ps.Add("Subject", r.Subject)
	}
	if len(r.HtmlBody) > 0 {
		ps.Add("HtmlBody", r.HtmlBody)
	}
	if len(r.TextBody) > 0 {
		ps.Add("TextBody", r.TextBody)
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

	var v map[string]string
	dec := json.NewDecoder(bytes.NewReader(body))
	if err := dec.Decode(&v); err != nil {
		return "", err
	}
	if _, ok := v["Code"]; ok {
		return "", errors.New(string(body))
	}
	return string(body), nil
}
