package main

import (
	api "mygin/router"
)



func main() {
	r := api.SetupRouter()
	// Listen and Server in 0.0.0.0:8888
	r.Run(":8888")
}
