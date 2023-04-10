package ppnet

import (
	"errors"
)

type routerPart struct {
	part      string
	isWild    bool
	nextParts []*routerPart
	handler   HFunc
}

func newFatherRouter(pathPart string, isWild bool) *routerPart {
	r := new(routerPart)
	r.part = pathPart
	r.isWild = isWild
	r.nextParts = make([]*routerPart, 0, 16)
	r.handler = nil
	return r
}

func (rp *routerPart) addRouterInTrie(routerPath []string, idx int, handler HFunc) error {

	if idx == len(routerPath) {
		rp.handler = handler
		return nil
	}

	for _, nxtPart := range rp.nextParts {
		if nxtPart != nil && nxtPart.part == routerPath[idx] {
			return nxtPart.addRouterInTrie(routerPath, idx+1, handler)
		}
	}

	if len(routerPath[idx]) >= 1 && (routerPath[idx][0] == '*' || routerPath[idx][0] == ':') {
		if routerPath[idx][0] == '*' && idx != len(routerPath)-1 {
			return errors.New("Add function Error")
		}
		rp.nextParts = append(rp.nextParts, newFatherRouter(routerPath[idx], true))
	} else {
		rp.nextParts = append(rp.nextParts, newFatherRouter(routerPath[idx], false))
	}

	return (rp.nextParts[len(rp.nextParts)-1]).addRouterInTrie(routerPath, idx+1, handler)

}

func (rp *routerPart) parseRouterInTrie(c *Context, routerPath []string, idx int) (HFunc, error) {

	if idx == len(routerPath) {
		if rp.handler != nil {
			return rp.handler, nil
		}
		return nil, errors.New("No path found")
	}

	for _, v := range rp.nextParts {
		if v != nil && (v.isWild || v.part == routerPath[idx]) {

			if len(v.part) >= 1 && v.part[0] == '*' {

				now_path := ""
				for ; idx < len(routerPath); idx++ {
					now_path += routerPath[idx] + "/"
				}
				c.Params[v.part[1:]] = now_path[:len(now_path)-1]

				if v.handler == nil {
					return nil, errors.New("No path found")
				}
				return v.handler, nil
			}

			handler, err := v.parseRouterInTrie(c, routerPath, idx+1)
			if err == nil {
				if v.isWild {
					c.Params[v.part[1:]] = routerPath[idx]
				}
				return handler, err
			}
		}
	}

	return nil, errors.New("No path found")

}
