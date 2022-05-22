package application

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"pomodoro.news/timer/internal/adapters/config"
)

func Run() {
	configPath := parseFlags()
	cfg := config.MustLoad(configPath)

	ctx, ctxCancelFunc := context.WithTimeout(context.Background(), cfg.Timeout)
	defer ctxCancelFunc()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(time.Now().String()))

	})

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: mux,
	}

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)

	go func() {
		log.Printf("%s listening on 0.0.0.0:%d with %d timeout", cfg.ServiceName, cfg.Port, cfg.Timeout)
		if err := srv.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}
	}()

	<-stop

	log.Printf("%s shutting down ...\n", cfg.ServiceName)

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Printf("%s down\n", cfg.ServiceName)
}
