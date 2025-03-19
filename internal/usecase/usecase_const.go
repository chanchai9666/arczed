package usecase

import (
	"github.com/jinzhu/copier"

	"arczed/internal/entities/models"
	"arczed/internal/entities/schemas"
	"arczed/internal/repositories"
)

type ConstService interface {
	CreateConst(req *schemas.ConfigConstant) error
	UpdateConst(req *schemas.ConfigConstant) error
	DeleteConst(req *schemas.ConfigConstant) error
	Find(req *schemas.ConfigConstant) (*schemas.Pagination[models.ConfigConstant], error) //หาค่าคงที่แบบแบ่งหน้า
	FindAll(req *schemas.ConfigConstant) ([]schemas.ConfigConstantResp, error)            //หาค่าคงที่ทั้งหมด
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

func (s *ConstRepository) Find(req *schemas.ConfigConstant) (*schemas.Pagination[models.ConfigConstant], error) {
	data, err := s.repo.FindPage(req)
	if err != nil {
		return nil, err
	}

	// constData := []schemas.ConfigConstantResp{}
	_ = data.Rows

	newData := &schemas.Pagination[models.ConfigConstant]{}
	if err := copier.Copy(&newData, data); err != nil {
		return nil, err
	}
	_ = newData
	return newData, nil
}

func (s *ConstRepository) FindAll(req *schemas.ConfigConstant) ([]schemas.ConfigConstantResp, error) {
	req.IsActive = "-1" //-1 คือทั้งหมด
	data, err := s.repo.FindAll(req)
	if err != nil {
		return nil, err
	}
	constData := []schemas.ConfigConstantResp{}
	if err := copier.Copy(&constData, data); err != nil {
		return nil, err
	}
	for i, v := range constData {
		if v.BaseColumn.IsActive != nil && *v.BaseColumn.IsActive == 1 {
			constData[i].IsActiveTxt = "ใช้งาน"
		} else {
			constData[i].IsActiveTxt = "ไม่ใช้งาน"
		}
	}

	return constData, nil
}
