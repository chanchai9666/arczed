package safety

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtConst struct {
	UserId   string `json:"user_id"`  //ผู้ใช้งาน
	Name     string `json:"name"`     //ชื่อ
	SurName  string `json:"sur_name"` //นามสกุล
	Email    string `json:"email"`    //อีเมล
	Level    string `json:"level"`    //เลเวล
	SafetyId string `json:"safety_id"`
}

type Claims struct {
	UserId   string `json:"user_id"`   //ผู้ใช้งาน
	Name     string `json:"name"`      //ชื่อ
	SurName  string `json:"sur_name"`  //นามสกุล
	Email    string `json:"email"`     //อีเมล
	Level    string `json:"level"`     //เลเวล
	SafetyId string `json:"safety_id"` //รหัสความปลอดภัย
	jwt.RegisteredClaims
}

// ฟังก์ชันสร้าง JWT
func GenerateJWT(jwtSecret string, req *JwtConst) (string, error) {
	// กำหนด claims
	claims := &Claims{
		UserId:   req.UserId,
		Name:     req.Name,
		SurName:  req.SurName,
		Email:    req.Email,
		Level:    req.Level,
		SafetyId: req.SafetyId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(6 * time.Hour)), // อายุ 6 ชั่วโมง
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "ArcZed", //ระบุแหล่งที่มาหรือผู้สร้างของ JWT นั้น ซึ่งมักจะเป็นชื่อของบริการ, ระบบ, หรือเซิร์ฟเวอร์ที่ออก token นี้
		},
	}

	// สร้าง token object
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// เซ็น token ด้วย secret key
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
