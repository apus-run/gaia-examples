package main

import (
	"context"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/apus-run/gaia"
	"github.com/apus-run/gaia/middleware"
	"github.com/apus-run/gaia/middleware/recovery"
	"github.com/apus-run/gaia/pkg/ginx"
	grpcserver "github.com/apus-run/gaia/transport/grpc"
	httpserver "github.com/apus-run/gaia/transport/http"
	"github.com/apus-run/sea-kit/log"

	pb "github.com/apus-run/gaia/examples/user-service/api"
)

var (
	// Name is the name of the compiled software.
	Name = "user-service-client"
	// Version is the version of the compiled software.
	Version = "v1.0.0"

	id, _ = os.Hostname()
)

var (
	UserServiceServerClient pb.UserServiceClient
)

func ConnectGrpcServer() pb.UserServiceClient {
	conn, err := grpcserver.DialInsecure(
		context.Background(),
		grpcserver.WithEndpoint("127.0.0.1:9000"),
		grpcserver.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	c := pb.NewUserServiceClient(conn)

	return c
}

func HandleGrpcErrorToHttp(err error, c *gin.Context) {
	if e, ok := status.FromError(err); ok {
		switch e.Code() {
		case codes.NotFound:
			c.JSON(http.StatusNotFound, gin.H{
				"msg": e.Message(),
			})
		case codes.Internal:
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg:": "内部错误",
			})
		case codes.InvalidArgument:
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "参数错误",
			})
		case codes.Unavailable:
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "用户服务不可用",
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": e.Code(),
			})
		}
		return
	}

}

type userForm struct {
	UserName string `form:"username" json:"username" binding:"required,min=2,max=100"`
	Email    string `form:"email" json:"email" binding:"required"`
}

func createUser(c *gin.Context) {
	user := userForm{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	request := &pb.CreateUserRequest{
		Username: user.UserName,
		Email:    user.Email,
	}

	u, err := UserServiceServerClient.CreateUser(context.Background(), request)
	if err != nil {
		HandleGrpcErrorToHttp(err, c)
		return
	}

	//rsp := &pb.CreateUserReply{
	//	Id:       10001,
	//	Username: user.Username,
	//	Email:    user.Email,
	//}

	rsp := map[string]interface{}{
		"id":       10001,
		"username": u.Username,
		"email":    u.Email,
	}

	c.JSON(http.StatusOK, rsp)
}

type login struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func Login(c *ginx.Context) {
	var param login
	err := c.Bind(&param)
	if err != nil {
		c.JSONE(1, err.Error(), nil)
	}
	if len(param.Username) < 2 || len(param.Password) > 20 {
		c.JSONE(1, "username length should between 2 ~ 20", "")
		return
	}

	log.Infof("用户: %v, %v", param.Password, param.Username)
	// 数据库操作

	// c.JSONOK("")
	// c.Success(gin.H{})
	c.Success(param)

	return
}

func customMiddleware(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
		log.Info("自定义标准插件")
		reply, err = handler(ctx, req)
		return
	}
}

func NewRouter() *gin.Engine {
	g := gin.New()
	// 使用gaia中间件
	g.Use(ginx.Middlewares(recovery.Recovery(), customMiddleware))

	g.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, map[string]string{"welcome": name})
	})

	g.POST("/user", createUser)
	g.POST("/login", ginx.Handle(Login))

	return g
}

func main() {
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
	)

	c := ConnectGrpcServer()
	UserServiceServerClient = c

	gin.SetMode("release")
	router := NewRouter()

	// http server
	httpServer := httpserver.NewServer(
		httpserver.Address(":8000"),
		//httpserver.Middleware(
		//	recovery.Recovery(),
		//),
	)

	httpServer.Handler = router

	app := gaia.New(
		gaia.WithName(Name),
		gaia.WithVersion(Version),
		gaia.WithLogger(logger),
		gaia.WithServer(
			httpServer,
		),
	)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
