package main

import (
	"context"

	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/apus-run/gaia"
	pb "github.com/apus-run/gaia/examples/user-service/api"
	"github.com/apus-run/gaia/examples/user-service/pkg"
	"github.com/apus-run/gaia/log"
	"github.com/apus-run/gaia/middleware/recovery"
	grpcserver "github.com/apus-run/gaia/transport/grpc"
	httpserver "github.com/apus-run/gaia/transport/http"
)

var (
	// Name is the name of the compiled software.
	Name = "user-service-client"
	// Version is the version of the compiled software.
	Version = "v1.0.0"
)

var (
	UserServiceServerClient pb.UserServiceClient
)

func ConnectGrpcServer() {
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

	UserServiceServerClient = pb.NewUserServiceClient(conn)
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

func NewRouter() *gin.Engine {
	g := gin.New()
	// 使用gaia中间件
	g.Use(pkg.Middlewares(recovery.Recovery()))

	g.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	g.POST("/user", createUser)

	return g
}

func main() {
	ConnectGrpcServer()

	// http server
	httpServer := httpserver.NewServer(
		httpserver.Address(":8000"),
		//httpserver.Middleware(
		//	recovery.Recovery(),
		//),
	)
	gin.SetMode("release")
	router := NewRouter()

	httpServer.Handler = router

	app := gaia.New(
		gaia.WithName(Name),
		gaia.WithVersion(Version),
		gaia.WithLogger(log.GetLogger()),
		gaia.WithServer(
			httpServer,
		),
	)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
