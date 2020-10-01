CREATE EXTENSION IF NOT EXISTS pg_trgm;
CREATE TABLE "user" (
	"id" UUID PRIMARY KEY,
	"lastName" VARCHAR(64) NOT NULL,
	"lastName" VARCHAR(64) NOT NULL,
	"joinedAt" timestamptz NOT NULL,
	"phoneNumber" VARCHAR(10) NOT NULL
);


