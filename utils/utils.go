package uitils

import (
	"errors"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

var ErrAlreadyExists = errors.New("record already exists")
