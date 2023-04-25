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

INSERT INTO "roles" (created_at, updated_at, name, slug, is_permitted)
VALUES (now(), now(), 'Super Admin', 'superadmin', False);

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
     "is_disconnected" boolean DEFAULT false,
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

INSERT INTO "users" (created_at, updated_at, first_name, last_name, email, password,
                     phone, last_login_at, is_active, is_disconnected, is_2fa, role_id)
VALUES (now(), now(), 'Global', 'Admin', 'admin@usersmanager.io',
        '$2a$10$.gtOT8enT0tU4J4t37XYcOTV26p7/5u.QXMkKBk/851bZi2ddcOWi', NULL, NULL, TRUE, FALSE, FALSE, 1);

CREATE TABLE "resources" (
     "id" bigserial,
     "created_at" timestamptz,
     "updated_at" timestamptz,
     "name" varchar(128),
     "uri_mask" varchar(512),
     "method_mask" int, -- byte mask of HTTP METHODS
     "is_active" boolean DEFAULT false,
     "res_group" varchar(128), -- group resource by specific name
     PRIMARY KEY ("id")
);
CREATE INDEX resources_name ON "resources" ("name");
CREATE INDEX resources_uri_mask ON "resources" ("uri_mask");
CREATE INDEX resources_group ON "resources" ("res_group");

INSERT INTO "resources" (created_at, updated_at, name, uri_mask, method_mask, is_active, res_group)
VALUES (now(), now(), 'Users', '/users', 12, true, 'Users');
INSERT INTO "resources" (created_at, updated_at, name, uri_mask, method_mask, is_active, res_group)
VALUES (now(), now(), 'User', '/users/\d+', 11, true, 'Users');
INSERT INTO "resources" (created_at, updated_at, name, uri_mask, method_mask, is_active, res_group)
VALUES (now(), now(), 'Users recovery', '/users/\d+/recovery', 4, false, 'Users');
INSERT INTO "resources" (created_at, updated_at, name, uri_mask, method_mask, is_active, res_group)
VALUES (now(), now(), 'Roles', '/roles', 12, true, 'Roles');
INSERT INTO "resources" (created_at, updated_at, name, uri_mask, method_mask, is_active, res_group)
VALUES (now(), now(), 'Role', '/roles/\d+', 11, true, 'Roles');
INSERT INTO "resources" (created_at, updated_at, name, uri_mask, method_mask, is_active, res_group)
VALUES (now(), now(), 'Actions', '/actions', 12, true, 'Actions');
INSERT INTO "resources" (created_at, updated_at, name, uri_mask, method_mask, is_active, res_group)
VALUES (now(), now(), 'Action', '/actions/\d+', 11, true, 'Actions');
INSERT INTO "resources" (created_at, updated_at, name, uri_mask, method_mask, is_active, res_group)
VALUES (now(), now(), 'Resources', '/resources', 12, true, 'Resources');
INSERT INTO "resources" (created_at, updated_at, name, uri_mask, method_mask, is_active, res_group)
VALUES (now(), now(), 'Resource', '/resources/\d+', 11, true, 'Resources');
INSERT INTO "resources" (created_at, updated_at, name, uri_mask, method_mask, is_active, res_group)
VALUES (now(), now(), 'Permissions', '/permissions', 12, true, 'Permissions');
INSERT INTO "resources" (created_at, updated_at, name, uri_mask, method_mask, is_active, res_group)
VALUES (now(), now(), 'Permission', '/permissions/\d+', 11, true, 'Permissions');

CREATE TABLE "permissions" (
    "id" bigserial,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "role_id" bigint,
    "resource_id" bigint,
    PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX role_action_resource_unique ON "permissions" ("role_id", "resource_id");
ALTER TABLE "permissions" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");
ALTER TABLE "permissions" ADD FOREIGN KEY ("resource_id") REFERENCES "resources" ("id");

CREATE TABLE "journal" (
    "id" bigserial,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "user_id" bigint,
    "resource_id" bigint,
    "request_data" text,
    "response_data" text,
    PRIMARY KEY ("id")
);
ALTER TABLE "journal" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "journal" ADD FOREIGN KEY ("resource_id") REFERENCES "resources" ("id");

---- create above / drop below ----

DROP TABLE "roles";
DROP TABLE "users";
DROP TABLE "resources";
DROP TABLE "permissions";
DROP TABLE "journal";

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
