package models

import(
	"backend/utils"
)

func Update(putPathway Pathway) {

	utils.DB.Save(&putPathway)
}