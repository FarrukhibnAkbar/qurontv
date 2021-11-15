package controllers

import(
	"log"

	model "backend/modules/pathway/models"

	"github.com/gin-gonic/gin"
)

func Delete(ctx *gin.Context){

	var deletePathway model.Pathway

	if err := ctx.BindJSON(&deletePathway);err != nil{
		log.Fatalf("failed pathway controllers delete %v", err)
	}

	model.Delete(deletePathway)
}