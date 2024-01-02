alter table "public"."users" add column "cocktail_id" bigint;

alter table "public"."users" add constraint "users_cocktail_id_check" CHECK ((cocktail_id >= 0)) not valid;

alter table "public"."users" validate constraint "users_cocktail_id_check";

alter table "public"."users" alter column "cocktail_id" set not null;
