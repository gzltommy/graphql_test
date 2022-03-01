// 从schema开始
// schema/schema.go
package schema

// 引入包 graphql-go
import (
	"github.com/graphql-go/graphql"
)

// 定义跟查询节点
var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name:        "RootQuery",
	Description: "Root Query",
	Fields: graphql.Fields{
		"hello": &queryHello, // queryHello 参考 schema/hello.go
	},
})

// 定义 Schema 用于 http handler 处理
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: nil, // 需要通过 GraphQL 更新数据，可以定义 Mutation
})
