package repositories

import (
	"fmt"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"arczed/internal/entities/models"
	"arczed/internal/entities/schemas"
	"arczed/pkg/safety"
)

func (r *userDB) CreateUsers(req *schemas.AddUsers) error {
	var user models.Users
	if err := copier.Copy(&user, req); err != nil {
		return fmt.Errorf("failed to copy user data: %w", err)
	}
	if req.BirthDay == "" {
		user.Birthday = nil
	}
	return Transaction(r.db, func(tx *gorm.DB) error {
		return Insert(r.db, &user)
	})
}
func (r *userDB) FindUsers(req *schemas.FindUsersReq) ([]models.Users, error) {

	var allusers []models.Users
	tx := r.db
	if req.Email != "" {
		tx = tx.Where("email=?", req.Email)
	}
	if req.Name != "" {
		tx = tx.Where("name LIKE ?", "%"+req.Name+"%")
	}
	if req.SurName != "" {
		tx = tx.Where("sur_name=?", "%"+req.SurName+"%")
	}
	if req.UserId != "" {
		tx = tx.Where("user_id=?", req.UserId)
	}

	err := tx.Preload("Level").Scopes(WhereIsActive()).Find(&allusers).Error
	if err != nil {
		return nil, err
	}
	return allusers, nil
}
func (r *userDB) UpdateUser(req *schemas.AddUsers) error {
	var users models.Users
	if err := copier.Copy(&users, req); err != nil {
		return fmt.Errorf("failed to copy user data: %w", err)
	}

	return Transaction(r.db, func(tx *gorm.DB) error {
		que := r.db.Select("name", "sur_name").Scopes(WhereUserId(req.UserId))
		return Updates(que, &users)
	})
}
func (r *userDB) DeletedUser(userID *string) error {
	return Transaction(r.db, func(d *gorm.DB) error {
		var UserUpdate models.Users
		// สร้างตัวแปรสำหรับอัปเดต
		active := 0
		UserUpdate.IsActive = &active
		UserUpdate.UserId = *userID
		// ลบผู้ใช้
		if err := Delete(
			r.db.Scopes(WhereUserId(*userID)), &UserUpdate); err != nil {
			return err
		}
		return nil
	})
}

func (r *userDB) NewJwt(req *schemas.JwtReq) (string, error) {
	jwt, err := safety.GenerateJWT(r.config.JwtSECRETKEY, &safety.JwtConst{
		UserId:   req.UserId,
		Name:     req.Name,
		SurName:  req.SurName,
		Email:    req.Email,
		Level:    req.Level,
		SafetyId: "xxxxxxx",
	})
	if err != nil {
		return "nil", err
	}
	return jwt, nil
}
