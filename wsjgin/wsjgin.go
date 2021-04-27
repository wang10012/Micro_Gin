package wsjgin

import (
	"net/http"
)

type HandleFunc func(c *Context)

// 构建RouterMap表 key:路径 value:handlefunc
type RouterMap struct {
	router *router
}

// 为RouterMap表添加路由和对应函数
func (routerMap *RouterMap) addRouter(method string, pattern string, handler HandleFunc) {
	routerMap.router.addRouter(method, pattern, handler)
}

// 调用addRouter()定义"GET"请求
func (routerMap *RouterMap) GET(pattern string, handler HandleFunc) {
	routerMap.addRouter("GET", pattern, handler)
}

// 调用addRouter()定义"POST"请求
func (routerMap *RouterMap) POST(pattern string, handler HandleFunc) {
	routerMap.addRouter("POST", pattern, handler)
}

// 构造函数
func Default() *RouterMap {
	return &RouterMap{
		router: NewRouter(),
	}
}

// 解析请求，查找RouterMap
func (routerMap *RouterMap) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := NewContext(w, r)
	routerMap.router.handleFunc(c)
}

// 启动服务
func (routerMap *RouterMap) Run(address string) (err error) {
	return http.ListenAndServe(address, routerMap)
}
