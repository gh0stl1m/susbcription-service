CREATE TABLE "subscriptions"."plans" (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "plan_name" character varying(255),
  "plan_amount" integer,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id")
);

CREATE TABLE "subscriptions"."users" (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "email" character varying(255),
  "first_name" character varying(255),
  "last_name" character varying(255),
  "password" character varying(60),
  "user_active" integer DEFAULT 0,
  "is_admin" integer default 0,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id")
);

CREATE TABLE "subscriptions"."user_plans" (
  "user_id" uuid REFERENCES "subscriptions".users("id"),
  "plan_id" uuid REFERENCES "subscriptions".plans("id"),
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT user_plans_pk
    PRIMARY KEY ("user_id", "plan_id")
);

