CREATE TABLE public.routines (
    id bigserial NOT NULL,
    date date NULL,
    user_id bigint NOT NULL,
    activity_id bigint NOT NULL,
    period_id bigint NOT NULL,
    status text NULL,
    CONSTRAINT fk_period_id FOREIGN KEY (period_id) REFERENCES public.periods (id),
    CONSTRAINT fk_activity_id FOREIGN KEY (activity_id) REFERENCES public.activities (id),
    CONSTRAINT routines_pkey PRIMARY KEY (id)
)
