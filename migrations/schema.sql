--
-- PostgreSQL database dump
--

-- Dumped from database version 14.18 (Homebrew)
-- Dumped by pg_dump version 14.18 (Homebrew)

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: productinformation; Type: TABLE; Schema: public; Owner: jamiuafolabi
--

CREATE TABLE public.productinformation (
    id character varying(255) NOT NULL,
    name character varying(255) NOT NULL,
    category character varying(255) NOT NULL,
    price numeric NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.productinformation OWNER TO jamiuafolabi;

--
-- Name: salestransaction; Type: TABLE; Schema: public; Owner: jamiuafolabi
--

CREATE TABLE public.salestransaction (
    id character varying(255) NOT NULL,
    product_id character varying(255) NOT NULL,
    quantity integer NOT NULL,
    date timestamp without time zone NOT NULL,
    customer_info character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.salestransaction OWNER TO jamiuafolabi;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: jamiuafolabi
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO jamiuafolabi;

--
-- Name: productinformation productinformation_pkey; Type: CONSTRAINT; Schema: public; Owner: jamiuafolabi
--

ALTER TABLE ONLY public.productinformation
    ADD CONSTRAINT productinformation_pkey PRIMARY KEY (id);


--
-- Name: salestransaction salestransaction_pkey; Type: CONSTRAINT; Schema: public; Owner: jamiuafolabi
--

ALTER TABLE ONLY public.salestransaction
    ADD CONSTRAINT salestransaction_pkey PRIMARY KEY (id);


--
-- Name: schema_migration schema_migration_pkey; Type: CONSTRAINT; Schema: public; Owner: jamiuafolabi
--

ALTER TABLE ONLY public.schema_migration
    ADD CONSTRAINT schema_migration_pkey PRIMARY KEY (version);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: jamiuafolabi
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- Name: salestransaction salestransaction_productinformation_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: jamiuafolabi
--

ALTER TABLE ONLY public.salestransaction
    ADD CONSTRAINT salestransaction_productinformation_id_fk FOREIGN KEY (product_id) REFERENCES public.productinformation(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

