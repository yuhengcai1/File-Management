CREATE TABLE "document" (
  "documentid" integer PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamp DEFAULT 'now()',
  "createdby" integer
);

CREATE TABLE "users" (
  "id" integer PRIMARY KEY,
  "username" varchar NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamp DEFAULT 'now()'
);

CREATE TABLE "admin" (
  "id" integer PRIMARY KEY
);

CREATE TABLE "normal" (
  "id" integer PRIMARY KEY,
  "createdby" integer
);

ALTER TABLE "Document" ADD FOREIGN KEY ("createdby") REFERENCES "users" ("id");

ALTER TABLE "admin" ADD FOREIGN KEY ("id") REFERENCES "users" ("id");

ALTER TABLE "normal" ADD FOREIGN KEY ("id") REFERENCES "users" ("id");

ALTER TABLE "normal" ADD FOREIGN KEY ("createdby") REFERENCES "admin" ("id");
