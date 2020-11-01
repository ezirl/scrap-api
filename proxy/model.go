package proxy

type Repository interface {
	FindByID(ID int) (*Proxy, error)
	Save(proxy *Proxy) error
}

type Proxy struct {
	ID       int    `db:"id" json:"id"`
	Address  string `db:"address" json:"address"`
	Port     string `db:"port" json:"port"`
	Login    string `db:"login" json:"login"`
	Password string `db:"password" json:"password"`
	Country  string `db:"country" json:"country"`
	Premium  bool   `db:"premium" json:"premium"`
	Type     string `db:"type" json:"type"`
}
