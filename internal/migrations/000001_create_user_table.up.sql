CREATE TABLE "user" (
    "createdAt"   timestamptz        NOT NULL,
    "updatedAt"   timestamptz        NOT NULL,
    "deletedAt"   timestamptz,
    "id"          UUID             PRIMARY KEY,
    "email"       VARCHAR(50)       NOT NULL,
    "nickname"    VARCHAR(100)     NOT NULL,
    "providers"   text[]             NOT NULL,
    CONSTRAINT "uniq_email" UNIQUE ("email")
);