package server

import (
	"github.com/durotimicodes/xanda_task_R3_D3/cmd"
	"github.com/durotimicodes/xanda_task_R3_D3/cmd/database/repository"
	"github.com/durotimicodes/xanda_task_R3_D3/handlers"
)

type Server struct {
	DB     repository.MySQLDb
	Router *cmd.Router
}

func StartServer() {

	mySQL := new(repository.MySQLDb)

	h := &handlers.Handler{
		DB: mySQL,
	}

	cmd.StartApi(h)

}
