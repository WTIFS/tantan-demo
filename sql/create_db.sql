--
-- PostgreSQL database dump
--

-- Dumped from database version 10.3
-- Dumped by pg_dump version 10.3

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


--
-- Name: relationship; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.relationship AS ENUM (
    'default',
    'liked',
    'disliked',
    'blocked'
);


ALTER TYPE public.relationship OWNER TO postgres;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: relationships; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.relationships (
    from_user_id bigint NOT NULL,
    to_user_id bigint NOT NULL,
    add_time timestamp(0) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    state public.relationship NOT NULL,
    id bigint NOT NULL
);


ALTER TABLE public.relationships OWNER TO postgres;

--
-- Name: relationships_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.relationships_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.relationships_id_seq OWNER TO postgres;

--
-- Name: relationships_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.relationships_id_seq OWNED BY public.relationships.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    name character varying(200)
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: COLUMN users.name; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.users.name IS '姓名';


--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: relationships id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.relationships ALTER COLUMN id SET DEFAULT nextval('public.relationships_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);



--
-- Name: relationships_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.relationships_id_seq', 30, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 20, true);


--
-- Name: relationships relationships_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.relationships
    ADD CONSTRAINT relationships_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: idx_user_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_user_id ON public.relationships USING btree (from_user_id, to_user_id);


--
-- Name: TABLE users; Type: ACL; Schema: public; Owner: postgres
--

GRANT ALL ON TABLE public.users TO chenyuanfei;


--
-- PostgreSQL database dump complete
--

