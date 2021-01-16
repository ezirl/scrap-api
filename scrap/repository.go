package scrap

import (
	"database/sql"
	"github.com/07sima07/scrap-api/user"
)

type Repo struct {
	db       *sql.DB
	userRepo user.Repo
}

func NewRepo(db *sql.DB, userRepo user.Repo) Repo {
	return Repo{
		db:       db,
		userRepo: userRepo,
	}
}

func (r *Repo) Save(call Call) error {
	_, err := r.db.Exec("INSERT INTO calls (user_id, url) VALUES (?,?)",
		call.UserID, call.Url)
	return err
}

func (r *Repo) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM calls WHERE id = ?", id)
	return err
}

func (r *Repo) All(limit int64) (*[]Call, error) {
	row, _ := r.db.Query("SELECT * FROM calls ORDER BY id DESC LIMIT ?", limit)

	var calls []Call
	for row.Next() {
		var call Call
		_ = row.Scan(&call.ID, &call.UserID, &call.Url)
		call.User, _ = r.userRepo.FindByID(call.UserID)
		calls = append(calls, call)
	}

	return &calls, nil
}

func (r *Repo) FindByID(ID int) (*Call, error) {
	row := r.db.QueryRow("SELECT * FROM calls where id = ?", ID)

	call := Call{}
	_ = row.Scan(&call.ID, &call.UserID, &call.Url)

	return &call, nil
}
