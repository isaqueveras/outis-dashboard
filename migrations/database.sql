-- DROP TABLE IF EXISTS t_watcher, t_routine, t_routine_execution, t_routine_log, t_routine_indicator;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE t_watcher (
	"id" UUID PRIMARY KEY,
	"name" TEXT NOT NULL CHECK (CHAR_LENGTH("name") > 10 AND CHAR_LENGTH("name") <= 150) UNIQUE,
	"last_run" TIMESTAMP NOT NULL,
	"created_at" TIMESTAMP NOT NULL
);

CREATE TABLE t_routine (
	"id" UUID PRIMARY KEY,
  "id_watcher" UUID NOT NULL REFERENCES t_watcher (id),
	"name" TEXT NOT NULL CHECK (CHAR_LENGTH("name") > 10 AND CHAR_LENGTH("name") <= 150) UNIQUE,
	"desc" VARCHAR,
	"state" VARCHAR NOT NULL,
  "interval" INTEGER NOT NULL,
	"load_interval" INTEGER,
	"start_hour" INTEGER,
	"end_time" INTEGER,
	"last_run" TIMESTAMP,
	"updated_at" TIMESTAMP NOT NULL,
	"created_at" TIMESTAMP NOT NULL
);

CREATE TABLE t_routine_execution (
	"id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "id_routine" UUID NOT NULL REFERENCES t_routine (id),
	"latency" FLOAT NOT NULL,
  "metadata" JSONB,
	"start" TIMESTAMP NOT NULL,
	"end" TIMESTAMP NOT NULL
);

CREATE TABLE t_routine_indicator (
  "id_routine_execution" UUID NOT NULL REFERENCES t_routine_execution (id),
	"key" TEXT NOT NULL,
	"value" TEXT NOT NULL,
	"created_at" TIMESTAMP NOT null DEFAULT NOW()
);

CREATE TABLE t_routine_histogram (
  "id_routine_execution" UUID NOT NULL REFERENCES t_routine_execution (id),
	"key" TEXT NOT NULL,
	"value" TEXT NOT NULL,
	"created_at" TIMESTAMP NOT null DEFAULT NOW()
);

CREATE TABLE t_routine_log (
	"id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "id_routine_execution" UUID NOT NULL REFERENCES t_routine_execution (id),
  "level" VARCHAR NOT NULL,
	"msg" VARCHAR NOT NULL,
	"ts" TIMESTAMP NOT NULL,
	"path" TIMESTAMP NOT NULL
);

-- DROP FUNCTION set_date_update();
-- DROP TRIGGER t_routine_set_date_update ON t_routine;

CREATE OR REPLACE FUNCTION set_date_update()
 RETURNS trigger
 LANGUAGE plpgsql
AS $function$
BEGIN
  IF (TG_OP = 'UPDATE') THEN
    NEW.updated_at = now();
    RETURN NEW;
  END IF;
END;
$function$;

CREATE TRIGGER t_routine_set_date_update BEFORE UPDATE ON t_routine FOR EACH ROW EXECUTE PROCEDURE set_date_update();
