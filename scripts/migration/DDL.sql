CREATE TABLE song_quotes(
    id bigserial,
    quote_text text,
    song_title varchar,
    album_title varchar,
    album_year int,
    band_name varchar,
    constraint pk_song_quotes_id primary key (id)
);