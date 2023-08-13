package main

import (
	"github.com/durotimicodes/xanda_task_R3_D3/cmd/database"
	"github.com/durotimicodes/xanda_task_R3_D3/server"
)

func main() {

	database.InitDatabase()
	server.StartServer()

}
