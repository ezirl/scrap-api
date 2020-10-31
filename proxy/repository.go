package proxy

import (
	"database/sql"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) Repo {
	return Repo{
		db: db,
	}
}

func (r *Repo) Save(proxy Proxy) error {
	_, err := r.db.Exec("INSERT INTO proxy (login, password, address, port) VALUES (?,?,?,?)",
		proxy.Login, proxy.Password, proxy.Address, proxy.Port)
	return err
}

func (r *Repo) FindByID(ID int) (*Proxy, error) {
	row := r.db.QueryRow("SELECT * FROM proxy where id = ?", ID)

	proxy := Proxy{}
	_ = row.Scan(&proxy.ID, &proxy.Login, &proxy.Password, &proxy.Port, &proxy.Address, &proxy.Country, &proxy.Premium)

	return &proxy, nil
}

func (r *Repo) GetRandom() (*Proxy, error) {
	row := r.db.QueryRow("SELECT * FROM proxy order by rand() limit 1")

	proxy := Proxy{}
	err := row.Scan(&proxy.ID, &proxy.Login, &proxy.Password, &proxy.Port, &proxy.Address, &proxy.Country, &proxy.Premium)

	if err != nil {
		return &proxy, err
	}

	return &proxy, nil
}

func (r *Repo) GetPremium() (*Proxy, error) {
	row := r.db.QueryRow("SELECT * FROM proxy where premium = true order by rand() limit 1")

	proxy := Proxy{}
	err := row.Scan(&proxy.ID, &proxy.Login, &proxy.Password, &proxy.Port, &proxy.Address, &proxy.Country, &proxy.Premium)

	if err != nil {
		return nil, err
	}

	return &proxy, nil
}

func (r *Repo) FindByCountry(country string) (*Proxy, error) {
	row := r.db.QueryRow("SELECT * FROM proxy where country = ? order by rand() limit 1", country)

	proxy := Proxy{}
	err := row.Scan(&proxy.ID, &proxy.Login, &proxy.Password, &proxy.Port, &proxy.Address, &proxy.Country, &proxy.Premium)

	if err != nil {
		return nil, err
	}

	return &proxy, nil
}
