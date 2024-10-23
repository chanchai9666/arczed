package schemas

type ConfigConstant struct {
	ConstId   string `json:"const_id" form:"const_id"`     //ไอดีค่าคงที่
	GroupId   string `json:"group_id" form:"group_id"`     //ไอดีกลุ่มค่าคงที่
	NameTH    string `json:"name_th" form:"name_th"`       //ชื่อค่าคงที่ TH
	NameEN    string `json:"name_en" form:"name_en"`       //ชื่อค่าคงที่ EN
	RefValue1 string `json:"ref_value1" form:"ref_value1"` //ค่าอ้างอิง 1
	RefValue2 string `json:"ref_value2" form:"ref_value2"` //ค่าอ้างอิง 2
	RefValue3 string `json:"ref_value3" form:"ref_value3"` //ค่าอ้างอิง 3
	Sort      int    `json:"sort" form:"sort"`             //ลำดับ
}
