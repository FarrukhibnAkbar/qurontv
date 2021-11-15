package controllers

import(
	"log"

	model "backend/modules/pathway/models"

	"github.com/gin-gonic/gin"
)

func Update(ctx *gin.Context){

	var putPathway model.Pathway

	if err := ctx.BindJSON(&putPathway);err != nil {
		log.Fatalf("failed pathway controllers update %v", err)
	}

	pathway := model.Update(putPathway)

	ctx.JSON(202, pathway)
}