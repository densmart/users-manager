package main

import (
	"context"
	"github.com/densmart/users-manager/internal/adapters/api/rest"
	"github.com/densmart/users-manager/internal/adapters/db"
	"github.com/densmart/users-manager/internal/configger"
	"github.com/densmart/users-manager/internal/domain/repo"
	"github.com/densmart/users-manager/internal/domain/services"
	"github.com/densmart/users-manager/internal/logger"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	configger.InitConfig()
	logger.InitLogger()
	logger.Infof("logger initialized. level: %s", logger.GetLevel())

	logger.Infof("starting app...")

	// init app WG and context
	appCtx, cancel := context.WithCancel(context.Background())
	appWg := new(sync.WaitGroup)

	// create DB connection
	logger.Infof("DB connection -> %s@%s:%s/%s", viper.GetString("db.username"),
		viper.GetString("db.host"), viper.GetString("db.port"), viper.GetString("db.dbname"))
	dbWrapper, err := db.NewDB(appCtx, "postgres", db.ConnectionConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslMode"),
	})
	if err != nil {
		logger.Fatalf("error starting DB: %s", err.Error())
	}

	// create new repo instance
	r := repo.NewRepo(dbWrapper)

	// initialize services
	s := services.NewService(r)

	// run migrations
	if err = s.Migrator.Up(); err != nil {
		logger.Fatalf("error DB migrate: %s", err.Error())
	}

	// initialize REST (router & server)
	restRouter := rest.NewRestRouter(s)
	restServer := rest.NewRestServer(restRouter.InitRoutes())
	// start REST http server
	go func() {
		if err = restServer.Run(); err != nil {
			logger.Errorf("error starting HTTP server: %s", err.Error())
		}
	}()

	// catch term OS signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	logger.Infof("app started!")
	// wait for interruption
	<-quit

	// Begin gracefully shutdown
	logger.Infof("stopping app...")

	// send stop signal to all goroutines
	cancel()
	// wait for all goroutines finished
	appWg.Wait()

	// close DB connection
	dbWrapper.Close()
	logger.Infof("all DB connections closed")

	// all components stopped successfully!
	logger.Infof("app stopped!")
}
