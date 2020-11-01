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
	_, err := r.db.Exec("INSERT INTO proxy (login, password, address, port, country, type) VALUES (?,?,?,?,?,?)",
		proxy.Login, proxy.Password, proxy.Address, proxy.Port, proxy.Country, proxy.Type)
	return err
}

func (r *Repo) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM proxy WHERE id = ?", id)
	return err
}

func (r *Repo) All() (*[]Proxy, error) {
	row, _ := r.db.Query("SELECT * FROM proxy")

	var proxies []Proxy
	for row.Next() {
		var proxy Proxy
		_ = row.Scan(&proxy.ID, &proxy.Login, &proxy.Password, &proxy.Port, &proxy.Address, &proxy.Country, &proxy.Premium, &proxy.Type)
		proxies = append(proxies, proxy)
	}

	return &proxies, nil
}

func (r *Repo) FindByID(ID int) (*Proxy, error) {
	row := r.db.QueryRow("SELECT * FROM proxy where id = ?", ID)

	proxy := Proxy{}
	_ = row.Scan(&proxy.ID, &proxy.Login, &proxy.Password, &proxy.Port, &proxy.Address, &proxy.Country, &proxy.Premium, &proxy.Type)

	return &proxy, nil
}

func (r *Repo) GetRandom() (*Proxy, error) {
	row := r.db.QueryRow("SELECT * FROM proxy order by rand() limit 1")

	proxy := Proxy{}
	err := row.Scan(&proxy.ID, &proxy.Login, &proxy.Password, &proxy.Port, &proxy.Address, &proxy.Country, &proxy.Premium, &proxy.Type)

	if err != nil {
		return &proxy, err
	}

	return &proxy, nil
}

func (r *Repo) GetPremium() (*Proxy, error) {
	row := r.db.QueryRow("SELECT * FROM proxy where premium = true order by rand() limit 1")

	proxy := Proxy{}
	err := row.Scan(&proxy.ID, &proxy.Login, &proxy.Password, &proxy.Port, &proxy.Address, &proxy.Country, &proxy.Premium, &proxy.Type)

	if err != nil {
		return nil, err
	}

	return &proxy, nil
}

func (r *Repo) FindByCountry(country string) (*Proxy, error) {
	row := r.db.QueryRow("SELECT * FROM proxy where country = ? order by rand() limit 1", country)

	proxy := Proxy{}
	err := row.Scan(&proxy.ID, &proxy.Login, &proxy.Password, &proxy.Port, &proxy.Address, &proxy.Country, &proxy.Premium, &proxy.Type)

	if err != nil {
		return nil, err
	}

	return &proxy, nil
}
