package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"

	psqlrepo "pomodoro.news/timer/internal/adapters/persistence/postgres"

	"github.com/hasanozgan/frodao/drivers/postgres"
	"pomodoro.news/timer/internal/adapters/config"

	"pomodoro.news/timer/internal/application/usecase"

	"pomodoro.news/timer/internal/adapters/graphql"
	"pomodoro.news/timer/internal/adapters/graphql/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	cfg := config.MustLoad("")

	if err := postgres.Connect(cfg.DSN); err != nil {
		log.Fatalf("DB Connection failed %s", cfg.DSN)
	}
	defer postgres.Close()

	pomodoroRepo := psqlrepo.NewPomodoro()

	router := chi.NewRouter()

	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(r.Header.Get("X-user_ID"))
			userID, _ := strconv.Atoi(r.Header.Get("X-USER_ID"))
			ctx := context.WithValue(r.Context(), "userID", userID)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	})

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graphql.Resolver{
		PomodoroUC: usecase.NewPomodoro(pomodoroRepo),
	}}))

	router.Handle("/", playground.Handler("GraphQL", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
