package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shurcooL/graphql"
)

var query struct {
	Posts []struct {
		Id    graphql.Int
		Title graphql.String
		Cover graphql.String
	}
}

func GetPosts(ctx *gin.Context) {

	client := graphql.NewClient("https://mockend.com/mockend/demo/graphql", nil)
	err := client.Query(context.Background(), &query, nil)
	if err != nil {
		fmt.Println("An error has occurred when querying graphql server")
		fmt.Println(err)
	}
	ctx.JSON(http.StatusOK, query.Posts)
}
