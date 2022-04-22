CREATE TABLE "products" (
  "id" SERIAL PRIMARY KEY,
  "title" varchar(255) NOT NULL,
  "price" int NOT NULL,
  "description" varchar(10000) NOT NULL,
  "image_0" varchar(255) NOT NULL,
  "image_1" varchar(255) NOT NULL,
  "image_2" varchar(255) NOT NULL,
  "image_3" varchar(255) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  "valid_until" timestamp NOT NULL
);

CREATE TABLE "orders" (
  "id" SERIAL PRIMARY KEY,
  "user_id" int NOT NULL,
  "product_id" int NOT NULL,
  "quantity" int NOT NULL,
  "total_price" int NOT NULL,
  "to_location" VARCHAR(1000) NOT NULL,
  "is_draft" boolean NOT NULL DEFAULT true,
  "is_payed" boolean NOT NULL DEFAULT false,
  "is_shipped" boolean NOT NULL DEFAULT false,
  "current_location" VARCHAR(1000) NOT NULL DEFAULT ''
);

CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "username" int UNIQUE NOT NULL,
  "encoded_hash" char(255) NOT NULL,
  "phone_number" varchar(10) NOT NULL,
  "email" varchar(255) NOT NULL,
  "address" varchar(500) NOT NULL
);

CREATE INDEX ON "products" ("title");

CREATE INDEX ON "users" ("username");

ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");
