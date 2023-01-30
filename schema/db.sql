
CREATE DATABASE recordings;

DROP TABLE IF EXISTS public.albums;
CREATE TABLE public.albums 
(
    id          VARCHAR(50)     PRIMARY KEY,
    title       VARCHAR(128)    NOT NULL,
    artist      VARCHAR(255)    NOT NULL,
    price       NUMERIC(5,2)    NOT NULL
);

INSERT INTO albums
  (id, title, artist, price)
VALUES
  ('1', 'Blue Train', 'John Coltrane', 56.99),
  ('2', 'Giant Steps', 'John Coltrane', 63.99),
  ('3', 'Jeru', 'Gerry Mulligan', 17.99),
  ('4', 'Sarah Vaughan', 'Sarah Vaughan', 34.98);