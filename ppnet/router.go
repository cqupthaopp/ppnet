package ppnet

import (
	"fmt"
	"net/http"
	"strings"
)

type HFunc func(ctx *Context)

type router struct {
	handlers map[string]*routerPart
}

func newRouter() *router {

	r := router{make(map[string]*routerPart)}

	for _, v := range []string{"GET", "POST", "DELETE", "PATCH", "PUT"} {
		r.handlers[v] = newFatherRouter("", false)
	}

	return &r
}

func (r *router) addHandlerFunc(method string, pattern []string, handlerFunc HFunc) error {
	return r.handlers[method].addRouterInTrie(pattern, 1, handlerFunc)
}

//实现 http.Handler 接口
func (e *engine) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	if routerPart, exist := e.router.handlers[req.Method]; exist {
		ctext := newContext(res, req)
		if hFunc, err := routerPart.parseRouterInTrie(ctext, strings.Split(req.URL.Path, "/"), 1); err == nil {
			for _, groups := range e.routers {
				if strings.HasPrefix(req.URL.Path, groups.prefix) {
					ctext.middlewares = append(ctext.middlewares, groups.middlewares...)
				}
			}
			ctext.middlewares = append(ctext.middlewares, hFunc)
			ctext.Next()
			return
		}
	}
	fmt.Fprintf(res, "寄咯，找不到页面！ %v", req.URL)

}
