package main

import (
	"github.com/durotimicodes/xanda_task_R3_D3/cmd/database"
	"github.com/durotimicodes/xanda_task_R3_D3/cmd"
)

func main() {

	database.InitDatabase()
	cmd.StartApi()
}
