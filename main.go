package main

import (
	"github.com/durotimicodes/xanda_task_R3_D3/cmd"
	"github.com/durotimicodes/xanda_task_R3_D3/cmd/database"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	database.InitDatabase()
	cmd.StartApi()

}
