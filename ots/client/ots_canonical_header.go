package client

import "strings"

type HKey int

func (this HKey) String() string {

	pos := int(this)
	length := len(x_header_key_bitmap)

	if pos >= length {
		return ""
	}

	kb := x_header_key_bitmap[pos]
	return kb.Key
}

func (this HKey) Value() int {
	return int(this)
}

var XHeaderKey = struct {
	AccessKeyId   HKey
	APIVersion    HKey
	ContentMd5    HKey
	ContentType   HKey
	Date          HKey
	InstanceName  HKey
	RequestId     HKey
	Signature     HKey
	Authorization HKey
}{
	0,
	1,
	2,
	3,
	4,
	5,
	6,
	7,
	8,
}

type HeaderKeyBitMap struct {
	BitMap int // 0x01: request 0x10 response
	Key    string
}

var x_header_key_bitmap []*HeaderKeyBitMap = []*HeaderKeyBitMap{
	&HeaderKeyBitMap{BitMap: 1, Key: "x-ots-accesskeyid"},
	&HeaderKeyBitMap{BitMap: 1, Key: "x-ots-apiversion"},
	&HeaderKeyBitMap{BitMap: 3, Key: "x-ots-contentmd5"},
	&HeaderKeyBitMap{BitMap: 2, Key: "x-ots-contenttype"},
	&HeaderKeyBitMap{BitMap: 3, Key: "x-ots-date"},
	&HeaderKeyBitMap{BitMap: 1, Key: "x-ots-instancename"},
	&HeaderKeyBitMap{BitMap: 2, Key: "x-ots-requestid"},
	&HeaderKeyBitMap{BitMap: 0, Key: "x-ots-signature"},
	&HeaderKeyBitMap{BitMap: 0, Key: "Authorization"},
}

type CanonicalHeaders struct {
	headers []string
	length  int
}

func NewCanonicalHeaders() *CanonicalHeaders {
	chs := new(CanonicalHeaders)
	chs.length = len(x_header_key_bitmap)
	chs.headers = make([]string, chs.length)

	return chs
}

func (this *CanonicalHeaders) AddHeader(key HKey, value string) *CanonicalHeaders {

	pos := key.Value()
	if pos >= this.length {
		return this
	}

	this.headers[pos] = strings.TrimSpace(value)

	return this
}

func (this *CanonicalHeaders) GetHeader(key HKey) (val string) {

	pos := key.Value()
	if pos >= this.length {
		return
	}

	return this.headers[pos]
}

func (this *CanonicalHeaders) RequestString() (str string) {

	for pos, value := range this.headers {
		kb := x_header_key_bitmap[pos]
		if kb.BitMap&0x1 == 0 {
			continue
		}
		str = str + kb.Key + ":" + value + "\n"
	}

	return str
}

func (this *CanonicalHeaders) ResponseString() (str string) {

	for pos, value := range this.headers {
		kb := x_header_key_bitmap[pos]
		if kb.BitMap&0x2 == 0 {
			continue
		}
		str = str + kb.Key + ":" + value + "\n"
	}

	return str
}
