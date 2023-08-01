CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
   "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
   "username" text UNIQUE NOT NULL,
   "email" text UNIQUE NOT NULL,
   "password" text NOT NULL
);

CREATE TABLE IF NOT EXISTS coffees (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "name" text NOT NULL,
    "roast" text NOT NULL,
    "region" text NOT NULL,
    "image" text NOT NULL,
    "price" real NOT NULL,
    "grind_unit" integer NOT NULL,
    "created_at" timestamptz DEFAULT NOW(),
    "updated_at" timestamptz DEFAULT NOW()
);