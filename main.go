package main

import (
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/godcong/wego-auth-manager/config"
	"github.com/godcong/wego-auth-manager/database"
	_ "github.com/godcong/wego-auth-manager/docs" // docs is generated by Swag CLI, you have to import it.
	"github.com/godcong/wego-auth-manager/log"
	"github.com/godcong/wego-auth-manager/model"
	"github.com/godcong/wego-auth-manager/service"
	"os"
	"os/signal"
	"syscall"
)

var configPath = flag.String("config", "config.toml", "load config from path")
var logPath = flag.String("log", "manager.log", "set log name")
var sync = flag.Bool("sync", false, "open to sync the model")
var elk = flag.Bool("elk", true, "set to open the elk")

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @schemes http https

// @license.name MIT
// @license.url https://github.com/godcong/wego-auth-manager/blob/master/LICENSE

// @host localhost:8080
// @BasePath /v0
func main() {
	var e error
	flag.Parse()

	if *elk {
		log.InitLog()
	}

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	cfg := config.InitConfig(*configPath)
	model.InitDB(cfg)

	if *sync {
		e = database.Migrate()
		if e != nil {
			panic(e)
		}
	}

	//start
	service.Start(cfg)

	go func() {
		sig := <-sigs
		fmt.Println(sig, "exiting")
		service.Stop()
		done <- true
	}()
	<-done

}
