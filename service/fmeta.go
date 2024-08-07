package service

import (
	"github.com/wangkebin/kbds-client/models"
	"gorm.io/gorm"
)

func CreateInBatches(db *gorm.DB, fmetas *[]models.FMeta, batchsize int) error {
	res := db.CreateInBatches(&fmetas, batchsize)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
