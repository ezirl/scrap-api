CREATE TABLE user
(
    id       INT AUTO_INCREMENT PRIMARY KEY,
    email    varchar(255) not null,
    password varchar(255) not null,
    token    varchar(255) not null,
    tariff   varchar(150) DEFAULT 1,
    requests int     DEFAULT 1000
)