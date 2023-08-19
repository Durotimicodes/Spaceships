package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/durotimicodes/xanda_task_R3_D3/cmd"
	"github.com/durotimicodes/xanda_task_R3_D3/cmd/database"
	"github.com/durotimicodes/xanda_task_R3_D3/cmd/database/repository"
	"github.com/durotimicodes/xanda_task_R3_D3/handlers"
)

type Server struct {
	DB     repository.MySQLDb
	Router *cmd.Router
}

func StartServer() error {

	value := database.InitDBParams()

	mysql := new(repository.MySQLDb)
	h := &handlers.Handler{DB: mysql}
	err := mysql.Init(value.Host, value.User, value.Password, value.DBname, value.Port)
	if err != nil {
		log.Println("Error trying to Init db", err)
		return err
	}
	route, port := cmd.StartApi(h)
	fmt.Println("connect on port", port)
	err = http.ListenAndServe(port, route)
	if err != nil {
		log.Printf("Error from Setting up router %v", err)
		return err
	}

	return nil
}
