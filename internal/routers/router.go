package routers

import (
	"net/http"
	"time"

	"github.com/go-programming-tour-book/blog-service/pkg/limiter"

	"github.com/go-programming-tour-book/blog-service/global"

	"github.com/gin-gonic/gin"

	// something import
	_ "github.com/go-programming-tour-book/blog-service/docs"
	"github.com/go-programming-tour-book/blog-service/internal/middleware"
	"github.com/go-programming-tour-book/blog-service/internal/routers/api"
	"github.com/go-programming-tour-book/blog-service/internal/routers/api/v1"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.LimiterBucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	},
)

// NewRouter cool
func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}

	r.Use(middleware.Tracing())
	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
	r.Use(middleware.Translations())

	member := v1.NewMember()
	product := v1.NewProduct()
	records := v1.NewRecords()
	upload := api.NewUpload()
	r.GET("/debug/vars", api.Expvar)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload/file", upload.UploadFile)
	r.POST("/auth", api.GetAuth)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	apiv1 := r.Group("/api/v1")
	apiv1.Use() //middleware.JWT()
	{

		// 創建用戶
		apiv1.POST("/members", member.Create)
		// 删除指定用戶
		apiv1.DELETE("/members/:id", member.Delete)
		// 更新指定文用戶
		apiv1.PUT("/members/:id", member.Update)
		// 獲取指定用戶
		apiv1.GET("/members/:id", member.Get)

		// 創建產品
		apiv1.POST("/products", product.Create)
		// 删除指定產品
		apiv1.DELETE("/products/:id", product.Delete)
		// 更新指定產品
		apiv1.PUT("/products/:id", product.Update)
		// 獲取指定產品
		apiv1.GET("/products/:id", product.Get)

		// 創建紀錄
		apiv1.POST("/recordss", records.Create)
		// 删除指定紀錄
		apiv1.DELETE("/recordss/:id", records.Delete)
		// 更新指定紀錄
		apiv1.PUT("/recordss/:id", records.Update)
		// 獲取指定紀錄
		apiv1.GET("/recordss/:id", records.Get)

	}

	return r
}
