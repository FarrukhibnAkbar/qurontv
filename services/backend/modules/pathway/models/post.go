package models

import(
	"backend/utils"
)

func Post(newPathway Pathway) Pathway {
	utils.DB.Table("pathways").Create(&newPathway)
	return newPathway
}