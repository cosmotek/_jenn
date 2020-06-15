
CREATE TABLE "user" (
	"id" UUID PRIMARY KEY,
	"firstName" VARCHAR(64) NOT NULL,
	"lastName" VARCHAR(64) NOT NULL,
	"joinedAt" timestamptz NOT NULL,
	"phoneNumber" VARCHAR(10) NOT NULL
);


CREATE TABLE "cocktail" (
	"id" UUID PRIMARY KEY,
	"name" VARCHAR(512) NOT NULL
);


CREATE INDEX CONCURRENTLY index_cocktail_on_name_trigram
ON "cocktail"
USING gin ("name" gin_trgm_ops);


