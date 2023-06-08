package gee

import (
    "fmt"
    "strings"
)

// /:lang/doc
// /:lang/tutorial
// /:lang/intro
// /p/blog
// /p/related/*filepath

type node struct {
    pattern  string  // 待匹配路由，例如 /p/:lang
    part     string  // 路由中的一部分，例如 :lang
    children []*node // 子节点，例如 [doc, tutorial, intro]
    isWild   bool    // 是否模糊匹配，part 中含有 : 或 * 时为 true
}

func (n *node) String() string {
    return fmt.Sprintf("node{pattern=%s, part=%s, isWild=%t}", n.pattern, n.part, n.isWild)
}

// 第一个匹配成功的节点，用于插入
func (n *node) matchChild(part string) *node {
    for _, child := range n.children {
        if child.part == part || child.isWild { // 精确匹配成功 或者 当前节点为 :lang, 模糊匹配
            return child
        }
    }
    return nil
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
    // /p/:lang/doc 只有在第三层节点，即 doc 节点，pattern才会设置为 /p/:lang/doc
    // 而 /p 与 /p/:lang 是没有对应 handler
    // 因为切片下标与 height 都是从 0 开始的，当 trie 完成创建时，height = len(parts) - 1，再次递归调用 height + 1
    if len(parts) == height { // 对应路由已完成 trie 创建
        n.pattern = pattern
        return
    }

    part := parts[height] // 一般从 height = 0 开始，递归进行匹配查找
    child := n.matchChild(part)
    if child == nil { // 无对应节点，新建并加入 n 的子节点中
        child = &node{
            part:   part,
            isWild: part[0] == ':' || part[0] == '*',
        }
        n.children = append(n.children, child)
    }
    child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
    if len(parts) == height || strings.HasPrefix(n.part, "*") {
        // 由于 isWild 为 true 时，存在 : 和 * 两种情况
        // 当前节点若为 *，直接匹配成功，只需再次判断是否有对应 handler 即可
        if n.pattern == "" { // 当前节点 pattern 为空，即该节点路径对应的路由没有 handler
            return nil
        }
        return n
    }

    part := parts[height]
    children := n.matchChildren(part)

    for _, child := range children { // 一一查找对应的路由，返回该节点
        result := child.search(parts, height+1)
        if result != nil {
            return result
        }
    }
    return nil
}

func (n *node) travel(list *([]*node)) {
    if n.pattern != "" { // 有 handler 才 append
        *list = append(*list, n)
    }
    for _, child := range n.children {
        child.travel(list)
    }
}
