-- public.routines definition

-- Drop table

-- DROP TABLE public.routines;

CREATE TABLE public.routines (
	id bigserial NOT NULL,
	"date" date NULL,
	user_id int8 NOT NULL,
	activity_id int8 NOT NULL,
	period_id int8 NOT NULL,
	status text NULL,
	CONSTRAINT routines_pkey PRIMARY KEY (id),
    CONSTRAINT fk_activity_id FOREIGN KEY (activity_id) REFERENCES public.activities(id),
    CONSTRAINT fk_period_id FOREIGN KEY (period_id) REFERENCES public.periods(id),
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES public.users(id)
);