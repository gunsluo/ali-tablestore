package client

import (
	"bytes"
	"net/http"
	"time"

	"github.com/gunsluo/ali-tablestore/ots/security"
)

const (
	API_VERSION         = "2014-08-08"
	HTTP_REQUEST_METHOD = "POST"
	DEFAULT_TIMEOUT     = 5
)

type RequestBuilder struct {
	AccessKeyID     string
	AccessKeySecret string
	instance        string
	region          string
	internal        bool
	method          string
}

func NewRequestBuilder(accessKeyID, accessKeySecret string) *RequestBuilder {
	builder := new(RequestBuilder)
	builder.method = HTTP_REQUEST_METHOD
	builder.AccessKeyID = accessKeyID
	builder.AccessKeySecret = accessKeySecret

	return builder
}

func (this *RequestBuilder) NewRequest(operation string, content []byte) *http.Request {

	canonicalHeaders := this.newCanonicalHeaders(content)
	stringToSign := this.stringToSign(operation, canonicalHeaders)

	// signature
	signature := security.Signature(this.AccessKeySecret, stringToSign)
	canonicalHeaders.AddHeader(XHeaderKey.Signature, signature)

	return this.newRequest(operation, content, canonicalHeaders)
}

func (this *RequestBuilder) newCanonicalHeaders(content []byte) *CanonicalHeaders {

	t := time.Now()
	utc := t.UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT")

	// security request body
	contentMd5 := security.ContentMd5(content)

	canonicalHeaders := NewCanonicalHeaders()
	canonicalHeaders.AddHeader(XHeaderKey.APIVersion, API_VERSION)
	canonicalHeaders.AddHeader(XHeaderKey.AccessKeyId, this.AccessKeyID)
	canonicalHeaders.AddHeader(XHeaderKey.ContentMd5, contentMd5)
	canonicalHeaders.AddHeader(XHeaderKey.Date, utc)
	canonicalHeaders.AddHeader(XHeaderKey.InstanceName, this.instance)

	return canonicalHeaders
}

func (this *RequestBuilder) stringToSign(operation string, headers *CanonicalHeaders) string {
	return "/" + operation + "\n" + this.method + "\n\n" + headers.RequestString()
}

func (this *RequestBuilder) address(operation string) string {
	if this.internal == true {
		return "http://" + this.instance + "." + this.region + ".ots-internal.aliyuncs.com/" + operation
	} else {
		return "http://" + this.instance + "." + this.region + ".ots.aliyuncs.com/" + operation
	}
}

func (this *RequestBuilder) newRequest(operation string, content []byte, canonicalHeaders *CanonicalHeaders) *http.Request {

	url := this.address(operation)
	utc := canonicalHeaders.GetHeader(XHeaderKey.Date)
	contentMd5 := canonicalHeaders.GetHeader(XHeaderKey.ContentMd5)
	signature := canonicalHeaders.GetHeader(XHeaderKey.Signature)

	req, err := http.NewRequest("POST", url, bytes.NewReader(content))
	if err != nil {
		return nil
	}

	req.Header.Set("x-ots-date", utc)
	req.Header.Set("x-ots-apiversion", API_VERSION)
	req.Header.Set("x-ots-accesskeyid", this.AccessKeyID)
	req.Header.Set("x-ots-instancename", this.instance)
	req.Header.Set("x-ots-contentmd5", contentMd5)
	req.Header.Set("x-ots-signature", signature)

	return req
}
