
-- migrate up
CREATE SCHEMA IF NOT EXISTS "example";

-- GLOBAL
CREATE OR REPLACE FUNCTION example.function_updated_at()
  RETURNS TRIGGER AS $$
  BEGIN
   NEW.updated_at = now();
   RETURN NEW;
  END;
  $$ LANGUAGE 'plpgsql';


-- PROCESS
CREATE TABLE example.person (
  id_person               SERIAL NOT NULL,
  name                    TEXT NOT NULL,
  age                     INTEGER,
  created_at              TIMESTAMP DEFAULT NOW(),
  updated_at              TIMESTAMP DEFAULT NOW(),
  CONSTRAINT example_id_person_pkey PRIMARY KEY (id_person)
);

CREATE TRIGGER trigger_person_updated_at BEFORE UPDATE
  ON example.person FOR EACH ROW EXECUTE PROCEDURE example.function_updated_at();


-- history
CREATE TABLE example.person_history (LIKE example.person);
ALTER TABLE example.person_history ADD COLUMN _operation TEXT NOT NULL;
ALTER TABLE example.person_history ADD COLUMN "_user" TEXT NOT NULL;
ALTER TABLE example.person_history ADD COLUMN _operation_at TIMESTAMP DEFAULT NOW();

CREATE OR REPLACE FUNCTION function_person_history() RETURNS TRIGGER AS $$
BEGIN
    IF (TG_OP = 'DELETE') THEN
        INSERT INTO example.person_history VALUES(OLD.*, 'D', user, now());
        RETURN OLD;
    ELSIF (TG_OP = 'UPDATE') THEN
        INSERT INTO example.person_history VALUES(NEW.*, 'U', user, now());
        RETURN NEW;
    ELSIF (TG_OP = 'INSERT') THEN
        INSERT INTO example.person_history VALUES(NEW.*, 'I', user, now());
        RETURN NEW;
    END IF;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_person_history
AFTER INSERT OR UPDATE OR DELETE ON example.person
    FOR EACH ROW EXECUTE PROCEDURE function_person_history();


-- migrate down
DROP TRIGGER trigger_person_history ON example.person;
DROP TRIGGER trigger_person_updated_at on example.person;
DROP FUNCTION example.function_updated_at();

DROP TABLE example.person_history;
DROP TABLE example.person;