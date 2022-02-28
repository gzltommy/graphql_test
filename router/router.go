// route/router.go
package router

import (
	"graphql_test/controller/graphql"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()
}

func SetRouter() {
	// GET 方法用于支持 GraphQL 的 web 界面操作
	// 如果不需要 web 界面，可根据自己需求用 GET/POST 或者其他都可以
	Router.POST("/graphql", graphql.GraphqlHandler())
	Router.GET("/graphql", graphql.GraphqlHandler())
}
