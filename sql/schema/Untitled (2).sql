CREATE TABLE "document" (
  "documentid" integer PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamp DEFAULT 'now()',
  "createdby" integer
);

CREATE TABLE "users" (
  "id" integer PRIMARY KEY,
  "username" varchar NOT NULL,
  "userhash" varchar NOT NULL,
  "created_at" timestamp DEFAULT 'now()',
  "admin" boolean
);

ALTER TABLE "document" ADD FOREIGN KEY ("createdby") REFERENCES "users" ("id");
