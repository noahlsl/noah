package resultx

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-module/dongle"
	"github.com/noahlsl/noah/models/res"
	"github.com/noahlsl/noah/tools/errx"
	"github.com/noahlsl/noah/tools/structx"
	"github.com/pkg/errors"
)

type Reply struct {
	error errx.Err
}

var (
	reply  *Reply
	cipher *dongle.Cipher
)

func init() {
	reply = newReply(NewErrManger())
}

func SetReply(e errx.Err) {
	reply = newReply(e)
}

func newReply(e errx.Err) *Reply {
	return &Reply{
		error: e,
	}
}

func WithCipher(c *dongle.Cipher) {
	cipher = c
}

func Result(c *gin.Context, data interface{}) {

	rs := res.NewRes().
		WithData(data).WithTrace(c.GetHeader("trace"))
	c.String(http.StatusOK, structx.StructToStr(rs))
}

func ResultAes(c *gin.Context, data interface{}) {

	rs := res.NewRes().
		WithData(data).WithTrace(c.GetHeader("trace")).ToAesBytes(cipher)
	c.String(http.StatusOK, structx.StructToStr(rs))
}

func ResultErr(c *gin.Context, code int, e error) {

	lang := c.GetHeader("lang")
	rs := res.NewRes().
		WithCode(code).
		WithMsg(reply.error.GetErr(code, lang).Error()).
		WithTrace(c.GetHeader("trace"))

	err := errors.Cause(e)
	c1 := reply.error.GetCode(err, lang)
	rs.WithCode(c1)
	if c.GetHeader("debug") != "" {
		rs.WithMsg(e.Error())
	} else {
		er := reply.error.GetErr(c1, lang)
		msg := er.Error()
		rs.WithMsg(msg)
	}

	c.String(http.StatusOK, structx.StructToStr(rs))
}
