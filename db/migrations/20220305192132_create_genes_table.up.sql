CREATE TABLE "genes" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "type" varchar NOT NULL,
  "produced_name" varchar NOT NULL,
  "produced_date" date NOT NULL,
  "availability" varchar NOT NULL,
  "description" text NULL,
  "history" text NULL,
  "links" jsonb,

  "created_at" timestamptz NOT NULL DEFAULT (now())
);