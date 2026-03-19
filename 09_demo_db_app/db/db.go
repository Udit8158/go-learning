package db

type UserDetails struct {
	Name  string
	Email string
	Todos []string
}

type users map[string]UserDetails

type ErrUserExist struct {
	File string
	Msg  string
}

func (e ErrUserExist) Error() string {
	return e.Msg
}

var UserTable = make(users)

func (u users) AddUser(details UserDetails) error {
	// search for the email
	_, ok := u[details.Email]
	if !ok {
		u[details.Email] = details
		return nil
	}
	// fmt.Printf("ERROR occurs %#v \n", ErrUserExist{File: "db.go", Msg: "User already exist in the db table"})
	return &ErrUserExist{File: "db.go", Msg: "User already exist in the db table"}

}
