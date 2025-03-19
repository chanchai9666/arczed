package usecase

import (
	"errors"
	"strings"

	"github.com/chanchai9666/aider"
	"github.com/jinzhu/copier"

	"arczed/internal/entities/models"
	"arczed/internal/entities/schemas"
	"arczed/internal/repositories"
)

type UsersService interface {
	FindUsers(req *schemas.FindUsersReq) ([]models.Users, error)
	FindUsersAll() ([]models.Users, error)
	CreateUsers(req *schemas.AddUsers) error
	FindUsersByEmail(req *schemas.FindUsersByEmailReq) (*models.Users, error)
	UpdateUsers(req *schemas.AddUsers) error
	DeleteUsers(req *schemas.AddUsers) error
	Login(req *schemas.LoginReq) (*schemas.LoginResp, error)
	RefreshToken(req *schemas.RefreshTokenReq) (*schemas.LoginResp, error)
}

type userRequest struct {
	repo repositories.UsersRepository
}

func NewUserService(repo repositories.UsersRepository) UsersService {
	return &userRequest{
		repo: repo,
	}
}

func (s *userRequest) FindUsers(req *schemas.FindUsersReq) ([]models.Users, error) {
	data, err := s.repo.FindUsers(req)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *userRequest) FindUsersByEmail(req *schemas.FindUsersByEmailReq) (*models.Users, error) {
	data, err := s.FindUsers(&schemas.FindUsersReq{Email: req.Email})
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New("user not found")
	}
	return &data[0], nil
}

func (s *userRequest) FindUsersAll() ([]models.Users, error) {
	data, err := s.repo.FindUsers(&schemas.FindUsersReq{})
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *userRequest) CreateUsers(req *schemas.AddUsers) error {
	return s.repo.CreateUsers(req)
}

func (s *userRequest) UpdateUsers(req *schemas.AddUsers) error {
	return s.repo.UpdateUser(req)
}

func (s *userRequest) DeleteUsers(req *schemas.AddUsers) error {
	return s.repo.DeletedUser(req.UserId)
}

func (s *userRequest) Login(req *schemas.LoginReq) (*schemas.LoginResp, error) {
	// ดึงข้อมูลผู้ใช้ตาม UserID
	result, err := s.FindUsersByEmail(&schemas.FindUsersByEmailReq{Email: req.Email})
	if err != nil {
		return nil, err
	}

	// ตรวจสอบการเข้าสู่ระบบ
	if result.UserId == 0 || result.Password == "" {
		return nil, aider.NewError(aider.ErrInternal, "ข้อมูลไม่ครบถ้วน")
	}

	// ตรวจสอบรหัสผ่าน
	if !aider.CheckPassword(req.Password, result.Password) {
		return nil, aider.NewError(aider.ErrInternal, "รหัสผ่านไม่ถูกต้อง")
	}

	// คัดลอกข้อมูลผู้ใช้
	var userLogin schemas.LoginResp
	var userData schemas.UserResp
	if err := copier.Copy(&userData, result); err != nil {
		return nil, err
	}

	levelVal := []string{}
	for _, v := range result.Level {
		levelVal = append(levelVal, v.Level)
	}
	strLvl := strings.Join(levelVal, ",")
	userLogin.User = userData
	userLogin.User.Level = levelVal

	// สร้าง JWT
	token, err := s.repo.NewJwt(&schemas.JwtReq{
		UserId:  aider.ToString(result.UserId),
		Name:    result.Name,
		SurName: result.SurName,
		Email:   result.Email,
		Level:   strLvl,
	})
	if err != nil {
		return nil, err
	}

	userLogin.AccessToken = "Bearer " + token
	aider.DDD(userLogin)
	return &userLogin, nil
}

func (s *userRequest) RefreshToken(req *schemas.RefreshTokenReq) (*schemas.LoginResp, error) {
	// ดึงข้อมูลผู้ใช้ตาม UserID
	result, err := s.FindUsersByEmail(&schemas.FindUsersByEmailReq{Email: req.Email})
	if err != nil {
		return nil, err
	}
	// ตรวจสอบความถูกต้อง
	if req.UserId != result.UserId {
		return nil, aider.NewError(aider.ErrInternal, "ข้อมูลไม่ครบถ้วน")
	}

	// คัดลอกข้อมูลผู้ใช้
	var userLogin schemas.LoginResp
	var userData schemas.UserResp
	if err := copier.Copy(&userData, result); err != nil {
		return nil, err
	}

	levelVal := []string{}
	for _, v := range result.Level {
		levelVal = append(levelVal, v.Level)
	}
	strLvl := strings.Join(levelVal, ",")
	userLogin.User = userData
	userLogin.User.Level = levelVal

	// สร้าง JWT
	token, err := s.repo.NewJwt(&schemas.JwtReq{
		UserId:  aider.ToString(result.UserId),
		Name:    result.Name,
		SurName: result.SurName,
		Email:   result.Email,
		Level:   strLvl,
	})
	if err != nil {
		return nil, err
	}

	userLogin.AccessToken = "Bearer " + token
	aider.DDD(userLogin)
	return &userLogin, nil

}
