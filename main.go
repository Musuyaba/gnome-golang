package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"time"

	"github.com/Musuyaba/gnome-golang/pkg/initializers"
	"github.com/Musuyaba/gnome-golang/sqlc/generated"
	_ "github.com/go-sql-driver/mysql"
)

var (
	// server *gin.Engine
	err error
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDbPostgreSQL(&config)

	// server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}
	ctx := context.Background()
	db, err := sql.Open("mysql", config.DATABASE_MYSQL_USER+":"+config.DATABASE_MYSQL_PASSWORD+"@("+config.DATABASE_MYSQL_HOST+":"+config.DATABASE_MYSQL_PORT+")/"+config.DATABASE_MYSQL_NAME+"?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	sqlc_query := generated.New(db)
	startTime := time.Now()
	rowPallete, err := sqlc_query.GetPalleteNotDone(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(rowPallete.ID)

	err = sqlc_query.UpdatePalleteToDone(ctx, rowPallete.ID)
	if err != nil {
		log.Fatal(err)
	}
	executionTime := time.Since(startTime)
	fmt.Println("Execution time:", executionTime)

	// routers.ApiRoutes(server)
	// routers.PublicRoutes(server)

	// log.Fatal(server.RunTLS(config.HOSTNAME+":"+config.PORT, "./certs/generated/server.crt", "./certs/generated/server.key"))
}
