--
-- PostgreSQL database dump
--

-- Dumped from database version 16.2 (Debian 16.2-1.pgdg120+2)
-- Dumped by pg_dump version 16.2 (Debian 16.2-1.pgdg120+2)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: pgcrypto; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS pgcrypto WITH SCHEMA public;


--
-- Name: EXTENSION pgcrypto; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION pgcrypto IS 'cryptographic functions';


--
-- Name: age_enum; Type: TYPE; Schema: public; Owner: exampleuser
--

CREATE TYPE public.age_enum AS ENUM (
    '18-25',
    '26-30',
    '31-36',
    '36-40',
    '41-50',
    '51-59',
    '60+'
);


ALTER TYPE public.age_enum OWNER TO exampleuser;

--
-- Name: fragrance_enum; Type: TYPE; Schema: public; Owner: exampleuser
--

CREATE TYPE public.fragrance_enum AS ENUM (
    'Cheesy',
    'Fungus Amongus',
    'Floral'
);


ALTER TYPE public.fragrance_enum OWNER TO exampleuser;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: config; Type: TABLE; Schema: public; Owner: exampleuser
--

CREATE TABLE public.config (
    "ConfigID" integer NOT NULL,
    "DateTime" timestamp with time zone NOT NULL,
    "Seed" character varying(500)
);


ALTER TABLE public.config OWNER TO exampleuser;

--
-- Name: config_ConfigID_seq; Type: SEQUENCE; Schema: public; Owner: exampleuser
--

CREATE SEQUENCE public."config_ConfigID_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public."config_ConfigID_seq" OWNER TO exampleuser;

--
-- Name: config_ConfigID_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: exampleuser
--

ALTER SEQUENCE public."config_ConfigID_seq" OWNED BY public.config."ConfigID";


--
-- Name: feet; Type: TABLE; Schema: public; Owner: exampleuser
--

CREATE TABLE public.feet (
    "FeetID" integer NOT NULL,
    "Size" integer NOT NULL,
    "Fragrance" public.fragrance_enum NOT NULL,
    "Length" integer NOT NULL,
    "Width" integer NOT NULL,
    "Calluses" boolean NOT NULL,
    "Callus Count" integer,
    "Age" public.age_enum NOT NULL,
    "Bunion" boolean NOT NULL,
    "Preferences" boolean NOT NULL,
    "SubmissionDateTime" timestamp with time zone NOT NULL,
    "UserName" character varying(100)
);


ALTER TABLE public.feet OWNER TO exampleuser;

--
-- Name: feet_FeetID_seq; Type: SEQUENCE; Schema: public; Owner: exampleuser
--

CREATE SEQUENCE public."feet_FeetID_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public."feet_FeetID_seq" OWNER TO exampleuser;

--
-- Name: feet_FeetID_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: exampleuser
--

ALTER SEQUENCE public."feet_FeetID_seq" OWNED BY public.feet."FeetID";


--
-- Name: logins; Type: TABLE; Schema: public; Owner: exampleuser
--

CREATE TABLE public.logins (
    "LoginID" integer NOT NULL,
    "LoginDate" timestamp with time zone NOT NULL,
    "LoginSuccess" boolean NOT NULL,
    "UserName" character varying(100)
);


ALTER TABLE public.logins OWNER TO exampleuser;

--
-- Name: logins_LoginID_seq; Type: SEQUENCE; Schema: public; Owner: exampleuser
--

CREATE SEQUENCE public."logins_LoginID_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public."logins_LoginID_seq" OWNER TO exampleuser;

--
-- Name: logins_LoginID_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: exampleuser
--

ALTER SEQUENCE public."logins_LoginID_seq" OWNED BY public.logins."LoginID";


--
-- Name: users; Type: TABLE; Schema: public; Owner: exampleuser
--

CREATE TABLE public.users (
    "UserID" integer NOT NULL,
    "UserName" character varying(100) NOT NULL,
    "PasswordHash" character varying(60) NOT NULL,
    "DateCreated" timestamp with time zone NOT NULL
);


ALTER TABLE public.users OWNER TO exampleuser;

--
-- Name: users_UserID_seq; Type: SEQUENCE; Schema: public; Owner: exampleuser
--

CREATE SEQUENCE public."users_UserID_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public."users_UserID_seq" OWNER TO exampleuser;

--
-- Name: users_UserID_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: exampleuser
--

