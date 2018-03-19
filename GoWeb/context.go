package main

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type Context struct {
	Params map[string]interface{}

	ResponseWriter http.ResponseWriter
	Request        *http.Request
}

type HandlerFunc func(*Context)

func (c *Context) RenderJson(v interface{}) {

	c.ResponseWriter.WriteHeader(http.StatusOK)

	c.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")

	if err := json.NewEncoder(c.ResponseWriter).Encode(v); err != nil {

		c.RenderErr(http.StatusInternalServerError, err)

	}

}

func (c *Context) RenderErr(code int, err error) {
	if err != nil {
		if code > 0 {
			// 정상적인 code를 전달하면 HTTP Status를 해당 code로 지정
			http.Error(c.ResponseWriter, http.StatusText(code), code)
		} else {
			// 정상적인 code가 아니면 HTTP Status를 StatusInternalServerError로 지정
			defaultErr := http.StatusInternalServerError
			http.Error(c.ResponseWriter, http.StatusText(defaultErr), defaultErr)
		}
	}
}

func (c *Context) RenderXml(v interface{}) {

	c.ResponseWriter.WriteHeader(http.StatusOK)
	c.ResponseWriter.Header().Set("Content-Type", "application/xml; charset=utf-8")

	if err := xml.NewEncoder(c.ResponseWriter).Encode(v); err != nil {

		c.RenderErr(http.StatusInternalServerError, err)

	}

}
