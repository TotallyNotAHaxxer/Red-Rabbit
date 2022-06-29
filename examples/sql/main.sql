CREATE TABLE Numbers
(
    uid serial NOT NULL,
    NumberPh character varying(100) NOT NULL,
    Username character varying(500) NOT NULL,
    City character varying(500) NOT NULL,
    StatePh character varying(500) NOT NULL,
    createdWhen date,
    CONSTRAINT Numbers_pkey PRIMARY KEY (uid)
)
WITH (OIDS=FALSE);