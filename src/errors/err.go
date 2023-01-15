package errors

import "log"

func CheckErr(err error) {
	if err != nil {
		log.Default().Print(err)
	}
}
