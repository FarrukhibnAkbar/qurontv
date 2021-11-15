package models

import(
	"backend/utils"
)

func Delete(deletePathway Pathway){
	utils.DB.Delete(&deletePathway) 
}