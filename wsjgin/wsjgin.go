package wsjgin

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strings"
)

type HandleFunc func(c *Context)

type (
	RouterGroup struct {
		prefix      string
		middlewares []HandleFunc
		// 所有的group共享一个routermap实例
		routermap *RouterMap
	}
	// 构建RouterMap表
	RouterMap struct {
		*RouterGroup
		router        *router
		groups        []*RouterGroup
		htmlTemplates *template.Template // for html render
		funcMap       template.FuncMap   // for html render
	}
)

// 构造函数
func Default() *RouterMap {
	routermap := &RouterMap{
		router: NewRouter(),
	}
	routermap.RouterGroup = &RouterGroup{
		routermap: routermap,
	}
	routermap.groups = []*RouterGroup{
		routermap.RouterGroup,
	}
	return routermap
}

// 设置自定义渲染函数
func (routermap *RouterMap) SetFuncMap(funcMap template.FuncMap) {
	routermap.funcMap = funcMap
}

// 加载模板
func (routermap *RouterMap) LoadHTMLGlob(pattern string) {
	routermap.htmlTemplates = template.Must(template.New("").Funcs(routermap.funcMap).ParseGlob(pattern))
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	routermap := group.routermap
	newGroup := &RouterGroup{
		prefix:    group.prefix + prefix,
		routermap: routermap,
	}
	routermap.groups = append(routermap.groups, newGroup)
	return newGroup
}

// 为group添加中间件
func (group *RouterGroup) Use(middlewares ...HandleFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

// 为RouterMap表添加路由和对应函数
func (group *RouterGroup) addRouter(method string, subpattern string, handler HandleFunc) {
	pattern := group.prefix + subpattern
	log.Printf("Router %4s - %s", method, pattern)
	group.routermap.router.addRouter(method, pattern, handler)
}

// 调用addRouter()定义"GET"请求
func (group *RouterGroup) GET(pattern string, handler HandleFunc) {
	group.addRouter("GET", pattern, handler)
}

// 调用addRouter()定义"POST"请求
func (group *RouterGroup) POST(pattern string, handler HandleFunc) {
	group.addRouter("POST", pattern, handler)
}

// 解析请求，查找RouterMap
func (routerMap *RouterMap) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var middlewares []HandleFunc
	for _, group := range routerMap.groups {
		if strings.HasPrefix(r.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := NewContext(w, r)
	// 中间件加入context的函数集合
	c.handlers = middlewares
	c.routermap = routerMap
	routerMap.router.handleFunc(c)
}

// 启动服务
func (routerMap *RouterMap) Run(address string) (err error) {
	return http.ListenAndServe(address, routerMap)
}

// 创建静态函数
func (group *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandleFunc {
	absolutePath := path.Join(group.prefix, relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(c *Context) {
		file := c.GetParameter("filepath")
		// 检查此文件是否存在
		if _, err := fs.Open(file); err != nil {
			c.SetStatusCdoe(http.StatusNotFound)
			return
		}
		fileServer.ServeHTTP(c.Writer, c.Request)
	}
}

// 后端路由和文件绝对地址映射，注册路由
func (group *RouterGroup) Static(relativePath string, root string) {
	handler := group.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")
	// 注册GET函数
	group.GET(urlPattern, handler)
}
