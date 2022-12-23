-- Write your migrate up statements here

CREATE TABLE "roles" (
     "id" bigserial,
     "created_at" timestamptz,
     "updated_at" timestamptz,
     "name" varchar(128) DEFAULT NULL,
     "slug" varchar(32) UNIQUE NOT NULL,
     "is_permitted" boolean DEFAULT true,
     PRIMARY KEY ("id")
);
CREATE INDEX roles_name ON "roles" ("name");
CREATE INDEX roles_is_permitted ON "roles" ("is_permitted");

CREATE TABLE "users" (
     "id" bigserial,
     "created_at" timestamptz,
     "updated_at" timestamptz,
     "email" varchar(255) UNIQUE NOT NULL,
     "password" varchar(255) NOT NULL,
     "first_name" varchar(128) NOT NULL,
     "last_name" varchar(128) NOT NULL,
     "phone" varchar(32) DEFAULT NULL,
     "is_active" boolean DEFAULT false,
     "is_2fa" boolean DEFAULT false,
     "token_2fa" varchar(64) DEFAULT NULL,
     "last_login_at" timestamptz DEFAULT NULL,
     "role_id" bigint,
     PRIMARY KEY ("id")
);
CREATE INDEX users_first_name ON "users" ("first_name");
CREATE INDEX users_last_name ON "users" ("last_name");
CREATE INDEX users_phone ON "users" ("phone");
CREATE INDEX users_is_2fa ON "users" ("is_2fa");
CREATE INDEX users_is_active ON "users" ("is_active");
CREATE INDEX users_last_login_at ON "users" ("last_login_at");
ALTER TABLE "users" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");

CREATE TABLE "actions" (
    "id" bigint,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "name" varchar(128),
    "method" varchar(16),
    PRIMARY KEY ("id")
);
CREATE INDEX actions_name ON "actions" ("name");
CREATE INDEX actions_method ON "actions" ("method");

CREATE TABLE "resources" (
     "id" bigint,
     "created_at" timestamptz,
     "updated_at" timestamptz,
     "name" varchar(128),
     "uri_mask" varchar(512),
     PRIMARY KEY ("id")
);
CREATE INDEX resources_name ON "resources" ("name");

CREATE TABLE "permissions" (
    "id" bigint,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "role_id" bigint,
    "action_id" bigint,
    "resource_id" bigint,
    PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX role_action_resource_unique ON "permissions" ("role_id", "action_id", "resource_id");
ALTER TABLE "permissions" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");
ALTER TABLE "permissions" ADD FOREIGN KEY ("action_id") REFERENCES "actions" ("id");
ALTER TABLE "permissions" ADD FOREIGN KEY ("resource_id") REFERENCES "resources" ("id");

CREATE TABLE "journal" (
    "id" bigint,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "user_id" bigint,
    "action_id" bigint,
    "resource_id" bigint,
    "request_data" text,
    "response_data" text,
    PRIMARY KEY ("id")
);
ALTER TABLE "journal" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "journal" ADD FOREIGN KEY ("action_id") REFERENCES "actions" ("id");
ALTER TABLE "journal" ADD FOREIGN KEY ("resource_id") REFERENCES "resources" ("id");

---- create above / drop below ----

DROP TABLE "roles";
DROP TABLE "users";
DROP TABLE "actions";
DROP TABLE "resources";
DROP TABLE "permissions";
DROP TABLE "journal";

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
