CREATE TABLE public.periods (
    id bigserial NOT NULL,
    name text NULL,
    days_of_week text NULL,
    days_of_month text NULL,
    months_of_year text NULL,
    CONSTRAINT periods_pkey PRIMARY KEY (id)
)