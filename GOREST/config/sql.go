package config

const InitSQL = `

DROP TABLE IF EXISTS "RECIPE";
DROP TABLE IF EXISTS "RATE";

CREATE TABLE public."RECIPE"
(
"UNIQUEID" serial,
"NAME" character varying(60) COLLATE pg_catalog."default" NOT NULL,
"PREPTIME" character varying(10) COLLATE pg_catalog."default" NOT NULL,
"DIFFICULTY" integer NOT NULL,
"VEGETARIAN" boolean NOT NULL,
CONSTRAINT "RECIPE_pkey" PRIMARY KEY ("UNIQUEID")
)
WITH (
OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public."RECIPE"
OWNER to hellofresh;


INSERT INTO public."RECIPE"(
"NAME", "PREPTIME", "DIFFICULTY", "VEGETARIAN")
VALUES ('New Recipe1', '1 min', 1, true),  ('New Recipe2', '2 min', 2, true), ('New Recipe3', '3 min', 3, true),('New Recipe4', '1 min', 1, false),('New Recipe5', '2 min', 2, false),('New Recipe6', '3 min', 3, false),('New Recipe7', '1 min', 1, true),  ('New Recipe8', '2 min', 2, true), ('New Recipe9', '3 min', 3, true),('New Recipe10', '1 min', 1, false),('New Recipe11', '2 min', 2, false),('New Recipe12', '3 min', 3, false),('New Recipe13', '1 min', 1, true),  ('New Recipe14', '2 min', 2, true), ('New Recipe15', '3 min', 3, true),('New Recipe16', '1 min', 1, false),('New Recipe17', '2 min', 2, false),('New Recipe18', '3 min', 3, false),('New Recipe19', '1 min', 1, true),  ('New Recipe20', '2 min', 2, true), ('New Recipe21', '3 min', 3, true),('New Recipe22', '1 min', 1, false),('New Recipe23', '2 min', 2, false),('New Recipe24', '3 min', 3, false),('New Recipe25', '1 min', 1, true);

CREATE TABLE public."RATE"
(
"INDEX" serial NOT NULL,
"UNIQUEID" integer NOT NULL,
"RATE" integer NOT NULL,
PRIMARY KEY ("INDEX")
)
WITH (
OIDS = FALSE
);

ALTER TABLE public."RATE"
OWNER to hellofresh;
`
