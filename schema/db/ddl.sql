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

CREATE TABLE "user_profiles"
(
    "user_id"    UUID      NOT NULL,
    "content"    BYTEA     NOT NULL,
    "created_at" TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP NOT NULL,
    PRIMARY KEY ("user_id"),
    CONSTRAINT "fk__user_profiles__user_id" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE
);

CREATE TABLE "private_kvs_entries"
(
    "user_id"    UUID         NOT NULL,
    "key"        VARCHAR(255) NOT NULL,
    "value"      BYTEA        NOT NULL,
    "created_at" TIMESTAMP    NOT NULL,
    "updated_at" TIMESTAMP    NOT NULL,
    PRIMARY KEY ("user_id", "key"),
    CONSTRAINT "fk__private_kvs_entries__user_id" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE
);

CREATE TABLE "private_kvs_etags"
(
    "user_id"    UUID      NOT NULL,
    "etag"       UUID      NOT NULL,
    "created_at" TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP NOT NULL,
    PRIMARY KEY ("user_id"),
    CONSTRAINT "fk__private_kvs_etags__user_id" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE
);

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

CREATE TABLE "leader_board"
(
    "id"         VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP    NOT NULL,
    "updated_at" TIMESTAMP    NOT NULL,
    PRIMARY KEY ("id")
);

CREATE TABLE "leader_board_scores"
(
    "leader_board_id" VARCHAR(255) NOT NULL,
    "score_id"        VARCHAR(255) NOT NULL,
    "score"           INTEGER      NOT NULL,
    "timestamp"       TIMESTAMP    NOT NULL,
    "created_at"      TIMESTAMP    NOT NULL,
    "updated_at"      TIMESTAMP    NOT NULL,
    PRIMARY KEY ("leader_board_id", "score_id"),
    CONSTRAINT "fk__leader_board_scores__leader_board_id" FOREIGN KEY ("leader_board_id") REFERENCES "leader_board" ("id") ON DELETE CASCADE
);

CREATE TABLE "rooms"
(
    "id"            UUID         NOT NULL,
    "owner_user_id" UUID         NOT NULL,
    "type"          INTEGER      NOT NULL,
    "max_capacity"  INTEGER      NOT NULL,
    "secret"        VARCHAR(255) NOT NULL,
    "details"       BYTEA        NOT NULL,
    "created_at"    TIMESTAMP    NOT NULL,
    "updated_at"    TIMESTAMP    NOT NULL,
    PRIMARY KEY ("id"),
    CONSTRAINT "fk__rooms__owner_user_id" FOREIGN KEY ("owner_user_id") REFERENCES "users" ("id") ON DELETE CASCADE
);

CREATE TABLE "room_users"
(
    "room_id"    UUID      NOT NULL,
    "user_id"    UUID      NOT NULL,
    "created_at" TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP NOT NULL,
    PRIMARY KEY ("room_id", "user_id"),
    CONSTRAINT "fk__room_users__room_id" FOREIGN KEY ("room_id") REFERENCES "rooms" ("id") ON DELETE CASCADE,
    CONSTRAINT "fk__room_users__user_id" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE
);

CREATE TABLE "master_data"
(
    "revision"   INTEGER      NOT NULL,
    "content"    BYTEA        NOT NULL,
    "is_active"  BOOLEAN      NOT NULL,
    "comment"    VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP    NOT NULL,
    "updated_at" TIMESTAMP    NOT NULL,
    PRIMARY KEY ("revision")
);
CREATE UNIQUE INDEX "uq__master_data__is_active" ON master_data ("is_active") WHERE "is_active" = TRUE;

CREATE TABLE "echos"
(
    "id"         UUID         NOT NULL,
    "message"    VARCHAR(255) NOT NULL,
    "timestamp"  TIMESTAMP    NOT NULL,
    "created_at" TIMESTAMP    NOT NULL,
    "updated_at" TIMESTAMP    NOT NULL,
    PRIMARY KEY ("id")
);
