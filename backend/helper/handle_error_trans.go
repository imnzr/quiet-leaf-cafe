package helper

import "log"

func HandleErrorTransaction(err error) {
	if err != nil {
		log.Println("failed to commit transaction:", err)
	}
}
