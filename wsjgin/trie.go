package wsjgin

import (
	"fmt"
	"strings"
)

// 只支持两种动态模糊匹配：
// 1. /:name  2. /*filepath

type trieNode struct {
	pattern  string
	part     string
	children []*trieNode
	isFuzzy  bool // 模糊匹配时:True
}

// 返回第一个匹配成功的节点,要么是一个精确的要么是一个模糊匹配
func (node *trieNode) matchChild(part string) *trieNode {
	for _, child := range node.children {
		// 精确匹配或模糊匹配
		if child.part == part || child.isFuzzy {
			return child
		}
	}
	return nil
}

// 返回所有匹配成功的节点,包含一个精确匹配的和一个模糊匹配的
func (node *trieNode) matchChildren(part string) []*trieNode {
	nodes := make([]*trieNode, 0)
	for _, child := range node.children {
		if child.part == part || child.isFuzzy {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// 插入函数：用于(在当前节点下)注册动态路由
// 参数为：模式串（输入的待匹配路由）,此pattern的组成部分集合(由pattern串解析得来)，路由层号
// height从0开始，代表层数
func (node *trieNode) insert(pattern string, parts []string, height int) {
	// 当匹配到最后一层时，才给当前节点pattern属性赋值，并且返回
	if len(parts) == height {
		node.pattern = pattern
		return
	}
	part := parts[height]
	child := node.matchChild(part)
	if child == nil {
		// 此时pattern为空
		child = &trieNode{
			part:    part,
			isFuzzy: part[0] == ':' || part[0] == '*',
		}
		node.children = append(node.children, child)
	}
	child.insert(pattern, parts, height+1)
}

// 搜索函数: 用于访问路由
func (node *trieNode) search(parts []string, height int) *trieNode {
	if len(parts) == height || strings.HasPrefix(node.part, "*") {
		if node.pattern == "" {
			return nil
		}
		return node
	}
	part := parts[height]
	children := node.matchChildren(part)
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}

// 遍历当前节点作为根节点的子树
func (node *trieNode) travel(nodes *([]*trieNode)) {
	if node.pattern != "" {
		*nodes = append(*nodes, node)
	}
	for _, child := range node.children {
		child.travel(nodes)
	}
}

// 打印节点内容
func (node *trieNode) ToString() string {
	return fmt.Sprintf("node{pattern=%s, part=%s, isFuzzy=%t}", node.pattern, node.part, node.isFuzzy)
}
