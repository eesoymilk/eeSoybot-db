CREATE TYPE AUTOREACT AS ( );

CREATE TABLE
    guilds (
        id BIGINT NOT NULL PRIMARY KEY,
        name TEXT NOT NULL,
        autoresponses
    );