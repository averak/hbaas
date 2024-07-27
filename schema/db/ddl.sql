-----------------------------------
-- ユーザ
-----------------------------------

CREATE TABLE "users"
(
    "id"         UUID         NOT NULL,
    "email"      VARCHAR(255) NOT NULL,
    "status"     INTEGER      NOT NULL,
    "is_deleted" BOOLEAN      NOT NULL,
    "created_at" TIMESTAMP    NOT NULL,
    "updated_at" TIMESTAMP    NOT NULL,
    PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "uq__users__email" ON "users" ("email") WHERE "is_deleted" = FALSE;

CREATE TABLE "user_authentications"
(
    "user_id"               UUID         NOT NULL,
    "baas_user_id"          VARCHAR(255) NOT NULL,
    "last_authenticated_at" TIMESTAMP    NOT NULL,
    "created_at"            TIMESTAMP    NOT NULL,
    "updated_at"            TIMESTAMP    NOT NULL,
    PRIMARY KEY ("user_id"),
    CONSTRAINT "fk__user_authentications__user_id" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE
);
CREATE UNIQUE INDEX "uq__user_authentications__baas_user_id" ON "user_authentications" ("baas_user_id");

-----------------------------------
-- その他
-----------------------------------

CREATE TABLE "global_kvs_entries"
(
    "key"        VARCHAR(255) NOT NULL,
    "value"      BYTEA        NOT NULL,
    "created_at" TIMESTAMP    NOT NULL,
    "updated_at" TIMESTAMP    NOT NULL,
    PRIMARY KEY ("key")
);

CREATE TABLE "echos"
(
    "id"         UUID         NOT NULL,
    "message"    VARCHAR(255) NOT NULL,
    "timestamp"  TIMESTAMP    NOT NULL,
    "created_at" TIMESTAMP    NOT NULL,
    "updated_at" TIMESTAMP    NOT NULL,
    PRIMARY KEY ("id")
);
