package repositories

import (
	"gorm.io/gorm"

	"arczed/configs"
	"arczed/internal/entities/models"
	"arczed/internal/entities/schemas"
)

type ProductRepository interface {
	SaveProduct() error
	FindProduct(req *models.AddProduct) error
}

func NewConstRepository(db *gorm.DB, config *configs.Config, userId string) ConstRepository {
	return &beseDB{
		baseRequest: &baseRequest{db: db},
		userId:      userId, // กำหนดค่าให้กับ Name
		config:      config,
	}
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &beseDB{
		baseRequest: &baseRequest{db: db},
	}
}

func NewUsersRepository(db *gorm.DB, config *configs.Config, userId string) UsersRepository {
	return &userDB{
		baseRequest: &baseRequest{db: db}, // ใช้ชื่อฟิลด์เพื่อกำหนดค่า
		userId:      userId,               // กำหนดค่าให้กับ Name
		config:      config,
	}
}

type UsersRepository interface {
	CreateUsers(req *schemas.AddUsers) error                     //เพิ่มข้อมูล Users
	FindUsers(req *schemas.FindUsersReq) ([]models.Users, error) //ค้นหาข้อมูล Users
	UpdateUser(req *schemas.AddUsers) error                      //อัพเดตข้อมูล Users
	DeletedUser(userID uint64) error                             //ลบข้อมูล Users
	NewJwt(req *schemas.JwtReq) (string, error)                  //สร้าง Jwt Token
}

type ConstRepository interface {
	Create(req *schemas.ConfigConstant) error                                         //เพิ่ม ค่าคงที่
	Update(req *schemas.ConfigConstant) error                                         //แก้ไข ค่าคงที่
	Delete(id, group string) error                                                    //ลบค่าคงที่
	FindPage(req *schemas.ConfigConstant) (*Pagination[models.ConfigConstant], error) //ค้าหาค่าคงที่แบบแบ่งหน้า
	FindAll(req *schemas.ConfigConstant) ([]models.ConfigConstant, error)             //ค้นหาค่าคงที่ทั้งหมด
}
