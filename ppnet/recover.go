package ppnet

import (
	"fmt"
	"log"
	"net/http"
)

func Recover(c *Context) {

	defer func() {
		if val := recover(); val != nil {
			log.Println(fmt.Sprintf("Get Error: %v", val))
			c.String(http.StatusInternalServerError, fmt.Sprintf("Get Error: %v", val))
			return
		}
	}()

	c.Next()

}
