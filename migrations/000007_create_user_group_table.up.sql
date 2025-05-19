CREATE TABLE public.user_groups (
    user_id integer NOT NULL,
    group_id integer NOT NULL,
    user_jid character varying(255) NOT NULL,
    group_jid character varying(255) NOT NULL,
    CONSTRAINT user_group_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id),
    CONSTRAINT user_group_group_id_fkey FOREIGN KEY (group_id) REFERENCES public.groups(id)
)