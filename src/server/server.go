package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/esc-chula/gearfest-backend/src/config"
	supa "github.com/nedpals/supabase-go"
	"gorm.io/gorm"
)

type Server struct {
	router http.Handler
	db     *gorm.DB
}

func New() *Server {
	config, err := config.New()
	if err != nil {
		fmt.Println("Error reading config.")
		os.Exit(0)
	}

	supabase := supa.CreateClient(config.SupabaseClientConfig.URL, config.SupabaseClientConfig.Key)

	db := LoadSupabase(config.SupabaseDatabaseConfig)

	server := &Server{
		router: loadRoutes(db, supabase),
		db:     db,
	}
	return server
}

func (s *Server) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":8080",
		Handler: s.router,
	}

	fmt.Println("Checking database connection.")
	sqlDB, err := s.db.DB()
	if err != nil {
		fmt.Println("Failed to get database: ", err)
		return err
	} else if err = sqlDB.Ping(); err != nil {
		fmt.Println("Failed to check database connection: ", err)
		return err
	}
	fmt.Println("Database connected.")

	defer func() {
		defer sqlDB.Close()
	}()

	w := make(chan error, 1)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			w <- err
		}
		close(w)
	}()

	select {
	case e := <-w:
		fmt.Println("Error during Listen and Serve:", e)
		return e
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*8)
		defer cancel()
		fmt.Println("Graceful shutdown.")
		err := server.Shutdown(timeout)
		if err != nil {
			fmt.Println("Error during graceful shutdown:", err)
			return err
		}
		return nil
	}
}
