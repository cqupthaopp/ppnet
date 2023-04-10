package ppnet

import (
	"log"
	"net/http"
	"strings"
)

type engine struct {
	*groupRouter
	routers []*groupRouter
	router  *router
}

func newEngine() *engine {

	e := engine{}
	e.groupRouter = newGroupRouter("", nil, &e)
	e.routers = make([]*groupRouter, 0, 16)
	e.routers = append(e.routers, e.groupRouter)
	e.router = newRouter()

	return &e
}

func New() *engine {
	return newEngine()
}

func Default() *engine {
	e := newEngine()
	e.Use(Recover)
	//e.Use(LogMiddleWare)
	return e
}

func (e *engine) Run(addr string) {
	log.Fatalf(http.ListenAndServe(addr, e).Error())
}

func (e *groupRouter) GET(pattern string, handlerFunc HFunc) {
	if err := e.engine.router.addHandlerFunc("GET", strings.Split(e.prefix+pattern, "/"), handlerFunc); err == nil {
		log.Println("Method : ", "GET", "   --   ", "path : ", e.prefix+pattern)
	}
}

func (e *groupRouter) POST(pattern string, handlerFunc HFunc) {
	if err := e.engine.router.addHandlerFunc("POST", strings.Split(e.prefix+pattern, "/"), handlerFunc); err == nil {
		log.Println("Method : ", "GET", "   --   ", "path : ", e.prefix+pattern)
	}
}

func (e *groupRouter) PUT(pattern string, handlerFunc HFunc) {
	if err := e.engine.router.addHandlerFunc("PUT", strings.Split(e.prefix+pattern, "/"), handlerFunc); err == nil {
		log.Println("Method : ", "GET", "   --   ", "path : ", e.prefix+pattern)
	}
}

func (e *groupRouter) DELETE(pattern string, handlerFunc HFunc) {
	if err := e.engine.router.addHandlerFunc("DELETE", strings.Split(e.prefix+pattern, "/"), handlerFunc); err == nil {
		log.Println("Method : ", "GET", "   --   ", "path : ", e.prefix+pattern)
	}
}

func (e *groupRouter) PATCH(pattern string, handlerFunc HFunc) {
	if err := e.engine.router.addHandlerFunc("PATCH", strings.Split(e.prefix+pattern, "/"), handlerFunc); err == nil {
		log.Println("Method : ", "GET", "   --   ", "path : ", e.prefix+pattern)
	}
}
