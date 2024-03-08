CREATE TABLE IF NOT EXISTS waitlists (
    "id" SERIAL PRIMARY KEY NOT NULL,
    "email" VARCHAR(50) UNIQUE NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT 'now()'
)