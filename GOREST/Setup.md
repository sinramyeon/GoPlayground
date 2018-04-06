
1. [Glide]

package management Tool like NPM

[https://glide.readthedocs.io/en/latest/getting-started/](https://glide.readthedocs.io/en/latest/getting-started/)

#### Install
1) mac
brew install glide
2) linux
curl https://glide.sh/get | sh
3) windows
https://github.com/Masterminds/glide/releases
Download : glide-v0.12.3-windows-amd64.zip
extract glide.exe file in zip, and copy it into your
GOPATH/bin. (Please check your GOPATH!)
#### Setup

1) clone project
2) `glide install`


2. [Docker]

- `docker-compose up -d`

3. [TABLE SCRIPT]

```
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
    VALUES ('New Recipe1', '1 min', 1, true), ('New Recipe2', '2 min', 2, true), ('New Recipe3', '3 min', 3, true),('New Recipe4', '1 min', 1, false),('New Recipe5', '2 min', 2, false),('New Recipe6', '3 min', 3, false),('New Recipe7', '1 min', 1, true), ('New Recipe8', '2 min', 2, true), ('New Recipe9', '3 min', 3, true),('New Recipe10', '1 min', 1, false),('New Recipe11', '2 min', 2, false),('New Recipe12', '3 min', 3, false),('New Recipe13', '1 min', 1, true), ('New Recipe14', '2 min', 2, true), ('New Recipe15', '3 min', 3, true),('New Recipe16', '1 min', 1, false),('New Recipe17', '2 min', 2, false),('New Recipe18', '3 min', 3, false),('New Recipe19', '1 min', 1, true), ('New Recipe20', '2 min', 2, true), ('New Recipe21', '3 min', 3, true),('New Recipe22', '1 min', 1, false),('New Recipe23', '2 min', 2, false),('New Recipe24', '3 min', 3, false),('New Recipe25', '1 min', 1, true);

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
```

4. [API EndPoint]

##### Recipes

| Name | Method | URL | Protected |
| --- | --- | --- | --- |
| List | `GET` | `/recipes` | ✘ |
| List | `GET` | `/recipes?p={p}` | ✘ |
| Create | `POST` | `/recipes` | ✓ |
| Get | `GET` | `/recipes/{id}` | ✘ |
| Update | `PUT/PATCH` | `/recipes/{id}` | ✓ |
| Delete | `DELETE` | `/recipes/{id}` | ✓ |
| Rate | `Post` | `/recipes/{id}/rating` | ✘ |
| Search | `GET` | `/recipes/name/search?q={q}` | ✘ |
| Search | `GET` | `/recipes/time/search?q={q}` | ✘ |
| Search | `GET` | `/recipes/difficulty/search?q={q}` | ✘ |
| Search | `GET` | `/recipes/vegeterain/search?q={q}` | ✘ |
| Search | `GET` | `/recipes/name/search?q={q}&p={p}` | ✘ |
| Search | `GET` | `/recipes/time/search?q={q}&p={p}` | ✘ |
| Search | `GET` | `/recipes/difficulty/search?q={q}&p={p}` | ✘ |
| Search | `GET` | `/recipes/vegeterain/search?q={q}&p={p}` | ✘ |

5. [Security]

- basic authorisation with username / password (hellofresh / hellofresh)
- jwt token

5. [Fresh]

#### Command line tool that builds and (re)starts your web application everytime you save a Go or template file
[https://github.com/pilu/fresh](https://github.com/pilu/fresh)

- `go get github.com/pilu/fresh`
- type `fresh` to start server

6. [Document]

- `godoc -http=":6060"`
- [http://localhost:6060/pkg/hero0926-api-test/](http://localhost:6060/pkg/hero0926-api-test/)


---

