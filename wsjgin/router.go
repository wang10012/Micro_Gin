package wsjgin

import (
	"log"
	"net/http"
)

type router struct {
	handlers map[string]HandleFunc
}

// 构造函数
func NewRouter() *router {
	return &router{handlers: make(map[string]HandleFunc)}
}

// 添加路由
func (r *router) addRouter(method string, pattern string, handler HandleFunc) {
	log.Printf("Router: %4s - %s", method, pattern) // 打印日志
	key := method + "-" + pattern
	r.handlers[key] = handler
}

// 处理路由请求
func (r *router) handleFunc(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
