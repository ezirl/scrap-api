package user

import "database/sql"

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) Repo {
	return Repo{
		db: db,
	}
}

func (r *Repo) Save(user *User) error {
	_, err := r.db.Exec("INSERT INTO user (email, password, token) VALUES (?,?,?)",
		user.Email, user.Password, user.Token)
	return err
}

func (r *Repo) FindByID(ID int) (*User, error) {
	row := r.db.QueryRow("SELECT * FROM user where id = ?", ID)

	user := User{}
	_ = row.Scan(&user.ID, &user.Email, &user.Password, &user.Token, &user.Tariff, &user.Requests)

	return &user, nil
}

func (r *Repo) FindByToken(token string) (*User, error) {
	row := r.db.QueryRow("SELECT * FROM user where token = ?", token)

	user := User{}
	_ = row.Scan(&user.ID, &user.Email, &user.Password, &user.Token, &user.Tariff, &user.Requests)

	return &user, nil
}

func (r *Repo) IncRequests(id int) bool {
	_, err := r.db.Exec("UPDATE scrap.user set requests = requests + 1 where id = ?", id)
	_err(err)
	return true
}

func _err(err error) {
	if err != nil {
		panic(err)
	}
}
