CREATE TABLE calls (
  id INT AUTO_INCREMENT PRIMARY KEY,
  user_id INT,
  url VARCHAR(5000),
  FOREIGN KEY (user_id) REFERENCES user(id)
)

