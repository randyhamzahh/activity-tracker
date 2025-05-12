CREATE TABLE public.user_activities (
	activity_id bigint NOT NULL,
	user_id bigint NOT NULL,
	CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES public.users (id),
	CONSTRAINT fk_activity_id FOREIGN KEY (activity_id) REFERENCES public.activities (id)
);
