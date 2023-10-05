CREATE DATABASE todos;

CREATE TABLE tasks(
    ID SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(500) NOT NULL,
    completed BOOLEAN NOT NULL DEFAULT false
);

INSERT INTO tasks (
    name,
    completed
) VALUES
('replace door hinges', false),
('wash the car', true);