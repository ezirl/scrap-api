package user

import "github.com/dgrijalva/jwt-go"

const (
	FREE       = "free"
	BASE       = "base"
	BUSINESS   = "business"
	ENTERPRISE = "enterprise"
)

type Repository interface {
	FindByID(ID int) (*User, error)
	Save(user *User) error
}

type User struct {
	jwt.StandardClaims
	ID       int    `db:"id" json:"id"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"-"`
	Tariff   string `db:"tariff" json:"tariff"`
	Requests int    `db:"requests" json:"requests"`
}

func (u *User) CheckLimitBan() bool {
	switch u.Tariff {
	case FREE:
		if u.Requests > 999 {
			return true
		}
		return false

	case BASE:
		if u.Requests > 99999 {
			return true
		}
		return false

	case BUSINESS:
		if u.Requests > 2999999 {
			return true
		}
		return false

	case ENTERPRISE:
		if u.Requests > 19999999 {
			return true
		}
		return false
	}

	return false
}
