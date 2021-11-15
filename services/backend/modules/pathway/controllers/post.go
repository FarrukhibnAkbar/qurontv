package controllers

import(
	"log"

	model "backend/modules/pathway/models"

	"github.com/gin-gonic/gin"
)

func Post(ctx *gin.Context){

	var newPathway model.Pathway

	if err := ctx.BindJSON(&newPathway);err != nil {
		log.Fatalf("failed pathway controllers post %v", err)
	}

	pathway := model.Post(newPathway)

	ctx.JSON(201, pathway
	)
}