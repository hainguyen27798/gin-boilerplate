package helpers

import "github.com/hainguyen27798/gin-boilerplate/global"

func Must(err error) {
	if err != nil {
		global.Logger.Error("Details: " + err.Error())
		panic(err)
	}
}
