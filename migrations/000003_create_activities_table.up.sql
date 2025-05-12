CREATE TABLE public.activities (
	id bigserial NOT NULL,
	name text NULL,
    period_id bigint NOT NULL,
    CONSTRAINT fk_period_id FOREIGN KEY (period_id) REFERENCES public.periods (id),
	CONSTRAINT activities_pkey PRIMARY KEY (id)
);
