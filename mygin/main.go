package main

import (
	api "mygin/router"
)

// @title gin demo API
// @version 1.0
// @description This is a sample gin demo.
// @termsOfService https://github.com/kyralo/go-gin-xorm
// @contact.email kyralo721@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

//@host 127.0.0.1:8888
// @BasePath /api/v2

func main() {
	r := api.SetupRouter()
	// Listen and Server in 0.0.0.0:8888

	r.Run(":8888")
}
