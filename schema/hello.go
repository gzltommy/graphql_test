// schema/hello.go
package schema

import (
	"graphql_test/model"

	"github.com/graphql-go/graphql"
)

// 定义查询对象的字段，支持嵌套
var helloType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Hello",
	Description: "Hello Model",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
	},
})

// 处理查询请求
var queryHello = graphql.Field{
	Name:        "QueryHello",
	Description: "Query Hello",
	Type:        graphql.NewList(helloType),

	// Args 是定义在 GraphQL 查询中支持的查询字段，
	// 可自行随意定义，如加上 limit,start 这类
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},

	// Resolve 是一个处理请求的函数，具体处理逻辑可在此进行
	Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
		// Args 里面定义的字段在 p.Args 里面，对应的取出来
		// 因为是 interface{} 的值，需要类型转换，可参考类型断言(type assertion): https://zengweigang.gitbooks.io/core-go/content/eBook/11.3.html
		id, _ := p.Args["id"].(int)
		name, _ := p.Args["name"].(string)

		// 调用 Hello 这个 model 里面的 Query 方法查询数据
		return (&model.Hello{}).Query(id, name)
	},
}
