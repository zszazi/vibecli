BEGIN;

SET ROLE 'admin';

-- Drop the trigger first since it depends on the function and table
DROP TRIGGER IF EXISTS track_applied_migrations ON schema_migrations;

-- Drop the function that was created
DROP FUNCTION IF EXISTS track_applied_migration();

-- Drop the schema_migrations_history_custom table
DROP TABLE IF EXISTS schema_migrations_history;

COMMIT;
