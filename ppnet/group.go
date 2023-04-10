package ppnet

type groupRouter struct {
	prefix      string
	parent      *groupRouter
	engine      *engine
	middlewares []HFunc
}

func newGroupRouter(pre string, parent *groupRouter, engine *engine) *groupRouter {

	return &groupRouter{
		prefix:      pre,
		parent:      parent,
		engine:      engine,
		middlewares: make([]HFunc, 0, 16),
	}
}

func (e *groupRouter) Group(prefix string) *groupRouter {
	r := newGroupRouter(e.prefix+prefix, e, e.engine)
	e.engine.routers = append(e.engine.routers, r)
	return r
}

func (e *groupRouter) Use(middleware HFunc) {
	e.middlewares = append(e.middlewares, middleware)
}
