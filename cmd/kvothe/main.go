//Package main is the core package of kvothe-server
package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/SurgicalSteel/kvothe/router"
	"github.com/SurgicalSteel/kvothe/util"
)

const appName string = "kvothe"

//func main is the main function of kvothe-server
func main() {
	log.SetOutput(os.Stdout)
	environmentFlag := flag.String("env", "dev", "specify the environment for the app to run. Default : dev")

	flag.Parse()

	err := util.CheckEnvironment(*environmentFlag)
	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Printf("Starting %s in environment %s\n", appName, *environmentFlag)

	routing := router.InitializeRoute()
	routing.RegisterHandler()

	server := &http.Server{
		Addr: "0.0.0.0:9000",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      routing.Router,
	}

	log.Printf("Starting %s in port 9000...\n", appName)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln("Failed to start service! Reason :", err.Error())
		}
	}()
	log.Println("Server Started")

	<-done
	log.Println("Server Stopped")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer func() {
		// extra handling here
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v\n", err)
	}

	log.Println("Server Exited Properly")
	log.Println("👋")
}
