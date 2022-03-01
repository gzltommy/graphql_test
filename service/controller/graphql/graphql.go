// 准备好了 GraphQL 在 Go 里面需要的东西之后，来看看如何跟 Gin 结合
// controller/graphql/graphql.go

package graphql

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
	schema2 "graphql_test/service/schema"
)

func GraphqlHandler() gin.HandlerFunc {
	h := handler.New(&handler.Config{
		Schema:   &schema2.Schema,
		Pretty:   true,
		GraphiQL: true,
	})

	// 只需要通过 Gin 简单封装即可
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
