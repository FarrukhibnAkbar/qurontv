package models

import(

	"backend/utils"
)

func Get() []Pathway {
	var pathway []Pathway
	utils.DB.Find(&pathway)
	return pathway
}