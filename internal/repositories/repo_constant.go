package repositories

import (
	"fmt"
	"log"

	"github.com/chanchai9666/aider"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"arczed/internal/entities/models"
	"arczed/internal/entities/schemas"
	"arczed/pkg"
)

// เพิ่มข้อมูลค่าคงที่
func (r *beseDB) Create(req *schemas.ConfigConstant) error {
	var data models.ConfigConstant
	if err := copier.Copy(&data, req); err != nil {
		return fmt.Errorf("failed to copy user data: %w", err)
	}
	if data.ConstID == "" {
		//รายการใหม่ให้ทำการสร้าง id
		countItem := r.CountItemByGroup(req.GroupId)
		countItem = countItem + 1
		maxSort := r.MaxSortByGroup(req.GroupId)
		if countItem >= maxSort {
			maxSort = countItem //ถ้าการนับจำนวน items ตามกลุ่มมากกว่า ค่า max ของ sort ให้เปลี่ยนการ running เป็นจำนวน items แทน
		}
		newId := req.GroupId + "-" + aider.PadZeros(3, aider.ToInt(countItem))
		data.ConstID = newId
		data.Sort = int(r.MaxSortByGroup(req.GroupId))
	}
	return Transaction(r.db, func(tx *gorm.DB) error {
		return Insert(r.db, &data)
	})
}

// นับ items ตามกลุ่ม
func (r *beseDB) CountItemByGroup(group string) int64 {
	numItem := Count(r.db.Unscoped().Scopes(WhereGroupId(group)), &models.ConfigConstant{})
	return numItem
}

// อัพเดตค่าคงที่
func (r *beseDB) Update(req *schemas.ConfigConstant) error {
	return Transaction(r.db, func(d *gorm.DB) error {
		var data models.ConfigConstant
		if err := copier.Copy(&data, req); err != nil {
			return err
		}
		return Updates(r.db.Scopes(WhereConstId(data.ConstID), WhereGroupId(data.GroupID)), &data)
	})
}

// ลบค่าคงที่
func (r *beseDB) Delete(id, group string) error {
	return Transaction(r.db, func(d *gorm.DB) error {
		var configConst models.ConfigConstant
		configConst.IsActive = pkg.IntPointer(0)
		return Delete(r.db.Scopes(WhereConstId(id), WhereGroupId(group)), &configConst)
	})
}

// query ค้นหาค่าคงที่
func (r *beseDB) FindConst(req *schemas.ConfigConstant) *gorm.DB {
	return r.db.
		Scopes(
			WhereIsActive(),
			WhereConstId(req.ConstId),
			WhereGroupId(req.GroupId),
		)
}

// ค้นหาแบบแบ่งหน้า
func (r *beseDB) FindPage(req *schemas.ConfigConstant) (*Pagination[models.ConfigConstant], error) {
	constData := []models.ConfigConstant{}
	pagination := &Pagination[models.ConfigConstant]{
		Sort: "group_id,sort asc",
	}
	if err := r.FindConst(req).
		Scopes(Paginate(r.db, models.ConfigConstant{}, pagination)).
		Find(&constData).Error; err != nil {
		return nil, err
	}
	pagination.Rows = constData
	return pagination, nil
}

// ค้นหาทั้งหมด
func (r *beseDB) FindAll(req *schemas.ConfigConstant) ([]models.ConfigConstant, error) {
	constData := []models.ConfigConstant{}
	if err := r.FindConst(req).Order("group_id,sort asc").Find(&constData).Error; err != nil {
		return nil, err
	}
	return constData, nil
}

// หาค่า Max Sort ของค่าคงที่ตามกลุ่มสำหรับ running sort ของแต่ละกลุ่ม
func (r *beseDB) MaxSortByGroup(group string) int64 {
	var maxSort int64
	err := r.db.Model(&models.ConfigConstant{}).Unscoped().Select("MAX(sort)").Scopes(WhereGroupId(group)).Scan(&maxSort).Error
	if err != nil {
		log.Println("Get Max Sort Error :", err.Error())
		return 0
	}
	maxSort = maxSort + 1
	return maxSort
}
