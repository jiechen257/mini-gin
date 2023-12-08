## 动态路由
动态路由包含 `/user/:id` 和 `/user/?filepath` 两种格式

针对这两种格式的算法设计:(len为串的长度)
- 使用 HashMap 遍历两次进行匹配 - 构造完 map 后进行查询，构造的时间复杂度 O(n * len), 查询则是 O(n) * O(1) = O(n)
- 使用 `trie` 树 - 边构造树边查询，时间复杂度直接是 O(len)

`trie` 树简化版示例：
```golang
// TrieNode 表示字典树的节点
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

// Trie 表示字典树
type Trie struct {
	root *TrieNode
}

// NewTrie 创建一个新的字典树
func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
			isEnd:    false,
		},
	}
}

// Insert 向字典树中插入一个单词
func (t *Trie) Insert(word string) {
	node := t.root
	for _, char := range word {
		if node.children == nil {
			node.children = make(map[rune]*TrieNode)
		}
		if _, ok := node.children[char]; !ok {
			node.children[char] = &TrieNode{
				children: make(map[rune]*TrieNode),
				isEnd:    false,
			}
		}
		node = node.children[char]
	}
	node.isEnd = true
}

// Search 在字典树中搜索一个单词
func (t *Trie) Search(word string) bool {
	node := t.root
	for _, char := range word {
		if _, ok := node.children[char]; !ok {
			return false
		}
		node = node.children[char]
	}
	return node.isEnd
}
```

## 中间件
数据结构还是用的洋葱模型，本质上就是通过闭包的方式把 `Context` 传递下去

洋葱模型简化版示例：
```golang
// MiddlewareFunc 定义中间件函数类型
type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// Middleware1 是第一个中间件
func Middleware1(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Entering Middleware1")
		next(w, r)
		fmt.Println("Leaving Middleware1")
	}
}

// Middleware2 是第二个中间件
func Middleware2(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Entering Middleware2")
		next(w, r)
		fmt.Println("Leaving Middleware2")
	}
}

// FinalHandler 是最终的处理函数
func FinalHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Executing FinalHandler")
	w.Write([]byte("Hello, World!\n"))
}

func main() {
	// 创建最终处理函数
	finalHandler := http.HandlerFunc(FinalHandler)
	// 构建洋葱模型
	handler := Middleware1(Middleware2(finalHandler))

	// 注册处理函数
	http.Handle("/", handler)
	// 启动 HTTP 服务器
	http.ListenAndServe(":8080", nil)
}
```

## 错误恢复
> 关于 golang 的错误处理，详细可以看这篇文章 [GO 处理错误优雅化](https://becase.top/post/20231105000000)

利用中间件 + `recover` 机制实现错误恢复

```golang
func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", trace(message))
				c.Fail(http.StatusInternalServerError, "Internal Server Error")
			}
		}()

		c.Next()
	}
}
```