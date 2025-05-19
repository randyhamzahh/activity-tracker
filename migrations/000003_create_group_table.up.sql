-- public."groups" definition

-- Drop table

-- DROP TABLE public."groups";

CREATE TABLE public."groups" (
	id bigserial NOT NULL,
	group_jid varchar(255) NOT NULL,
	"name" varchar(255) NOT NULL,
	CONSTRAINT groups_pkey PRIMARY KEY (id)
);