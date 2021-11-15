package main

import(
	"fmt"

	ctrl "backend/modules/pathway/controllers"

	"github.com/gin-gonic/gin"
)

func main(){
	
	r := gin.Default()

	//PATHWAY
	pathways := r.Group("/pathways")
	pathways.GET("/", ctrl.Get)
	pathways.POST("/", ctrl.Post)	
	pathways.PUT("/", ctrl.Update)

	fmt.Println(":8080")

	r.Run()

}