package api

import (
	"os"

	cfg "github.com/dredfort42/tools/configreader"
	loger "github.com/dredfort42/tools/logprinter"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Host               string
	Port               string
	CorsStatus         string
	AuthServerURL      string
	IdentifyPathUser   string
	IdentifyPathDevice string
	ChangePathEmail    string
	DeletePathUser     string
}

var server Server

// ApiInit starts the profile service
func ApiInit() {
	server.Host = cfg.Config["profile.host"]
	if server.Host == "" {
		panic("profile.host is not set")
	}

	server.Port = cfg.Config["profile.port"]
	if server.Port == "" {
		panic("profile.port is not set")
	}

	server.CorsStatus = cfg.Config["profile.cors"]
	if server.CorsStatus == "" {
		loger.Warning("profile.cors is not set | CORS is disabled")
		server.CorsStatus = "false"
	}

	server.AuthServerURL = cfg.Config["auth.url"]
	if server.AuthServerURL == "" {
		panic("auth.url is not set")
	}

	server.IdentifyPathUser = cfg.Config["auth.path.identify.user"]
	if server.IdentifyPathUser == "" {
		panic("auth.path.identify.user is not set")
	}

	server.IdentifyPathDevice = cfg.Config["auth.path.identify.device"]
	if server.IdentifyPathDevice == "" {
		panic("auth.path.identify.device is not set")
	}

	server.ChangePathEmail = cfg.Config["auth.path.change.email"]
	if server.ChangePathEmail == "" {
		panic("auth.path.change.email is not set")
	}

	server.DeletePathUser = cfg.Config["auth.path.delete.user"]
	if server.DeletePathUser == "" {
		panic("auth.path.delete.user is not set")
	}

	if os.Getenv("DEBUG") != "true" && os.Getenv("DEBUG") != "1" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.Default()

	if server.CorsStatus == "true" || server.CorsStatus == "1" {
		router.Use(cors.Default())
	}

	// Apply the middleware to the routes you want to protect
	authorized := router.Group("/", AuthMiddleware())
	{
		authorized.POST("/api/v1/profile/user", UserCreate)
		authorized.GET("/api/v1/profile/user", UserGet)
		authorized.PATCH("/api/v1/profile/user", UserUpdate)
		authorized.DELETE("/api/v1/profile/user", UserDelete)
		authorized.POST("/api/v1/profile/user/email", UserChangeEmail)
		authorized.POST("/api/v1/profile/devices", DeviceCreate)
		authorized.GET("/api/v1/profile/devices", DevicesGet)
		authorized.PUT("/api/v1/profile/devices", DeviceUpdate)
		authorized.DELETE("/api/v1/profile/devices", DeviceDelete)
	}

	// // Unprotected route
	// router.GET("/unprotected", UnprotectedEndpoint)

	url := server.Host + ":" + server.Port
	loger.Success("Service successfully started", url)
	router.Run(url)
}
