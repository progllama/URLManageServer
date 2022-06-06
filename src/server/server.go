package server

import (
	"url_manager/app/api"
	"url_manager/app/middlewares"
	"url_manager/session"

	"github.com/gin-gonic/gin"
)

// dependency
// var (
// 	sessionFactory = session.NewMemSessionFactory()
// )

func Open(port string) {
	router := gin.Default()

	session.InitializeSession(router)

	router.Static("favicon.ico", "app/assets/favicon.ico")
	router.Use(middlewares.ServeFavicon("./favicon.ico"))

	{
		errorHandler := middlewares.NewErrorHandler()
		router4user := router.Group("/users")
		router4user.Use(errorHandler.Handle)
		api := api.NewUserApi(nil, api.UserApiConfig{})
		router4user.GET("", api.Index)
		router4user.GET("/:userId", api.Show)
		router4user.POST("/:userId", api.Create)
		router4user.PUT("/:userId", api.Update)
		router4user.DELETE("/:userId", api.Delete)
	}

	// store, err := redis.NewStore(10, "tcp", "redis:6379", "", []byte("32bytes-secret-auth-key"))
	// if err != nil {
	// 	panic(err)
	// }
	// router.Use(sessions.Sessions("URLManager", store))

	// store := memstore.NewStore([]byte("secret"))
	// router.Use(sessions.Sessions("mysession", store))

	// router.Use(cors.New(cors.Config{
	// 	// 許可したいHTTPメソッドの一覧
	// 	AllowMethods: []string{
	// 		"POST",
	// 		"GET",
	// 		"OPTIONS",
	// 		"PUT",
	// 		"DELETE",
	// 	},
	// 	// 許可したいHTTPリクエストヘッダの一覧
	// 	AllowHeaders: []string{
	// 		"Access-Control-Allow-Headers",
	// 		"Content-Type",
	// 		"Content-Length",
	// 		"Accept-Encoding",
	// 		"X-CSRF-Token",
	// 		"Authorization",
	// 	},
	// 	// 許可したいアクセス元の一覧
	// 	AllowOrigins: []string{
	// 		"http://localhost:80",
	// 	},
	// 	AllowCredentials: true,
	// 	// 自分で許可するしないの処理を書きたい場合は、以下のように書くこともできる
	// 	// AllowOriginFunc: func(origin string) bool {
	// 	//  return origin == "https://www.example.com:8080"
	// 	// },
	// 	// preflight requestで許可した後の接続可能時間
	// 	// https://godoc.org/github.com/gin-contrib/cors#Config の中のコメントに詳細あり
	// 	MaxAge: 24 * time.Hour,
	// }))

	router.Run(port)
}
