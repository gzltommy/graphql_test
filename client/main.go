package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hasura/go-graphql-client"
)

func main() {
	//src := oauth2.StaticTokenSource(
	//	&oauth2.Token{AccessToken: os.Getenv("GRAPHQL_TOKEN")},
	//)
	//httpClient := oauth2.NewClient(context.Background(), src)

	//client := graphql.NewClient("https://example.com/graphql", httpClient)
	// Use client...

	client := graphql.NewClient("https://hub.snapshot.org/graphql", nil)

	// 2.Arguments and Variables
	//{
	//	{
	//		/*
	//			query Spaces {
	//			  spaces(first: 20, skip: 0, orderBy: "created", orderDirection: desc) {
	//			    id
	//			    name
	//			}
	//
	//		*/
	//		var query struct {
	//			Spaces []struct {
	//				ID   graphql.String
	//				Name graphql.String
	//			} `graphql:"spaces(first: 20, skip: 0, orderBy: \"created\", orderDirection: desc)"`
	//		}
	//
	//		err := client.Query(context.Background(), &query, nil)
	//		if err != nil {
	//			fmt.Println("==========", err)
	//			return
	//		}
	//
	//		fmt.Printf("--------- %+v", query)
	//	}
	//
	//	// 变量参数
	//	{
	//		/*
	//				query {
	//					human(id: "1000") {
	//						name
	//						height(unit: METER)
	//				}
	//			}
	//
	//		*/
	//		var query struct {
	//			Spaces []struct {
	//				ID   graphql.String
	//				Name graphql.String
	//			} `graphql:"spaces(first: $first, skip: $skip, orderBy: \"created\", orderDirection: desc)"`
	//		}
	//
	//		variables := map[string]interface{}{
	//			"first": graphql.Int(20),
	//			"skip":  graphql.Int(0),
	//		}
	//
	//		err := client.Query(context.Background(), &query, variables)
	//		if err != nil {
	//			fmt.Println("==========", err)
	//			return
	//		}
	//
	//		buf, _ := json.Marshal(query)
	//		fmt.Printf("---------%s  \n", string(buf))
	//	}
	//}

	// 3.Custom scalar tag

	// 4.Mutations
	{
		/*
			mutation($ep: Episode!, $review: ReviewInput!) {
				createReview(episode: $ep, review: $review) {
					stars
					commentary
				}
			}
			variables {
				"ep": "JEDI",
				"review": {
					"stars": 5,
					"commentary": "This is a great movie!"
				}
			}
		*/

		var mutations struct {
			CreateReview struct {
				Stars      graphql.Int
				Commentary graphql.String
			} `graphql:"createReview(episode: $ep, review: $review)"`
		}

		variables := map[string]interface{}{
			"ep": starwars.Episode("JEDI"),
			"review": starwars.ReviewInput{
				Stars:      graphql.Int(5),
				Commentary: graphql.String("This is a great movie!"),
			},
		}

		err := client.Mutate(context.Background(), &mutations, variables)
		if err != nil {
			// Handle error.
		}
		fmt.Printf("Created a %v star review: %v\n", mutations.CreateReview.Stars, mutations.CreateReview.Commentary)

		// Output:
		// Created a 5 star review: This is a great movie!

	}

}
