
CREATE DATABASE recordings;

DROP TABLE IF EXISTS public.artists;
CREATE TABLE public.artists 
(
    id            VARCHAR(50)     PRIMARY KEY,
    name          VARCHAR(128)    NOT NULL,
    country_code  VARCHAR(4)      NOT NULL,
    profile_image VARCHAR(255),
    updated_at    TIMESTAMP       NOT NULL,
    created_at    TIMESTAMP       NOT NULL
);

DROP TABLE IF EXISTS public.albums;
CREATE TABLE public.albums 
(
    id          VARCHAR(50),
    title       VARCHAR(128)    NOT NULL,
    artist_id   VARCHAR(50)     NOT NULL,
    price       NUMERIC(5,2)    NOT NULL,
    album_art   VARCHAR(255),
    updated_at  TIMESTAMP       NOT NULL,
    created_at  TIMESTAMP       NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT fk_artist FOREIGN KEY (artist_id) REFERENCES artists (id)
);
