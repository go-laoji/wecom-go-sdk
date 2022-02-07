package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-laoji/wework"
	"github.com/go-laoji/wework/config"
	"github.com/go-laoji/wework/pkg/demo"
	"github.com/go-laoji/wework/pkg/svr"
	"github.com/go-laoji/wework/pkg/svr/logic"
	"github.com/go-laoji/wework/pkg/svr/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	c := config.ParseFile("")
	wwconfig := wework.WeWorkConfig{
		CorpId:              c.CorpId,
		ProviderSecret:      c.ProviderSecret,
		SuiteId:             c.SuiteId,
		SuiteSecret:         c.SuiteSecret,
		SuiteToken:          c.SuiteToken,
		SuiteEncodingAesKey: c.SuiteEncodingAesKey,
		Dsn:                 c.Dsn,
	}

	router := gin.Default()
	ww := wework.NewWeWork(wwconfig)
	logic.Migrate(wwconfig.Dsn)
	router.Use(middleware.InjectSdk(ww))

	svr.InjectRouter(router)
	demo.InjectRouter(router)
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": ww.UserGet(1, "jifengwei")})
	})
	srv01 := &http.Server{
		Addr:           fmt.Sprintf("127.0.0.1:%v", c.Port),
		Handler:        router,
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		if err := srv01.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv01.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown:", err)
	}
}
