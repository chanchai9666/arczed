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

// where user_id
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

// where const_id
func WhereConstId(data ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		fieldName := "const_id"
		x := data[0] == ""
		if len(data) == 0 || x {
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

// where group_id
func WhereGroupId(data ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		fieldName := "group_id"
		x := data[0] == ""
		if len(data) == 0 || x {
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

// where amount>1000 [แบบฟิกค่า]
func AmountGreaterThan1000(db *gorm.DB) *gorm.DB {
	return db.Where("amount > ?", 1000)
}

// where pay_mode=card [แบบฟิกค่า]
func PaidWithCreditCard(db *gorm.DB) *gorm.DB {
	return db.Where("pay_mode = ?", "card")
}

// where pay_mode=cod [แบบฟิกค่า]
func PaidWithCod(db *gorm.DB) *gorm.DB {
	return db.Where("pay_mode = ?", "cod")
}
