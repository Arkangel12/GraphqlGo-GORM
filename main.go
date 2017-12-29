package main

import (
	"fmt"
	"net/http"
	_ "github.com/lib/pq"
	"github.com/graphql-go/handler"
	"github.com/Arkangel12/curso/graphql-gorm/config"
	"github.com/Arkangel12/curso/graphql-gorm/models"
	"github.com/Arkangel12/curso/graphql-gorm/schema"
	"context"
)

func init() {
	config.LoadAppConfig()
}

func main() {
	h := handler.New(&handler.Config{
		Schema:   &schema.Schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", graphqlHandler(h))

	fmt.Println("Now server is running on ", config.AppConfig.Server)

	http.ListenAndServe(":8000", nil)
}

func graphqlHandler(h *handler.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Requesting something")

		var ctx context.Context

		viewer := models.User{
			Username: "guest",
		}

		ctx = context.WithValue(context.Background(), "currentUser", viewer)
		h.ContextHandler(ctx, w, r)
	})
}