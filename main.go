package main

import (
	"awesomeProject2/ppnet"
)

func main() {

	e := ppnet.Default()

	e.GET("/xuan/*file", func(c *ppnet.Context) {

		c.String(200, c.Params["file"])

	})

	e.Run(":80")

}
