package gee

import "strings"

type node struct {
	pattern  string  // 待匹配路由，是否一个完整的 url，不是则为空字符串
	part     string  // URL 块值，用 / 分割的部分，比如 /abc/123，abc 和 123 就是两个 part
	children []*node // 子节点
	isWild   bool    // 是否精确匹配，以实现动态路由匹配 part 中含有 : 或 * 时为 true
}

// 找到匹配的子节点，场景是用在插入时使用，找到1个匹配的就立即返回
func (n *node) matchChild(part string) *node {
	// 遍历当前节点的所有子节点，看是否能找到匹配的子节点，将其返回
	for _, child := range n.children {
		// 如果有模糊匹配的也会成功匹配上
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 返回所有匹配的子节点，用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// 一边匹配一边插入的方法
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		// 如果已经匹配结束，那么将 pattern 赋值给该 node，表示它是一个完整的 url
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		// 没有匹配上，那么就新建节点，作为当前节点的子节点
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	// 如果上面经过了 if 那么此时 child 为上面新建的节点
	// 接着插入后续 part 的节点
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	// 递归终止条件：找到末尾或者通配符
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		// pattern 为空表示它不是一个完整的 url，匹配失败
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	// 获取所有可能得子路径
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
