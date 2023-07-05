package helpers

import "log"

func HandlerErr(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}
