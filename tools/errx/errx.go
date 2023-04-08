package errx

type Err interface {
	GetCode(err error, lang string) int
	GetErr(code int, lang string) error
}
