package wsjgin

import (
	"net/http"
	"strings"
)

// roots: key:请求方式(get/post) value:Trie Tree
// handlers: key: (example) GET-/a/b/c value:handlefunc
type router struct {
	roots    map[string]*trieNode
	handlers map[string]HandleFunc
}

// 构造函数
func NewRouter() *router {
	return &router{
		roots:    make(map[string]*trieNode),
		handlers: make(map[string]HandleFunc),
	}
}

// 解析pattern
func parsePattern(pattern string) []string {
	patternList := strings.Split(pattern, "/")
	parts := make([]string, 0)
	// 只能有一个“*”
	for _, item := range patternList {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

// 添加路由
func (r *router) addRouter(method string, pattern string, handler HandleFunc) {
	// last update
	//log.Printf("Router: %4s - %s", method, pattern) // 打印日志
	//key := method + "-" + pattern
	//r.handlers[key] = handler
	key := method + "-" + pattern
	parts := parsePattern(pattern)
	if _, ok := r.roots[method]; !ok {
		r.roots[method] = &trieNode{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

// 返回搜索到的节点和解析到的参数集合
// example：如 {filepath: "css/geektutu.css"}
func (r *router) getRouter(method string, path string) (*trieNode, map[string]string) {
	pathParts := parsePattern(path)
	parameters := make(map[string]string)
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}
	node := root.search(pathParts, 0)
	if node != nil {
		patternParts := parsePattern(node.pattern)
		for index, part := range patternParts {
			if part[0] == ':' {
				parameters[part[1:]] = pathParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				parameters[part[1:]] = strings.Join(pathParts[index:], "/")
			}
		}
		return node, parameters
	}
	return nil, nil
}

// 返回以当前节点作为root的子树节点集合
func (r *router) getRouters(method string) []*trieNode {
	nodes := make([]*trieNode, 0)
	root, ok := r.roots[method]
	if !ok {
		return nil
	}
	root.travel(&nodes)
	return nodes

}

// 处理路由请求
func (r *router) handleFunc(c *Context) {
	node, parameters := r.getRouter(c.Method, c.Path)
	if node != nil {
		c.Parameters = parameters
		key := c.Method + "-" + node.pattern
		// 将用户自定义路由函数加入context的handlers集合中
		c.handlers = append(c.handlers, r.handlers[key])
	} else {
		c.handlers = append(c.handlers, func(c *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		})
	}
	// 开始执行context中handlers中的函数
	c.Next()
}
