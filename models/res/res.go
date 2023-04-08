package res

import (
	"github.com/noahlsl/noah/tools/strx"
	"time"

	"github.com/goccy/go-json"
	"github.com/golang-module/dongle"
)

type Res struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
	Trace string      `json:"trace"`
	Ts    int64       `json:"ts"`
}

func NewRes() *Res {
	return &Res{
		Code: 0,
		Msg:  "succeed",
		Ts:   time.Now().UnixMicro(),
	}
}

func (r *Res) WithCode(code int) *Res {
	r.Code = code
	return r
}

func (r *Res) WithMsg(msg string) *Res {
	r.Msg = msg
	return r
}

func (r *Res) WithData(data interface{}) *Res {
	r.Data = data
	return r
}

func (r *Res) WithTrace(trace string) *Res {
	r.Trace = trace
	return r
}

func (r *Res) ToStr() string {
	marshal, _ := json.Marshal(r)
	return strx.B2s(marshal)
}

func (r *Res) ToBytes() []byte {
	marshal, _ := json.Marshal(r)
	return marshal
}

// ToAesStr Aes秘钥加密输出字符串
func (r *Res) ToAesStr(cipher *dongle.Cipher) string {

	marshal, _ := json.Marshal(r)
	data := dongle.Encrypt.FromBytes(marshal).ByAes(cipher).ToRawString()
	return data
}

// ToAesBytes Aes秘钥加密输出Bytes
func (r *Res) ToAesBytes(cipher *dongle.Cipher) []byte {

	marshal, _ := json.Marshal(r)
	data := dongle.Encrypt.FromBytes(marshal).ByAes(cipher).ToRawBytes()
	return data
}

// ToRsaStr Rsa私钥加密输出字符串
func (r *Res) ToRsaStr(privateKey string) string {

	marshal, _ := json.Marshal(r)
	data := dongle.Encrypt.FromBytes(marshal).ByRsa(privateKey).ToRawString()
	return data
}

// ToRsaBytes Rsa私钥加密输出Bytes
func (r *Res) ToRsaBytes(privateKey string) []byte {

	marshal, _ := json.Marshal(r)
	data := dongle.Encrypt.FromBytes(marshal).ByRsa(privateKey).ToRawBytes()
	return data
}
