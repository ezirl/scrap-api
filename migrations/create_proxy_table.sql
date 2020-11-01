CREATE TABLE proxy
(
    id       INT AUTO_INCREMENT PRIMARY KEY,
    login    varchar(255),
    password varchar(255),
    port     int          not null,
    address  varchar(255) not null,
    country  varchar(5) default 'none',
    premium  bool default false,
    type varchar(10) default 'http'
)