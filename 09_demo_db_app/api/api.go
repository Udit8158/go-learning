package api

import (
	"errors"
	"fmt"

	"github.com/Udit8158/go-learning/09_demo_db_app/db"
)

func AddUserToDB(userDetails db.UserDetails) {
	err := db.UserTable.AddUser(userDetails)

	if err != nil {
		// fmt.Println("ERROR:", err.Error())
		// var userErr *db.ErrUserExist

		// if errors.As(err, &userErr) {
		// 	fmt.Printf("ERROR occured in %q - %q\n", userErr.File, userErr.Msg)
		// }
		//

		e, errTypeMatched := errors.AsType[*db.ErrUserExist](err)
		if errTypeMatched {
			fmt.Printf("ERROR occured in %q - %q\n", e.File, e.Msg)
		}
	}

	x := 10
	p := &x
	fmt.Println(*p)
}
