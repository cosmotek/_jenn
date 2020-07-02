CREATE EXTENSION IF NOT EXISTS pg_trgm;
CREATE TABLE "user" (
	"id" UUID PRIMARY KEY,
	"_archived" BOOLEAN NOT NULL DEFAULT FALSE,
	"firstName" VARCHAR(64) NOT NULL,
	"lastName" VARCHAR(64) NOT NULL,
	"joinedAt" timestamptz NOT NULL,
	"phoneNumber" VARCHAR(10) NOT NULL
);


CREATE TABLE "cocktail" (
	"id" UUID PRIMARY KEY,
	"_archived" BOOLEAN NOT NULL DEFAULT FALSE,
	"name" VARCHAR(512) NOT NULL
);


CREATE INDEX index_cocktail_on_name_trigram
ON "cocktail"
USING gin ("name" gin_trgm_ops);


CREATE TABLE "beverage" (
	"id" UUID PRIMARY KEY,
	"_archived" BOOLEAN NOT NULL DEFAULT FALSE,
	"name" VARCHAR(512) NOT NULL,
	"proof" smallint NOT NULL,
	"type" VARCHAR(6) NOT NULL
);


CREATE INDEX index_beverage_on_name_trigram
ON "beverage"
USING gin ("name" gin_trgm_ops);


CREATE INDEX index_beverage_on_proof_trigram
ON "beverage"
USING BTREE ("proof");


CREATE INDEX index_beverage_on_type_trigram
ON "beverage"
USING BTREE ("type");


