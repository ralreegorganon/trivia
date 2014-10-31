
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table trivia
(
      trivia_id serial not null,
      question character varying,
      answer character varying,
      constraint trivia_pkey primary key (trivia_id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table trivia;
