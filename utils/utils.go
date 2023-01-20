package uitils

import (
	"errors"
	"log"
)

func CheckErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

var ErrAlreadyExists = errors.New("record already exists")