ALTER SEQUENCE public."users_UserID_seq" OWNED BY public.users."UserID";


--
-- Name: config ConfigID; Type: DEFAULT; Schema: public; Owner: exampleuser
--

ALTER TABLE ONLY public.config ALTER COLUMN "ConfigID" SET DEFAULT nextval('public."config_ConfigID_seq"'::regclass);


--
-- Name: feet FeetID; Type: DEFAULT; Schema: public; Owner: exampleuser
--

ALTER TABLE ONLY public.feet ALTER COLUMN "FeetID" SET DEFAULT nextval('public."feet_FeetID_seq"'::regclass);


--
-- Name: logins LoginID; Type: DEFAULT; Schema: public; Owner: exampleuser
--

ALTER TABLE ONLY public.logins ALTER COLUMN "LoginID" SET DEFAULT nextval('public."logins_LoginID_seq"'::regclass);


--
-- Name: users UserID; Type: DEFAULT; Schema: public; Owner: exampleuser
--

ALTER TABLE ONLY public.users ALTER COLUMN "UserID" SET DEFAULT nextval('public."users_UserID_seq"'::regclass);


--
-- Data for Name: config; Type: TABLE DATA; Schema: public; Owner: exampleuser
--

COPY public.config ("ConfigID", "DateTime", "Seed") FROM stdin;
\.


--
-- Data for Name: feet; Type: TABLE DATA; Schema: public; Owner: exampleuser
--

COPY public.feet ("FeetID", "Size", "Fragrance", "Length", "Width", "Calluses", "Callus Count", "Age", "Bunion", "Preferences", "SubmissionDateTime", "UserName") FROM stdin;
\.


--
-- Data for Name: logins; Type: TABLE DATA; Schema: public; Owner: exampleuser
--

COPY public.logins ("LoginID", "LoginDate", "LoginSuccess", "UserName") FROM stdin;
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: exampleuser
--

COPY public.users ("UserID", "UserName", "PasswordHash", "DateCreated") FROM stdin;
\.


--
-- Name: config_ConfigID_seq; Type: SEQUENCE SET; Schema: public; Owner: exampleuser
--

SELECT pg_catalog.setval('public."config_ConfigID_seq"', 1, false);


--
-- Name: feet_FeetID_seq; Type: SEQUENCE SET; Schema: public; Owner: exampleuser
--

SELECT pg_catalog.setval('public."feet_FeetID_seq"', 1, false);


--
-- Name: logins_LoginID_seq; Type: SEQUENCE SET; Schema: public; Owner: exampleuser
--

SELECT pg_catalog.setval('public."logins_LoginID_seq"', 1, false);


--
-- Name: users_UserID_seq; Type: SEQUENCE SET; Schema: public; Owner: exampleuser
--

SELECT pg_catalog.setval('public."users_UserID_seq"', 1, false);


--
-- Name: config config_pkey; Type: CONSTRAINT; Schema: public; Owner: exampleuser
--

ALTER TABLE ONLY public.config
    ADD CONSTRAINT config_pkey PRIMARY KEY ("ConfigID");


--
-- Name: feet feet_pkey; Type: CONSTRAINT; Schema: public; Owner: exampleuser
--

ALTER TABLE ONLY public.feet
    ADD CONSTRAINT feet_pkey PRIMARY KEY ("FeetID");


--
-- Name: logins logins_pkey; Type: CONSTRAINT; Schema: public; Owner: exampleuser
--

ALTER TABLE ONLY public.logins
    ADD CONSTRAINT logins_pkey PRIMARY KEY ("LoginID");


--
-- Name: users users_UserName_key; Type: CONSTRAINT; Schema: public; Owner: exampleuser
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT "users_UserName_key" UNIQUE ("UserName");


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: exampleuser
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY ("UserID");


--
-- Name: feet feet_UserName_fkey; Type: FK CONSTRAINT; Schema: public; Owner: exampleuser
--

ALTER TABLE ONLY public.feet
    ADD CONSTRAINT "feet_UserName_fkey" FOREIGN KEY ("UserName") REFERENCES public.users("UserName");


--
-- Name: logins logins_UserName_fkey; Type: FK CONSTRAINT; Schema: public; Owner: exampleuser
--

ALTER TABLE ONLY public.logins
    ADD CONSTRAINT "logins_UserName_fkey" FOREIGN KEY ("UserName") REFERENCES public.users("UserName");


--
-- PostgreSQL database dump complete
--

