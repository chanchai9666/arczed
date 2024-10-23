package models

import (
	"errors"

	"github.com/chanchai9666/aider"
	"gorm.io/gorm"
)

func (u *Users) BeforeCreate(tx *gorm.DB) (err error) {
	// เรียก BeforeCreate ของ BaseColumn
	if err := u.BaseColumn.BeforeCreate(tx); err != nil {
		return err
	}

	if u.Password == "" {
		u.Password = "1234"
	}
	HashPassword, err := aider.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = HashPassword
	return
}

func (u *Users) AfterCreate(tx *gorm.DB) (err error) {
	//เพิ่มข้อมูล Level เริ่มต้นหลังจากเพิ่มข้อมูล User
	leveUser := UsersLevels{
		UserID: u.UserId,
		Level:  "user",
	}
	leveUser.CreatedAt = u.CreatedAt
	leveUser.CreatedUser = u.CreatedUser

	tx.Create(&leveUser)
	if tx.Error != nil {
		return tx.Error
	}
	return
}

func (u *Users) BeforeUpdate(tx *gorm.DB) (err error) {
	// if tx.Statement.Changed("DeletedAt") {
	// 	// ถ้ากำลังทำ Soft Delete ไม่ต้องเช็ค UserId
	// 	return nil
	// }
	// ตรวจสอบว่ากำลังลบหรือไม่
	if u.UserId == 0 {
		return errors.New("UserId is Empty")
	}
	return nil
}

func (u *Users) BeforeDelete(tx *gorm.DB) (err error) {
	// อัปเดตฟิลด์ IsActive และ DeletedAt โดยใช้ฟังก์ชัน Updates ของ GORM
	return nil
}
