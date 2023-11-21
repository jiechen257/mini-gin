package gin

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type HandlerFunc func(*Context)
type HandlersChain []HandlerFunc

type Context struct {
	Request *http.Request
	Writer  http.ResponseWriter
}

func (c *Context) String(format string, data ...interface{}) {
	fmt.Fprintf(c.Writer, format, data...)
	return
}

type Engine struct {
	pool   sync.Pool
	router map[string]map[string]HandlersChain
}

func New() *Engine {

	engine := &Engine{}
	engine.pool.New = func() interface{} {
		return engine.allocateContext()
	}
	engine.router = make(map[string]map[string]HandlersChain)
	return engine
}

func (engine *Engine) allocateContext() *Context {
	return &Context{}
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := engine.pool.Get().(*Context)

	c.Writer = w
	c.Request = req

	engine.handleHTTPRequest(c)

	engine.pool.Put(c)
}

func (engine *Engine) handleHTTPRequest(c *Context) {
	httpMethod := c.Request.Method
	rPath := c.Request.URL.Path

	routers, ok := engine.router[httpMethod]
	if ok {
		handles, ok := routers[rPath]
		if ok {
			for _, handle := range handles {
				handle(c)
			}
			return
		}
	}
	c.String("%s", httpMethod+" "+rPath+" doesn't exist")
	return
}

func (engine *Engine) AddRoute(method, path string, handlers ...HandlerFunc) {

	_, ok := engine.router[method]
	if !ok {
		engine.router[method] = make(map[string]HandlersChain)
	}
	//2.判断该路径是否存在，如果不存在则插入，如果存在，则不处理
	_, ok = engine.router[method][path]
	if !ok {
		engine.router[method][path] = handlers
	}
}

func (engine *Engine) Run(address string) {
	err := http.ListenAndServe(address, engine)
	if err != nil {
		log.Fatal(err)
	}

}
