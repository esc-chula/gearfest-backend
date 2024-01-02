alter table "public"."checkins" drop constraint "unique_user_location";

drop index if exists "public"."unique_user_location";

alter table "public"."checkins" alter column "location_id" set data type bigint using "location_id"::bigint;

alter table "public"."checkins" add constraint "checkins_location_id_check" CHECK ((location_id >= 0)) not valid;

alter table "public"."checkins" validate constraint "checkins_location_id_check";

ALTER TABLE "public"."checkins"
ADD CONSTRAINT unique_user_location UNIQUE (user_id, location_id);