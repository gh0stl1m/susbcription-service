
INSERT INTO "subscriptions"."users"("email","first_name","last_name","password","user_active", "is_admin")
VALUES
    (E'admin@example.com',E'Admin',E'User',E'$2a$12$1zGLuYDDNvATh4RA4avbKuheAMpb1svexSzrQm7up.bnpwQHs0jNe',1,1);


INSERT INTO "subscriptions"."plans"("plan_name","plan_amount")
VALUES
    (E'Bronze Plan',1000),
    (E'Silver Plan',2000),
    (E'Gold Plan',3000);
