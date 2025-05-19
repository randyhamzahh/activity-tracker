-- public.users definition

-- Drop table

-- DROP TABLE public.users;

CREATE TABLE public.users (
	id bigserial NOT NULL,
	"name" text NULL,
	email text NULL,
	phone text NULL,
	"password" text NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	user_jid varchar(255) NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id)
);

