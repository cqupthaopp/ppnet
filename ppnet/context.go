package ppnet

import "C"
import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	Res  http.ResponseWriter
	Req  *http.Request
	Path string

	Params     map[string]string
	Method     string
	StatusCode int

	middlewares []HFunc
	runningIdx  int
}

func newContext(res http.ResponseWriter, req *http.Request) *Context {

	return &Context{
		Res:         res,
		Req:         req,
		Path:        req.URL.Path,
		Params:      make(map[string]string),
		Method:      req.Method,
		StatusCode:  http.StatusOK,
		middlewares: make([]HFunc, 0, 16),
		runningIdx:  -1,
	}

}

func (c *Context) Next() {
	if c.runningIdx == len(c.middlewares)-1 {
		return
	}
	c.runningIdx++
	c.middlewares[c.runningIdx](c)

	c.Next()

}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) Get(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) SetHeader(key, value string) {
	c.Res.Header().Set(key, value)
}

func (c *Context) SetStatusCode(code int) {
	c.StatusCode = code
}

func (c *Context) JSON(code int, data H) {
	c.SetHeader("Content-Type", "application/json")
	c.SetStatusCode(code)
	encode := json.NewEncoder(c.Res)
	if err := encode.Encode(data); err != nil {
		http.Error(c.Res, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.SetStatusCode(code)
	c.Res.Write([]byte(html))
}

func (c *Context) String(code int, format string, vales ...any) {
	c.SetHeader("Content-Type", "text/plain")
	c.SetStatusCode(code)
	c.Res.Write([]byte(fmt.Sprintf(format, vales...)))
}

func (c *Context) Data(code int, data []byte) {
	c.SetStatusCode(code)
	c.Res.Write(data)
}
