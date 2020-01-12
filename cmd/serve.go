package cmd

import (
	"context"
	"fmt"
	"github.com/Zygimantass/beer-backend/api"
	"github.com/Zygimantass/beer-backend/app"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
)

func serveAPI(ctx context.Context, api *api.API) {
	corsRules := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: true,
		MaxAge: 300,
	})

	router := chi.NewRouter()
	router.Use(corsRules.Handler)
	router.Use(middleware.Logger)

	router.Mount("/api/v1", api.Init())

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", api.Config.Port),
		Handler: router,
	}

	done := make(chan struct{})

	go func() {
		<-ctx.Done()

		if err := server.Shutdown(context.Background()); err != nil {
			log.Fatal(err.Error())
		}

		log.Println("beer-backend service is booting down")
		close(done)
	}()

	log.Printf("serving beer-backend at port %d\n", api.Config.Port)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err.Error())
	}

	<-done
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serves the baceknd to a public port",
	RunE: func(cmd *cobra.Command, args []string) error {
		app, err := app.New()
		if err != nil {
			return err
		}

		api, err := api.New(app)
		if err != nil {
			return err
		}

		ctx, cancel := context.WithCancel(context.Background())

		go func() {
			ch := make(chan os.Signal, 1)
			signal.Notify(ch, os.Interrupt)
			<-ch
			cancel()
		}()

		var wg sync.WaitGroup

		wg.Add(1)

		go func() {
			defer wg.Done()
			defer cancel()
			defer serveAPI(ctx, api)
		}()

		wg.Wait()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
