package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}
type Context struct {
	Writer     http.ResponseWriter
	request    *http.Request //请求
	Path       string
	Method     string //写入
	Statuscode int
}

func Newcontext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer:  w,
		request: req,
		Path:    req.URL.Path,
		Method:  req.Method,
	}
}
func (c *Context) PostForm(key string) string {
	return c.request.FormValue(key)
}
func (c *Context) Query(key string) string {
	return c.request.URL.Query().Get(key)
}
func (c *Context) Status(code int) {
	c.Statuscode = code
	c.Writer.WriteHeader(code)
}
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("context type", "text/explain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))

}
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("context type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}
func (c *Context) HTML(code int, html string) {
	c.Status(code)
	c.Writer.Write([]byte(html))
}
