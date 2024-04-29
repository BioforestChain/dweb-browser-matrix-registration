package tests

import (
	"context"
	"develop-template/internal/app/model/common"
	"develop-template/pkg/support-go/bootstrap"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"testing"
)

var DB *gorm.DB

func Setup() {
	testing.Init()
	bootstrap.DevEnv = bootstrap.EnvLocal
	bootstrap.Init()
	bootstrap.InitWeb([]gin.HandlerFunc{})
	DB = common.ConnectionObject(context.TODO()).DB
}
