package main

import "github.com/gin-gonic/gin"

func main()  {
	router := gin.Default()

	//userGroup =  router.Group("users")
	router.Run(":3000")
}

