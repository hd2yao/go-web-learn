package gee

import (
    "fmt"
    "strings"
)

type node struct {
    pattern  string  // 待匹配路由 /p/:lang/doc 只有在第三层节点, 即 doc 节点, pattern 才会设置为 /p/:lang/doc, p 和 :lang 节点的 pattern 属性皆为空
    part     string  // 路由中的一部分 p :lang doc
    children []*node // 子节点
    isWild   bool    // 表示该节点是否有通配符子节点（即带有 : 或 * 的路径参数节点）
}

func (n *node) String() string {
    return fmt.Sprintf("node{pattern=%s, part=%s, isWild=%t}", n.pattern, n.part, n.isWild)
}

// 第一个匹配成功的节点，用于插入
func (n *node) matchChild(part string) *node {
    var wildChild *node
    for _, child := range n.children {
        // 这里查找不够严谨，可能会发生精准匹配和通配符匹配都匹配成功的结果
        // 可以修改为：优先精准匹配
        //if child.part == part || child.isWild {
        //    return child
        //}
        if child.part == part {
            return child
        }
        if child.isWild {
            wildChild = child
        }
    }
    return wildChild
}

// 所有匹配成功的节点，用于查找
func (n *node) matchChildren(part string) []*node {
    nodes := make([]*node, 0)
    for _, child := range n.children {
        if child.part == part || child.isWild {
            nodes = append(nodes, child)
        }
    }
    return nodes
}

func (n *node) insert(pattern string, parts []string, height int) {
    if len(parts) == height {
        n.pattern = pattern
        return
    }

    part := parts[height]
    child := n.matchChild(part)
    if child == nil {
        child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
        n.children = append(n.children, child)
    }
    child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
    if len(parts) == height || strings.HasPrefix(n.part, "*") {
        if n.pattern == "" {
            return nil
        }
        return n
    }

    part := parts[height]
    children := n.matchChildren(part)

    for _, child := range children {
        result := child.search(parts, height+1)
        if result != nil {
            return result
        }
    }

    return nil
}

func (n *node) travel(list *([]*node)) {
    if n.pattern != "" {
        *list = append(*list, n)
    }
    for _, child := range n.children {
        child.travel(list)
    }
}
