package wsjgin

import (
	"fmt"
	"net/http"
)

type HandleFunc func(http.ResponseWriter, *http.Request)

// 构建RouterMap表 key:路径 value:handlefunc
type RouterMap struct {
	router map[string]HandleFunc
}

// 为RouterMap表添加路由和对应函数
func (routerMap *RouterMap) addRouter(method string, pattern string, handler HandleFunc) {
	key := method + "-" + pattern
	routerMap.router[key] = handler
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
		router: make(map[string]HandleFunc),
	}
}

// 解析请求，查找RouterMap
func (routerMap *RouterMap) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if handler, ok := routerMap.router[key]; ok {
		handler(w, r)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL)
	}
}

// 启动服务
func (routerMap *RouterMap) Run(address string) (err error) {
	return http.ListenAndServe(address, routerMap)
}

