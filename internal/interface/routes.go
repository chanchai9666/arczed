package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"nhooyr.io/websocket"
	"xorm.io/xorm"

	"arczed/docs"
	"arczed/internal/entities/models"
	"arczed/internal/handlers"
	"arczed/internal/repositories"
	"arczed/internal/usecase"
	"arczed/pkg/middleware"
)

// @title User API
// @version 1.0
// @description This is a sample server for user management.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func (s *Server) RegisterRoutes() http.Handler {

	r := gin.Default()

	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true, // อนุญาตให้ทุกโดเมนเข้าถึง
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Accept", "Authorization", "Custom-Header"},
	}))

	db := s.db.MainConnect()
	db.AutoMigrate(&models.Users{}, &models.UsersLevels{}, &models.ConfigConstant{})

	// กำหนดรายละเอียดของส่วน auth Bearer
	// @securityDefinitions.apikey ApiKeyAuth
	// @name Authorization
	// @in ใส่ค่า Bearer เว้นวรรคและตามด้วย TOKEN  ex(Bearer ?????????)
	// END กำหนดค่าให้ swagger
	// =======================================================
	// เพิ่ม middleware สำหรับการเข้าถึง Swagger UI ด้วยควบคุมสิทธิ์

	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Title = "ArcZed"
	docs.SwaggerInfo.Description = "ArcZed-API "
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.BasePath = "/"

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	jwtAuth := middleware.AuthMiddleware(s.config.JwtSECRETKEY)

	r.GET("/", s.HelloWorldHandler)
	r.GET("/docs/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	c := gin.Context{}
	api := r.Group("/api")
	userRep := repositories.NewUsersRepository(db, &s.config, middleware.GetUserProfile(&c))
	userService := usecase.NewUserService(userRep)
	userEndPoint := handlers.NewUserEndPoint(userService)
	api.POST("/login", userEndPoint.Login)
	api.POST("/refreshToken", jwtAuth, userEndPoint.RefreshToken)

	us := api.Group("/users")
	us.Use(jwtAuth)
	{
		us.GET("/", userEndPoint.FindUser)
		us.GET("/:user_id", userEndPoint.FindUsersByUserId)
		us.GET("/usersAll", userEndPoint.FindUserAll)
		us.POST("/createUsers", userEndPoint.CreateUsers)
		us.POST("/updateUsers", userEndPoint.UpdateUsers)
		us.DELETE("/deleteUsers/:user_id", userEndPoint.DeleteUsers)
	}

	ct := api.Group("/const")
	ct.Use(jwtAuth)
	constRepo := repositories.NewConstRepository(db, &s.config, "admin")
	constService := usecase.NewConstRepository(constRepo)
	constEndPoint := handlers.NewConstEndpoint(constService)
	{
		ct.GET("/findConstAll", constEndPoint.FindConstAll)
		ct.GET("/findConst", constEndPoint.FindConst)
		ct.POST("/createConst", constEndPoint.CreateConst)
		ct.POST("/updateConst", constEndPoint.UpdateConst)
		ct.DELETE("/deleteConst/:group_id/:const_id", constEndPoint.DeleteConst)

	}

	r.GET("/health", s.healthHandler)

	r.GET("/websocket", s.websocketHandler)

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"
	c.JSON(http.StatusOK, resp)
}

// HealthCheckHandler
// @summary Health Check
// @description Health checking for the service
// @id HealthCheckHandler
// @produce plain
// @response 200 {string} string "OK"
// @router /health [get]
func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}

func (s *Server) websocketHandler(c *gin.Context) {
	w := c.Writer
	r := c.Request
	socket, err := websocket.Accept(w, r, nil)

	if err != nil {
		log.Printf("could not open websocket: %v", err)
		_, _ = w.Write([]byte("could not open websocket"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer socket.Close(websocket.StatusGoingAway, "server closing websocket")

	ctx := r.Context()
	socketCtx := socket.CloseRead(ctx)

	for {
		payload := fmt.Sprintf("server timestamp: %d", time.Now().UnixNano())
		err := socket.Write(socketCtx, websocket.MessageText, []byte(payload))
		if err != nil {
			break
		}
		time.Sleep(time.Second * 2)
	}
}

// ฟังก์ชันเพื่อแปลงจาก camel case เป็น snake case
func toSnakeCase(str string) string {
	var result []rune
	for i, char := range str {
		if i > 0 && isUpperCase(char) {
			result = append(result, '_')
		}
		result = append(result, char)
	}
	return strings.ToLower(string(result))
}

// ฟังก์ชันเพื่อตรวจสอบว่าเป็นตัวพิมพ์ใหญ่หรือไม่
func isUpperCase(char rune) bool {
	return char >= 'A' && char <= 'Z'
}

func UpdateComments(engine *xorm.Engine, model interface{}) error {
	val := reflect.ValueOf(model).Elem()
	typeOfModel := val.Type()

	tableName := strings.ToLower(typeOfModel.Name())
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		comment := field.Tag.Get("comment")
		if comment != "" {
			columnName := toSnakeCase(field.Name) // แปลงเป็น snake case
			// อัปเดต comment ใน PostgreSQL
			query := fmt.Sprintf("COMMENT ON COLUMN %s.%s IS '%s';", tableName, columnName, comment)
			_, err := engine.Exec(query)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
