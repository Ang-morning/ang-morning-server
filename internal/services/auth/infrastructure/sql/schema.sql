-- sqlc에서는 버저닝을 할 수 없기 때문에 migration 파일은 따로 관리해야한다.
CREATE TABLE "refreshToken" (
    "createdAt" timestamptz NOT NULL,
    "updatedAt" timestamptz NOT NULL,
    "id" SERIAL PRIMARY KEY,
    "userId" UUID NOT NULL,
    "value" VARCHAR(255) NOT NULL,
    "clientInfo" VARCHAR(255)
);