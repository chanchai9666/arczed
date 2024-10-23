package models

import (
	"errors"

	"gorm.io/gorm"
)

func (u *ConfigConstant) BeforeCreate(tx *gorm.DB) (err error) {
	// เรียก BeforeCreate ของ BaseColumn
	if err := u.BaseColumn.BeforeCreate(tx); err != nil {
		return err
	}

	if u.ConstID == "" && u.GroupID == "" {
		return errors.New("ConstID และ GroupID ค่าคงที่ ต้องไม่เป็นค่าว่าง")
	}

	return
}