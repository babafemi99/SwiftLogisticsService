CREATE TABLE IF NOT EXISTS user_table (
                        user_id uuid ,
                        firstname varchar,
                        lastname varchar,
                        email varchar PRIMARY KEY UNIQUE,
                        profile_picture varchar,
                        phone varchar,
                        user_status varchar,
                        password varchar,
                        default_location varchar,
                        updated_at varchar,
                        created_at varchar
);

CREATE TABLE IF NOT EXISTS rider_table (
                               "rider_id" uuid NOT NULL,
                               "guarantor_id" uuid NOT NULL,
                               "first_name" varchar NOT NULL,
                               "last_name" varchar NOT NULL,
                               "email" varchar NOT NULL,
                               "phone" varchar NOT NULL,
                               "password" varchar NOT NULL ,
                               "dob" varchar NOT NULL,
                               "gender" varchar,
                               "marital_status" varchar,
                               "education_level" varchar,
                               "residential_address" varchar NOT NULL,
                               "driver_license" varchar NOT NULL,
                               "identity_card" varchar NOT NULL,
                               "verification_status" varchar NOT NULL,
                               "account_status" varchar,
                                "created_at" varchar,
                               PRIMARY KEY ("rider_id", "email", "phone")
);

CREATE TABLE IF NOT EXISTS guarantors_table (
                              "guarantor_id" uuid PRIMARY KEY ,
                              "rider_id" uuid ,
                              "email" varchar,
                              "phone" varchar,
                              "first_name" varchar ,
                              "last_name" varchar ,
                              "guarantor_address" varchar ,
                              "guarantor_identification" varchar
);