-- public.user_activities definition

-- Drop table

-- DROP TABLE public.user_activities;

CREATE TABLE public.user_activities (
	activity_id int8 NOT NULL,
	user_id int8 NOT NULL,
	group_jid varchar(255) NULL,
	group_id int4 NULL,
	user_jid varchar(255) NOT NULL,
	CONSTRAINT fk_activity_id FOREIGN KEY (activity_id) REFERENCES public.activities(id),
	CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES public.users(id)
);
