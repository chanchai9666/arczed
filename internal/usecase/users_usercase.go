package usecase

import (
	"errors"

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
	FindUsersByUserId(req *schemas.FindUsersByUserIdReq) (*models.Users, error)
	UpdateUsers(req *schemas.AddUsers) error
	DeleteUsers(req *schemas.AddUsers) error
	Login(req *schemas.LoginReq) (*schemas.LoginResp, error)
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

func (s *userRequest) FindUsersByUserId(req *schemas.FindUsersByUserIdReq) (*models.Users, error) {
	data, err := s.FindUsers(&schemas.FindUsersReq{UserId: req.UserId})
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
	return s.repo.DeletedUser(&req.UserId)
}

func (s *userRequest) Login(req *schemas.LoginReq) (*schemas.LoginResp, error) {
	result, err := s.FindUsersByUserId(&schemas.FindUsersByUserIdReq{UserId: req.UserId})
	if err != nil {
		return nil, err
	}

	var userLogin schemas.LoginResp
	if result.UserId != "" && result.Password != "" {
		if aider.CheckPassword(req.Password, result.Password) {
			aider.DDD(result)
			var userData schemas.UserResp
			if err := copier.Copy(&userData, result); err != nil {
				return nil, err
			}
			userLogin.User = userData

			s.repo.NewJwt()
			//สร้าง Jwt

		} else {
			return nil, aider.NewError(aider.ErrInternal, "ไม่สามารถเข้าสู่ระบบได้")
		}
	}
	return &userLogin, err
}
