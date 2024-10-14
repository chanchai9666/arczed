package models

// ข้อมูล User
type Users struct {
	BaseColumn
	UserId      string        `json:"user_id" gorm:"type:varchar(50);primaryKey"`                   //ไอดี ของผู้ใช้งาน
	Password    string        `json:"password" gorm:"type:varchar(255);comment:รหัสผ่าน"`           //ชื่อ โปรไฟล์
	Name        string        `json:"name" gorm:"type:varchar(50);comment:ชื่อ"`                    //ชื่อ
	SurName     string        `json:"sur_name" gorm:"type:varchar(50);comment:นามสกุล"`             //นามสกุล
	Birthday    *string       `json:"birth_day" gorm:"type:date;comment:วันเกิด;default:null;"`     //วันเกิด
	PhoneNumber *string       `json:"phone_number" gorm:"type:varchar(20);comment:หมายเลขโทรศัพท์"` //หมายเลขโทรศัพท์
	IdCard      *string       `json:"id_card" gorm:"type:varchar(30);comment:รหัสบัตรประจำตัว"`     //รหัสบัตรประจำตัว
	Email       *string       `json:"email" gorm:"type:varchar(100);comment:อีเมล"`                 //อีเมล
	Level       []UsersLevels `gorm:"foreignKey:UserID;references:UserId"`                          //ความสัมพันธ์กับ UserLevels
}

type UsersLevels struct {
	BaseColumn
	ID     uint   `gorm:"column:id;type:int;primaryKey;autoIncrement"`
	UserID string `gorm:"column:user_id;index;type:varchar(50);not null"` // ต้องกำหนดประเภทและขนาดให้ตรงกับ Users.UserId
	Level  string `gorm:"column:level;type:varchar(50);not null"`
}
