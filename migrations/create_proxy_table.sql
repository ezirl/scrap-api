CREATE TABLE proxy
(
    id       INT AUTO_INCREMENT PRIMARY KEY,
    address  varchar(255) not null,
    country  varchar(5),
    premium  bool default false
)