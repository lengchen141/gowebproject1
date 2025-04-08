package main

import (
	"gowebproject1/common"
	"gowebproject1/router"
)

func main() {

	common.InitDb()

	route := router.Router()

	route.Run(":8081")
}
