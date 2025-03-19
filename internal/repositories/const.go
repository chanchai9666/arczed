package repositories

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"arczed/configs"
)

type constRequest struct {
	db *gorm.DB
}

type mainDB struct {
	*constRequest
	userId string
	config *configs.Config
}

type constDB struct {
	*constRequest
	userId string
	config *configs.Config
}

// กำหนดให้ใหม่ให้ repo ของ user ใหม่เนื่องจากมีการรับค่า string ที่จำเป็นต้องใช้งาน
type userDB struct {
	*constRequest
	userId string
	config *configs.Config
}

// Insert function for inserting data into the database
func Insert[T any](database *gorm.DB, data T) error {
	return database.Create(&data).Error // ใช้ &data เพื่อส่ง pointer
}
func Updates[T any](database *gorm.DB, data *T) error {
	tx := database.Updates(data)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("update error : record not found")
	}
	return nil
}
func Delete[T any](database *gorm.DB, data *T) error {

	tx := database.Select("is_active").Updates(data)
	if tx.Error != nil {
		return tx.Error
	}
	tx = database.Delete(data)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("delete error : record not found")
	}
	return nil
}
func UpdateInterface[T any](database *gorm.DB, model *T, data map[string]interface{}) error {
	return database.Model(model).Updates(data).Error
}
func Find[T any](database *gorm.DB, data *T) error {
	if err := database.Find(data).Error; err != nil {
		return err
	}
	return nil
}
func Count[T any](database *gorm.DB, data *T) int64 {
	var count int64
	if err := database.Model(data).Count(&count).Error; err != nil {
		return 0
	}
	return count
}

// สำหรับ Transaction
func Transaction(db *gorm.DB, handler func(*gorm.DB) error) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	fmt.Println("======")
	fmt.Println("Transaction Begin")
	fmt.Println("======")
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r) // ต่อให้มี panic ก็ตาม, ต้อง Rollback ให้เรียบร้อย
		} else if err := tx.Commit().Error; err != nil {
			tx.Rollback()
		}
		fmt.Println("======")
		fmt.Println("Transaction Commit")
		fmt.Println("======")
	}()

	if err := handler(tx); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// Create create record database
func Create(database *gorm.DB, i interface{}) error {
	if err := database.Create(i).Error; err != nil {
		return err
	}
	return nil
}
