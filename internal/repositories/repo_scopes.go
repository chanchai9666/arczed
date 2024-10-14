package repositories

import (
	"fmt"

	"gorm.io/gorm"
)

func Where2(data ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		tableName := data[0]
		fieldName := data[1]
		value := data[2]
		condition := fmt.Sprintf("%s.%s = ?", tableName, fieldName)
		return db.Where(condition, value)
	}
}
func WhereIsActive(table ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		condition := "is_active=?"
		if len(table) > 0 && table[0] != "" {
			condition = table[0] + ".is_active=?"
		}
		return db.Where(condition, 1)
	}
}

// func WhereUserId(data ...string) func(db *gorm.DB) *gorm.DB {
// 	return func(db *gorm.DB) *gorm.DB {
// 		if len(data) == 0 {
// 			return db // ไม่มีค่า input คืนค่า db กลับไปโดยไม่ทำอะไร
// 		}
// 		value := data[0]
// 		condition := "user_id = ?"
// 		if len(data) > 1 {
// 			tableName := data[0]
// 			value = data[1]
// 			condition = fmt.Sprintf("%s.user_id = ?", tableName)
// 		}

// 		return db.Where(condition, value)
// 	}
// }

func WhereGroupId(data ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		fieldName := "group_id"
		if len(data) == 0 {
			return db // ไม่มีค่า input คืนค่า db กลับไปโดยไม่ทำอะไร
		}
		condition := fieldName + " = ?"
		value := data[0]
		if len(data) > 1 {
			// เรียกใช้งานฟังก์ชันที่คืนค่าจาก Where2
			return Where2(data[0], fieldName, data[1])(db)
		}
		return db.Where(condition, value)
	}
}

func WhereUserId(data ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		fieldName := "user_id"
		if len(data) == 0 {
			return db // ไม่มีค่า input คืนค่า db กลับไปโดยไม่ทำอะไร
		}
		condition := fieldName + " = ?"
		value := data[0]
		if len(data) > 1 {
			// เรียกใช้งานฟังก์ชันที่คืนค่าจาก Where2
			return Where2(data[0], fieldName, data[1])(db)
		}
		return db.Where(condition, value)
	}
}

func WhereConstId(data ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		fieldName := "const_id"
		if len(data) == 0 {
			return db // ไม่มีค่า input คืนค่า db กลับไปโดยไม่ทำอะไร
		}
		condition := fieldName + " = ?"
		value := data[0]
		if len(data) > 1 {
			// เรียกใช้งานฟังก์ชันที่คืนค่าจาก Where2
			return Where2(data[0], fieldName, data[1])(db)
		}
		return db.Where(condition, value)
	}
}
