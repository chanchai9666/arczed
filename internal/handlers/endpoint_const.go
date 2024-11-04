package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"

	_ "arczed/internal/entities/models"
	"arczed/internal/entities/schemas"
	"arczed/internal/rander"
	"arczed/internal/usecase"
)

type constEndPoint struct {
	service usecase.ConstService
}

func NewConstEndpoint(service usecase.ConstService) *constEndPoint {
	return &constEndPoint{service: service}
}

// @Tags Const
// @Summary เพิ่มข้อมูลค่าคงที่
// @Description เพิ่มข้อมูลค่าคงที่
// @Accept  json
// @Produce  json
// @Param Accept-Language header string false "(en, th)" default(th)
// @Param request body schemas.ConfigConstant false " request body "
// @Success 200 {object} schemas.HTTPError
// @Failure 400 {object} schemas.HTTPError
// @Router /api/const/createConst [post]
// @Security ApiKeyAuth
func (en *constEndPoint) CreateConst(c *gin.Context) {
	rander.RespSuccess(c, en.service.CreateConst, &schemas.ConfigConstant{})
}

// @Tags Const
// @Summary แก้ไขข้อมูลค่าคงที่
// @Description แก้ไขข้อมูลค่าคงที่
// @Accept  json
// @Produce  json
// @Param Accept-Language header string false "(en, th)" default(th)
// @Param request body schemas.ConfigConstant false " request body "
// @Success 200 {object} schemas.HTTPError
// @Failure 400 {object} schemas.HTTPError
// @Router /api/const/updateConst [post]
// @Security ApiKeyAuth
func (en *constEndPoint) UpdateConst(c *gin.Context) {
	rander.RespSuccess(c, en.service.UpdateConst, &schemas.ConfigConstant{})
}

// @Tags Const
// @Summary ลบข้อมูล ค่าคงที่
// @Description ลบข้อมูล ค่าคงที่
// @Accept  json
// @Produce  json
// @Param Accept-Language header string false "(en, th)" default(th)
// @Param group_id path string true "group id"
// @Param const_id path string true "const id"
// @Success 200 {object} schemas.HTTPError
// @Failure 400 {object} schemas.HTTPError
// @Router /api/const/deleteConst/{group_id}/{const_id} [delete]
// @Security ApiKeyAuth
func (en *constEndPoint) DeleteConst(c *gin.Context) {
	rander.RespSuccess(c, en.service.DeleteConst, &schemas.ConfigConstant{})
}

// @Tags Const
// @Summary แสดงข้อมูลค่าคงที่ทั้งหมด
// @Description แสดงข้อมูลค่าคงที่ทั้งหมด
// @Accept  json
// @Produce  json
// @Param Accept-Language header string false "(en, th)" default(th)
// @Success 200 {object} schemas.Pagination[models.ConfigConstant]
// @Failure 400 {object} schemas.HTTPError
// @Router /api/const/findConstAll [get]
// @Security ApiKeyAuth
func (en *constEndPoint) FindConstAll(c *gin.Context) {
	fmt.Println("11111")
	rander.RespJsonNoReq(c, en.service.Find)
}
