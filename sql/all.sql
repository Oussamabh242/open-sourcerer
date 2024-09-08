
CREATE TABLE articles(
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  title text NOT NULL,
  content text NOT NULL ,
  created_at text NOT NULL , 
  view_count INTEGER DEFAULT 0  
)


CREATE TABLE confirmations (
    email TEXT NOT NULL PRIMARY KEY,
    id TEXT NOT NULL,
    expires DATETIME
);

CREATE TABLE subscribe (
  id INTEGER PRIMARY KEY AUTOINCREMENT ,
  email TEXT NOT NULL 
);

CREATE TRIGGER set_expiration_time
AFTER INSERT ON confirmations
FOR EACH ROW
WHEN NEW.expires IS NULL
BEGIN
    UPDATE confirmations
    SET expires = DATETIME('now', '+1 day')
    WHERE email = NEW.email;
END;

CREATE TRIGGER delete_expired_rows_after_insert_or_delete
AFTER INSERT ON confirmations
BEGIN
  DELETE FROM confirmations
  WHERE expires < DATETIME('now');
END;



