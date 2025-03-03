package controller

import (
	"log"
	"net/http"

	sampleService "github.com/ahpehgit/golang-boilerplate/service/sample"
	"github.com/gin-gonic/gin"
)

type SampleController struct {
	sampleService *sampleService.SampleService
}

func NewSampleController(sampleService *sampleService.SampleService) *SampleController {
	return &SampleController{
		sampleService,
	}
}

func (controller *SampleController) Ping(ctx *gin.Context) {
	log.Println(controller.sampleService.GetServiceName())

	ctx.IndentedJSON(http.StatusBadRequest, gin.H{
		"message": "Ok",
	})
}
