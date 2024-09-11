
-- Table to store articles
CREATE TABLE articles(
  id serial PRIMARY KEY,
  title text NOT NULL,
  content text NOT NULL,
  created_at TEXT not null, 
  view_count INTEGER DEFAULT 0
);

-- Table to store confirmations
CREATE TABLE confirmations (
  email TEXT NOT NULL PRIMARY KEY,
  id TEXT NOT NULL,
  expires TIMESTAMPTZ
);

-- Table to store subscribers
CREATE TABLE subscribe (
  id serial PRIMARY KEY,
  email TEXT UNIQUE NOT NULL
);

-- Create a function to set expiration time if not provided
CREATE OR REPLACE FUNCTION set_expiration_time_func()
RETURNS TRIGGER AS $$
BEGIN
  IF NEW.expires IS NULL THEN
    NEW.expires := NOW() + INTERVAL '1 day';
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger to call set_expiration_time_func after insert
CREATE TRIGGER set_expiration_time
BEFORE INSERT ON confirmations
FOR EACH ROW
EXECUTE FUNCTION set_expiration_time_func();

-- Create a function to delete expired rows after insert
CREATE OR REPLACE FUNCTION delete_expired_rows_func()
RETURNS TRIGGER AS $$
BEGIN
  DELETE FROM confirmations
  WHERE expires < NOW();
  RETURN NULL;
END;
$$ LANGUAGE plpgsql;

-- Trigger to call delete_expired_rows_func after insert
CREATE TRIGGER delete_expired_rows_after_insert
AFTER INSERT ON confirmations
EXECUTE FUNCTION delete_expired_rows_func();

