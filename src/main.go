package main

import (
	"fmt"
	"net/http"
	"strconv"

	"log"
	"time"

	"github.com/ahpehgit/golang-boilerplate/config"
	"github.com/ahpehgit/golang-boilerplate/constants"
	"github.com/ahpehgit/golang-boilerplate/setup"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"

	sampleController "github.com/ahpehgit/golang-boilerplate/controller/sample"
	sampleService "github.com/ahpehgit/golang-boilerplate/service/sample"
)

var validate *validator.Validate
var translation *ut.Translator

func main() {
	fmt.Println("Starting my Golang project...")

	setup.SetupConfig()
	validate = setup.RegisterValidators()
	translation = setup.RegisterTranslations(validate)

	// Set run mode
	pkg := viper.GetString("OSPRESERVATION_PACKAGE")
	mode := gin.DebugMode
	if pkg == "release" {
		log.Print("*** Running in Release mode ***")
		mode = gin.ReleaseMode
	}

	gin.SetMode(mode)

	router := gin.Default()
	router.SetTrustedProxies([]string{"localhost"})

	corsConfig := &constants.Cors{}

	// Set CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     corsConfig.GetOrigins(),
		AllowMethods:     corsConfig.GetAllowMethods(),
		AllowHeaders:     corsConfig.GetAllowHeaders(),
		ExposeHeaders:    corsConfig.GetExposedHeaders(),
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	var (
		sampleService    *sampleService.SampleService       = sampleService.NewSampleService("Sample Service")
		sampleController *sampleController.SampleController = sampleController.NewSampleController(sampleService)
	)

	//* Root
	router.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusPermanentRedirect, constants.URL_PING)
	})

	//* Ping health status
	router.GET(constants.URL_PING, sampleController.Ping)

	// Get host and port from env variables, if not default to config values
	host := viper.GetString("OSPRESERVATION_HOST")
	port := viper.GetString("OSPRESERVATION_PORT")

	if host == "" {
		host = config.Configuration.Server.Host
	}

	if port == "" {
		port = strconv.Itoa(config.Configuration.Server.Port)
	}

	env := viper.GetString("OSPRESERVATION_ENVIRONMENT")
	if env == "" {
		env = config.Configuration.Server.Environment
		if env == "" {
			log.Fatal("***** ERROR: OSPRESERVATION_ENVIRONMENT NOT DEFINED *****")
		}

		log.Printf("Environment is \"%s\"", env)
	}

	config.Configuration.Server.Environment = env

	// Start server
	log.Printf("Server [%s] running in %s:%s", env, host, port)
	router.Run(host + ":" + port)
}
