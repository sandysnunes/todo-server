create table todo (
	id bigserial primary key,
	title character varying(100),
	description text,
	favorite boolean
);