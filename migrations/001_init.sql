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

CREATE TABLE "actions" (
    "id" bigserial,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "name" varchar(128),
    "method" varchar(16),
    PRIMARY KEY ("id")
);
CREATE INDEX actions_name ON "actions" ("name");
CREATE INDEX actions_method ON "actions" ("method");

INSERT INTO "actions" (created_at, updated_at, name, method)
VALUES (now(), now(), 'add', 'POST');
INSERT INTO "actions" (created_at, updated_at, name, method)
VALUES (now(), now(), 'edit', 'PATCH');
INSERT INTO "actions" (created_at, updated_at, name, method)
VALUES (now(), now(), 'list', 'GET');
INSERT INTO "actions" (created_at, updated_at, name, method)
VALUES (now(), now(), 'view', 'GET');
INSERT INTO "actions" (created_at, updated_at, name, method)
VALUES (now(), now(), 'delete', 'DELETE');

CREATE TABLE "resources" (
     "id" bigserial,
     "created_at" timestamptz,
     "updated_at" timestamptz,
     "name" varchar(128),
     "uri_mask" varchar(512),
     PRIMARY KEY ("id")
);
CREATE INDEX resources_name ON "resources" ("name");

INSERT INTO "resources" (created_at, updated_at, name, uri_mask)
VALUES (now(), now(), 'Users CRUD', '/users');
INSERT INTO "resources" (created_at, updated_at, name, uri_mask)
VALUES (now(), now(), 'Roles CRUD', '/roles');
INSERT INTO "resources" (created_at, updated_at, name, uri_mask)
VALUES (now(), now(), 'Actions CRUD', '/actions');
INSERT INTO "resources" (created_at, updated_at, name, uri_mask)
VALUES (now(), now(), 'Resources CRUD', '/resources');
INSERT INTO "resources" (created_at, updated_at, name, uri_mask)
VALUES (now(), now(), 'Permissions CRUD', '/permissions');

CREATE TABLE "permissions" (
    "id" bigserial,
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

INSERT INTO "permissions" (created_at, updated_at, role_id, action_id, resource_id)
VALUES (now(), now(), 1, 1, 1);
INSERT INTO "permissions" (created_at, updated_at, role_id, action_id, resource_id)
VALUES (now(), now(), 1, 2, 1);
INSERT INTO "permissions" (created_at, updated_at, role_id, action_id, resource_id)
VALUES (now(), now(), 1, 3, 1);
INSERT INTO "permissions" (created_at, updated_at, role_id, action_id, resource_id)
VALUES (now(), now(), 1, 4, 1);
INSERT INTO "permissions" (created_at, updated_at, role_id, action_id, resource_id)
VALUES (now(), now(), 1, 5, 1);

INSERT INTO "permissions" (created_at, updated_at, role_id, action_id, resource_id)
VALUES (now(), now(), 1, 1, 2);
INSERT INTO "permissions" (created_at, updated_at, role_id, action_id, resource_id)
VALUES (now(), now(), 1, 2, 2);
INSERT INTO "permissions" (created_at, updated_at, role_id, action_id, resource_id)
VALUES (now(), now(), 1, 3, 2);
INSERT INTO "permissions" (created_at, updated_at, role_id, action_id, resource_id)
VALUES (now(), now(), 1, 4, 2);
INSERT INTO "permissions" (created_at, updated_at, role_id, action_id, resource_id)
VALUES (now(), now(), 1, 5, 2);

INSERT INTO "permissions" (created_at, updated_at, role_id, action_id, resource_id)
VALUES (now(), now(), 1, 1, 3);
INSERT INTO "permissions" (created_at, updated_at, role_id, action_id, resource_id)
VALUES (now(), now(), 1, 2, 3);
INSERT INTO "permissions" (created_at, updated_at, role_id, action_id, resource_id)
VALUES (now(), now(), 1, 3, 3);
INSERT INTO "permissions" (created_at, updated_at, role_id, action_id, resource_id)
VALUES (now(), now(), 1, 4, 3);
INSERT INTO "permissions" (created_at, updated_at, role_id, action_id, resource_id)
VALUES (now(), now(), 1, 5, 3);

INSERT INTO "permissions" (created_at, updated_at, role_id, action_id, resource_id)
VALUES (now(), now(), 1, 1, 4);
INSERT INTO "permissions" (created_at, updated_at, role_id, action_id, resource_id)
VALUES (now(), now(), 1, 2, 4);
INSERT INTO "permissions" (created_at, updated_at, role_id, action_id, resource_id)
VALUES (now(), now(), 1, 3, 4);
INSERT INTO "permissions" (created_at, updated_at, role_id, action_id, resource_id)
VALUES (now(), now(), 1, 4, 4);
INSERT INTO "permissions" (created_at, updated_at, role_id, action_id, resource_id)
VALUES (now(), now(), 1, 5, 4);

INSERT INTO "permissions" (created_at, updated_at, role_id, action_id, resource_id)
VALUES (now(), now(), 1, 1, 5);
INSERT INTO "permissions" (created_at, updated_at, role_id, action_id, resource_id)
VALUES (now(), now(), 1, 2, 5);
INSERT INTO "permissions" (created_at, updated_at, role_id, action_id, resource_id)
VALUES (now(), now(), 1, 3, 5);
INSERT INTO "permissions" (created_at, updated_at, role_id, action_id, resource_id)
VALUES (now(), now(), 1, 4, 5);
INSERT INTO "permissions" (created_at, updated_at, role_id, action_id, resource_id)
VALUES (now(), now(), 1, 5, 5);

CREATE TABLE "journal" (
    "id" bigserial,
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
