# go-lang-backend

Backend API in Golang
JWT, MYSQL and Tests;


package user

import (
	"database/sql"
	"fmt"
	"srinjay.com/m/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {

	return &Store{db: db}
}

func (s *Store) getUserEmailId(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * from users WHERE email =?", email)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next(){
		u , err := scanRowsIntoUser(rows)
		if(err != nil) {
			return nil,err
		}
	}

	if u.Id == 0{
		return nil,fmt.Errorf("User not Found!")

	}
	return u,nil;

}

func scanRowsIntoUser(rows *sql.Rows) (*types.User, error){
	user := new(types.User)

	err := rows.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.ModifiedAt,
	)

	if(err != nil) return nil,err;

	return user,nil
}


func (s *Store)  getUserById(id int) (*types.User, error) {

	return nil,nil
}

func (s *Store)  createUser(user types.User) error {

	return nil;
}