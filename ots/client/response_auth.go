package client

import (
	"errors"
	"net/http"
	"time"

	"github.com/gunsluo/ali-tablestore/ots/security"
)

type ResponseAuthorization struct {
	AccessKeyID     string
	AccessKeySecret string
}

func NewResponseAuthorization(accessKeyID, accessKeySecret string) *ResponseAuthorization {
	author := new(ResponseAuthorization)
	author.AccessKeyID = accessKeyID
	author.AccessKeySecret = accessKeySecret

	return author
}

func (this *ResponseAuthorization) stringToAuthorization(operation string, headers *CanonicalHeaders) string {

	stringToSign := headers.ResponseString() + "/" + operation
	// signature
	signature := security.Signature(this.AccessKeySecret, stringToSign)

	return "OTS " + this.AccessKeyID + ":" + signature
}

func (this *ResponseAuthorization) Auth(operation string, content []byte, resp *http.Response) error {

	contentMd5 := resp.Header.Get(XHeaderKey.ContentMd5.String())
	date := resp.Header.Get(XHeaderKey.Date.String())
	contentType := resp.Header.Get(XHeaderKey.ContentType.String())
	requestId := resp.Header.Get(XHeaderKey.RequestId.String())
	authorization := resp.Header.Get(XHeaderKey.Authorization.String())

	// auth signature
	canonicalHeaders := NewCanonicalHeaders()
	canonicalHeaders.AddHeader(XHeaderKey.ContentMd5, contentMd5)
	canonicalHeaders.AddHeader(XHeaderKey.ContentType, contentType)
	canonicalHeaders.AddHeader(XHeaderKey.Date, date)
	canonicalHeaders.AddHeader(XHeaderKey.RequestId, requestId)
	stringToAuth := this.stringToAuthorization(operation, canonicalHeaders)
	if authorization != stringToAuth {
		return errors.New("Signature mismatch from response in header.")
	}

	// auth date
	otsTime, err := time.Parse("Mon, 02 Jan 2006 15:04:05 GMT", date)
	if err != nil {
		return errors.New("x-ots-date of response in header pattern error.")
	}

	t := time.Now()
	systemTime := t.UTC()
	duration := systemTime.Sub(otsTime).Minutes()
	if duration > 15 || duration < -15 {
		return errors.New("Mismatch between system time and x-ots-date of response in header.")
	}

	// auth md5
	expectMd5 := security.ContentMd5(content)
	if contentMd5 != expectMd5 {
		return errors.New("Mismatch between MD5 value of response content and x-ots-contentmd5 in header.")
	}

	return nil
}
