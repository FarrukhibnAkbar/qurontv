package models

import(
	"backend/utils"
)

func Update(putPathway Pathway) Pathway{

	utils.DB.Save(&putPathway)
	return putPathway
}