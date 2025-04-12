package main

import (
	"log"
	"net/http"
	"os"

	"github.com/astaxie/beego/orm"
	"github.com/greedy_game/targeting_engine/service"
	"github.com/greedy_game/targeting_engine/transport"
	"github.com/greedy_game/targeting_engine/util"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)
	db := orm.NewOrm()

	model := service.NewDatabseModel(
		&db,
	)

	//initialize the database
	util.Init()

	// Create service

	svc := service.NewService(logger, model)

	// Create HTTP handler
	handler := transport.NewHTTPHandler(svc)

	// Start server
	logger.Print("Starting server on :8080")
	logger.Fatal(http.ListenAndServe(":8080", handler))
}
