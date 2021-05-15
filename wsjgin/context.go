package wsjgin

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 别名
type H map[string]interface{}

type Context struct {
	// 响应
	Writer     http.ResponseWriter
	StatusCode int
	// 请求
	Request *http.Request
	Path    string
	Method  string
	// 新增参数集合
	Parameters map[string]string
	// 新增中间件有关属性
	// 中间件函数和用户定义的handler都会储存在内
	handlers []HandleFunc
	index    int
}

// 构造函数
func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer:  w,
		Request: r,
		Path:    r.URL.Path,
		Method:  r.Method,
		index:   -1,
	}
}

func (c *Context) Next() {
	// 在next函数中才调用函数
	c.index++
	length := len(c.handlers)
	for ; c.index < length; c.index++ {
		c.handlers[c.index](c)
	}
}

// fail函数
func (c *Context) Fail(code int, err string) {
	c.index = len(c.handlers)
	c.JSON(code, H{"message": err})
}

func (c *Context) GetParameter(key string) string {
	value, _ := c.Parameters[key]
	return value
}

// 状态码设置
func (c *Context) SetStatusCdoe(statusCode int) {
	c.StatusCode = statusCode
	c.Writer.WriteHeader(statusCode)
}

// 响应头设置
func (c *Context) SetHeader(contentType string, value string) {
	c.Writer.Header().Set(contentType, value)
}

// query方法，查询query=?参数
func (c *Context) Query(queryValue string) string {
	return c.Request.URL.Query().Get(queryValue)
}

// PostForm方法,表单提交
func (c *Context) PostForm(formValue string) string {
	return c.Request.FormValue(formValue)
}

// 以下提供不同类型数据的构造方法
// JSON/HTML/String/Data响应

// 构造JSON响应
func (c *Context) JSON(statusCode int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.SetStatusCdoe(statusCode)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// 构造HTML响应
func (c *Context) HTML(statusCode int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.SetStatusCdoe(statusCode)
	c.Writer.Write([]byte(html))
}

// 构造String响应
func (c *Context) String(statusCode int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.SetStatusCdoe(statusCode)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// 构造Data响应
func (c *Context) Data(statusCode int, data []byte) {
	c.SetStatusCdoe(statusCode)
	c.Writer.Write(data)
}
