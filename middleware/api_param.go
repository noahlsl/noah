package middleware

import (
	"bytes"
	"github.com/noahlsl/noah/tools/strx"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/goccy/go-json"
)

func ParamMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var param string
		switch r.Method {
		case "GET":
			val := strings.Split(r.URL.RawQuery, "&")
			m := map[string]interface{}{}
			for _, v := range val {
				v1 := strings.Split(v, "=")
				if len(v1) != 2 {
					continue
				}
				m[v1[0]] = v1[1]
			}
			if len(m) == 0 {
				break
			}
			b, _ := json.Marshal(m)
			param = strx.B2s(b)

		case "POST", "PUT", "DELETE":
			contentType := r.Header.Get("Content-Type")
			// 从原有Request.Body读取
			b, _ := ioutil.ReadAll(r.Body)
			// 新建缓冲区并替换原有Request.body
			r.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			if strings.Contains(contentType, "application/json") {
				param = strx.B2s(b)
			} else if strings.Contains(contentType, "application/x-www-form-urlencoded") {
				val := strings.Split(strx.B2s(b), "&")
				m := map[string]interface{}{}
				for _, v := range val {
					v1 := strings.Split(v, "=")
					if len(v1) != 2 {
						continue
					}
					m[v1[0]] = v1[1]
				}
				if len(m) == 0 {
					break
				}
				b, _ = json.Marshal(m)
				param = strx.B2s(b)
			} else {
				// 不支持的类型
			}
		}

		r.Header.Set("param", param)
		next(w, r)
	}
}

func OriginalParamMiddleware(w http.ResponseWriter, r *http.Request) error {
	var param string
	switch r.Method {
	case "GET":
		val := strings.Split(r.URL.RawQuery, "&")
		m := map[string]interface{}{}
		for _, v := range val {
			v1 := strings.Split(v, "=")
			if len(v1) != 2 {
				continue
			}
			m[v1[0]] = v1[1]
		}
		if len(m) == 0 {
			break
		}
		b, _ := json.Marshal(m)
		param = strx.B2s(b)

	case "POST", "PUT", "DELETE":
		contentType := r.Header.Get("Content-Type")
		// 从原有Request.Body读取
		b, _ := ioutil.ReadAll(r.Body)
		// 新建缓冲区并替换原有Request.body
		r.Body = ioutil.NopCloser(bytes.NewBuffer(b))
		if strings.Contains(contentType, "application/json") {
			param = strx.B2s(b)
		} else if strings.Contains(contentType, "application/x-www-form-urlencoded") {
			val := strings.Split(strx.B2s(b), "&")
			m := map[string]interface{}{}
			for _, v := range val {
				v1 := strings.Split(v, "=")
				if len(v1) != 2 {
					continue
				}
				m[v1[0]] = v1[1]
			}
			if len(m) == 0 {
				break
			}
			b, _ = json.Marshal(m)
			param = strx.B2s(b)
		} else {
			// 不支持的类型
		}
	}

	r.Header.Set("param", param)

	return nil
}
