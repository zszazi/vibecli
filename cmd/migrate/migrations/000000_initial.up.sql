-- I copied this code from https://github.com/golang-migrate/migrate/issues/179#issuecomment-479740765 , since go-migrate doesn't have a migration_history table like flyway
BEGIN;

SET ROLE 'admin';

CREATE TABLE IF NOT EXISTS schema_migrations_history(        
    id SERIAL PRIMARY KEY NOT NULL,
    version BIGINT NOT NULL,
    applied_at timestamptz NOT NULL DEFAULT NOW()
);

CREATE OR REPLACE FUNCTION track_applied_migration()
RETURNS TRIGGER AS $$
DECLARE _current_version integer;
BEGIN
    SELECT COALESCE(MAX(version),0) FROM schema_migrations_history INTO _current_version;
    IF new.dirty = 'f' AND new.version > _current_version THEN
        INSERT INTO schema_migrations_history(version) VALUES (new.version);
    END IF;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- TRIGGER
-- TODO : handle if not exists for trigger
CREATE TRIGGER rack_applied_migrations AFTER INSERT ON schema_migrations FOR EACH ROW EXECUTE PROCEDURE track_applied_migration();

COMMIT;