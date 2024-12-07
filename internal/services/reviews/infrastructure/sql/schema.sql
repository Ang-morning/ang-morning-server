-- sqlc에서는 버저닝을 할 수 없기 때문에 migration 파일은 따로 관리해야한다.
CREATE TABLE "review" (
    "createdAt"   timestamptz        NOT NULL,
    "updatedAt"   timestamptz        NOT NULL,
    "id"          UUID             PRIMARY KEY,
    "userId"      UUID             NOT NULL,
    "hospitalId"   UUID             NOT NULL,
    "content"     TEXT             NOT NULL,
    "rating"      INT              NOT NULL
);