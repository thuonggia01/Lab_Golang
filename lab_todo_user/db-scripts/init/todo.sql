CREATE TABLE "todos" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint,
  "created_at" timestamp NOT NULL DEFAULT 'now()',
  "updated_at" timestamp NOT NULL DEFAULT 'now()',
  "deleted_at" timestamp DEFAULT 'now()',
  "title" varchar(50) NOT NULL,
  "description" varchar(100) NOT NULL,
  "status" boolean

);


ALTER TABLE "todos" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");