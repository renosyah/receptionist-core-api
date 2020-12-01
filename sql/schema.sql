CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

DROP TABLE IF EXISTS "customer" CASCADE;
DROP TABLE IF EXISTS "owner" CASCADE;
DROP TABLE IF EXISTS "store" CASCADE;
DROP TABLE IF EXISTS "seats" CASCADE;
DROP TABLE IF EXISTS "product" CASCADE;
DROP TABLE IF EXISTS "booking" CASCADE;
DROP TABLE IF EXISTS "booking_detail" CASCADE;
DROP TABLE IF EXISTS "transaction" CASCADE;

CREATE TABLE "customer" (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    name TEXT NOT NULL DEFAULT '',
    phone_number TEXT NOT NULL DEFAULT '',
    password TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()::TIMESTAMPTZ,
    flag_status INT NOT NULL DEFAULT 0,
    PRIMARY KEY (id)
);

CREATE TABLE "owner" (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    name TEXT NOT NULL DEFAULT '',
    email TEXT NOT NULL DEFAULT '',
    password TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()::TIMESTAMPTZ,
    flag_status INT NOT NULL DEFAULT 0,
    PRIMARY KEY (id)
);

CREATE TABLE "store" (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    owner_id UUID NOT NULL REFERENCES owner (id),
    name TEXT NOT NULL DEFAULT '',
    description TEXT NOT NULL DEFAULT '',
    image_url TEXT NOT NULL DEFAULT '',
    latitude FLOAT NOT NULL DEFAULT 0.0,
    longitude FLOAT NOT NULL DEFAULT 0.0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()::TIMESTAMPTZ,
    flag_status INT NOT NULL DEFAULT 0,
    PRIMARY KEY (id)
);

CREATE TABLE "seats" (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    store_id UUID NOT NULL REFERENCES store (id),
    name TEXT NOT NULL DEFAULT '',
    description TEXT NOT NULL DEFAULT '',
    position INT NOT NULL DEFAULT 0,
    price DECIMAL NOT NULL DEFAULT '0', 
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()::TIMESTAMPTZ,
    flag_status INT NOT NULL DEFAULT 0,
    PRIMARY KEY (id)
);

CREATE TABLE "product" (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    store_id UUID NOT NULL REFERENCES store (id),
    name TEXT NOT NULL DEFAULT '',
    description TEXT NOT NULL DEFAULT '', 
    image_url TEXT NOT NULL DEFAULT '',
    price DECIMAL NOT NULL DEFAULT '0',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()::TIMESTAMPTZ,
    flag_status INT NOT NULL DEFAULT 0,
    PRIMARY KEY (id)
);

CREATE TABLE "booking" (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    customer_id UUID NOT NULL REFERENCES customer (id),
    seats_id UUID NOT NULL REFERENCES seats (id), 
    price DECIMAL NOT NULL DEFAULT '0',
    total DECIMAL NOT NULL DEFAULT '0',
    duration_from TIMESTAMPTZ NULL,
    duration_to TIMESTAMPTZ NULL,
    payment_status INT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()::TIMESTAMPTZ,
    flag_status INT NOT NULL DEFAULT 0,
    PRIMARY KEY (id)
);

CREATE TABLE "booking_detail" (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    booking_id UUID NOT NULL REFERENCES booking (id),
    product_id UUID NOT NULL REFERENCES product (id), 
    price DECIMAL NOT NULL DEFAULT '0',
    quantity INT NOT NULL DEFAULT 0,
    sub_total DECIMAL NOT NULL DEFAULT '0',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()::TIMESTAMPTZ,
    flag_status INT NOT NULL DEFAULT 0,
    PRIMARY KEY (id)
);

CREATE TABLE "transaction" (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    booking_id UUID NOT NULL REFERENCES booking (id),
    customer_id UUID NOT NULL REFERENCES customer (id),
    total DECIMAL NOT NULL DEFAULT '0',
    payment_type INT NOT NULL DEFAULT 0, 
    payment_status INT NOT NULL DEFAULT 0,
    payment_order_id TEXT NOT NULL DEFAULT '',
    payment_id TEXT NOT NULL DEFAULT '',
    payment_time TEXT NOT NULL DEFAULT '', 
    approval_code TEXT NOT NULL DEFAULT '',
    bank_name TEXT NOT NULL DEFAULT '',
    va TEXT NOT NULL DEFAULT '',
    cstore_code TEXT NOT NULL DEFAULT '',
    cstore_name TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()::TIMESTAMPTZ,
    flag_status INT NOT NULL DEFAULT 0,
    PRIMARY KEY (id) 
);
