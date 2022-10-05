CREATE SEQUENCE IF NOT EXISTS public.seq_planet
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

ALTER SEQUENCE public.seq_planet
    OWNER TO root;


CREATE TABLE IF NOT EXISTS public.planets
(
    id integer NOT NULL DEFAULT nextval('seq_planet'::regclass),
    nome text COLLATE pg_catalog."default" NOT NULL,
    terreno text COLLATE pg_catalog."default" NOT NULL,
    clima text COLLATE pg_catalog."default" NOT NULL,
    aparicoes integer DEFAULT 0,
    CONSTRAINT pk_planeta PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.planets
    OWNER to root;


INSERT INTO public.planets(nome, terreno, clima, aparicoes)
	VALUES ('tathoine', 'desertico', 'arido', 2);
INSERT INTO public.planets(nome, terreno, clima, aparicoes)
	VALUES ('hoth', 'tundra', 'congelado', 1);
INSERT INTO public.planets(nome, terreno, clima, aparicoes)
	VALUES ('coruscant', 'urbano', 'temperado', 5);