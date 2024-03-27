package unibookBackend

import (
	"log"
)

func Log(msg ...interface{}) {
	if !IS_DEBUG {
		return
	}
	msg = append([]interface{}{"UNIBOOK-BACKEND: "}, msg...)
	log.Println(msg...)
}
