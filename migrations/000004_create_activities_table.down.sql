-- public.activities definition

-- Drop table

DROP TABLE IF EXISTS public.activities;

-- CREATE TABLE public.activities (
-- 	id bigserial NOT NULL,
-- 	"name" text NULL,
-- 	period_id int8 NOT NULL,
-- 	CONSTRAINT activities_pkey null,
-- 	CONSTRAINT fk_period_id FOREIGN KEY (period_id) REFERENCES public.periods(id)
-- );
