package usecase

import (
	"arczed/internal/entities/schemas"
	"arczed/internal/repositories"
)

type ConstService interface {
	CreateConst(req *schemas.ConfigConstant) error
	UpdateConst(req *schemas.ConfigConstant) error
	DeleteConst(req *schemas.ConfigConstant) error
}

type ConstRepository struct {
	repo repositories.ConstRepository
}

func NewConstRepository(repo repositories.ConstRepository) ConstService {
	return &ConstRepository{
		repo: repo,
	}
}

func (s *ConstRepository) CreateConst(req *schemas.ConfigConstant) error {
	return s.repo.Create(req)
}

func (s *ConstRepository) UpdateConst(req *schemas.ConfigConstant) error {
	return s.repo.Update(req)
}

func (s *ConstRepository) DeleteConst(req *schemas.ConfigConstant) error {
	return s.repo.Delete(req.ConstId, req.GroupId)
}
