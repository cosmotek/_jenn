
CREATE TABLE "user" (
	"firstName" VARCHAR(64) NOT NULL,
	"lastName" VARCHAR(512) NOT NULL,
	"phoneNumber" VARCHAR(10) NOT NULL,
	"joinedAt" timestamptz NOT NULL
);


CREATE INDEX CONCURRENTLY index_user_on_firstName_trigram
ON "user"
USING gin ("firstName" gin_trgm_ops);


