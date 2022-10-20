-- Table: public.TYPE_TRANSACION

-- DROP TABLE IF EXISTS public."TYPE_TRANSACION";

CREATE TABLE IF NOT EXISTS public."TYPE_TRANSACTION"
(
    type integer NOT NULL,
    description text COLLATE pg_catalog."default",
    nature text COLLATE pg_catalog."default",
    CONSTRAINT "TYPE_TRANSACION_pkey" PRIMARY KEY (type)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public."TYPE_TRANSACTION"
    OWNER to root;


-- Table: public.PRODUCER

-- DROP TABLE IF EXISTS public."PRODUCER";

CREATE TABLE IF NOT EXISTS public."PRODUCER"
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    NAME text COLLATE pg_catalog."default",
    CONSTRAINT "PRODUCER_pkey" PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public."PRODUCER"
    OWNER to root;


-- Table: public.TRANSACTION

-- DROP TABLE IF EXISTS public."TRANSACTION";

CREATE TABLE IF NOT EXISTS public."TRANSACTION"
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    type_id integer,
    product text COLLATE pg_catalog."default",
    value numeric,
    producer_id integer NOT NULL,
    date date,
    created_at date,
    updated_at date,
    deleted_at date,
    CONSTRAINT "TRANSACTION_pkey" PRIMARY KEY (id),
    CONSTRAINT fk_type FOREIGN KEY (type_id)
        REFERENCES public."TYPE_TRANSACTION" (type) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT fk_producer FOREIGN KEY (producer_id)
        REFERENCES public."PRODUCER" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public."TRANSACTION"
    OWNER to root;

INSERT INTO public."TYPE_TRANSACTION"(type, description, nature)	VALUES (1, 'Venda produtor', 'Entrada');
INSERT INTO public."TYPE_TRANSACTION"(type, description, nature)	VALUES (2, 'Venda afiliado', 'Entrada');
INSERT INTO public."TYPE_TRANSACTION"(type, description, nature)	VALUES (3, 'Comissão paga', 'Saída');
INSERT INTO public."TYPE_TRANSACTION"(type, description, nature)	VALUES (4, 'Comissão recebida', 'Entrada');
