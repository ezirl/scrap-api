package scrap

import "github.com/07sima07/scrap-api/user"

type Repository interface {
	FindByID(ID int) (*Call, error)
	Save(call *Call) error
}

type Call struct {
	ID     int    `db:"id" json:"id"`
	UserID int    `db:"user_id" json:"user_id"`
	Url    string `db:"url" json:"url"`
	User   *user.User `json:"user"`
}
