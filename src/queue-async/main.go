package main

import (
	"Go/src/queue-async/queue2"
	"github.com/gin-gonic/gin"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	server *http.Server
	osSignal chan os.Signal
)

func main() {

	osSignal = make(chan os.Signal, 10000)
	signal.Notify(osSignal, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("In app queue system")

	emailQueue := queue2.NewEmailQueue()

	appEngine := gin.Default()
	appEngine.POST("/users", func(c *gin.Context) {

		emailQueue.Enqueue("Send email to the user")

		c.JSON(http.StatusCreated, gin.H{
			"data": gin.H{
				"username": "user1",
				"email":    "user1@gamil.com",
			},
			"message": "success create new users",
			"status":  http.StatusCreated,
		})

	})

	server = &http.Server{
		Addr:    ":8080",
		Handler: appEngine,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Unexpected server error because of: %v\n", err)
		}
	}()

	for i := 0 ; i < 10 ; i ++ {
		go emailQueue.Work()
	}
	<-osSignal

	fmt.Println("Terminating server")
	server.Shutdown(context.Background())

	fmt.Println("Terminating email queue")

	for emailQueue.Size() > 0 {
		time.Sleep(time.Millisecond * 500)
	}

	fmt.Println("Complete terminating application")

}
