package routing

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/RomainC75/todo2/config"
)

func Serve() {
	configs := config.Get()
	r := GetRouter()

	// err := r.Run(fmt.Sprintf(":%v", configs.Server.Port))
	// if err != nil {
	// 	log.Fatal("Error in routing !")
	// 	return
	// }

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", configs.Server.Port),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	<-ctx.Done()
	log.Println("Server exiting")
}
