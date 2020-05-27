create table todo (
	id bigserial primary key,
	title character varying(100) NOT NULL,
	description text,
	favorite boolean not null default false
);

CREATE TABLE todo_tag(
   todo_id bigint NOT NULL,
   tag character varying(100) NOT NULL,
   CONSTRAINT fk_todo FOREIGN KEY (todo_id) REFERENCES todo (id) ON UPDATE NO ACTION ON DELETE NO ACTION,
   CONSTRAINT todo_tag_pk PRIMARY KEY (todo_id, tag)
);

ALTER TABLE todo
  ADD COLUMN completed boolean NOT NULL DEFAULT false;