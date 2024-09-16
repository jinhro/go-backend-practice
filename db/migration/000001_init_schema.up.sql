CREATE TABLE IF NOT EXISTS "accounts" (
    "id" BIGSERIAL PRIMARY KEY NOT NULL,
    "owner" VARCHAR NOT NULL,
    "balance" BIGINT NOT NULL,
    "currency" VARCHAR NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "entries" (
    "id" BIGSERIAL PRIMARY KEY NOT NULL,
    "account_id" BIGINT NOT NULL,
    "amount" BIGINT NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "transfers" (
    "id" BIGSERIAL PRIMARY KEY NOT NULL,
    "from_account_id" BIGINT NOT NULL,
    "to_account_id" BIGINT NOT NULL,
    "amount" BIGINT NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT(NOW())
);

CREATE INDEX ON "accounts" ("owner");

CREATE INDEX ON "entries" ("account_id");

CREATE INDEX ON "transfers" ("from_account_id");

CREATE INDEX ON "transfers" ("to_account_id");

CREATE INDEX ON "transfers" ("from_account_id", "to_account_id");

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");