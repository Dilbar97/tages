CREATE TABLE IF NOT EXISTS "media" (
  "id"            SERIAL      PRIMARY KEY,
  "name"          varchar     NOT NULL,
  "created_at"    timestamp   DEFAULT now(),
  "updated_at"    timestamp   DEFAULT now()
);