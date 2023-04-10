package ppnet

import "log"

func LogMiddleWare(c *Context) {

	log.Println("Mehod: %s   ::  URL: %s", c.Method, c.Req.URL.Path)

}
