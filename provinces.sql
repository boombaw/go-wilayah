--
-- PostgreSQL database dump
--

-- Dumped from database version 9.5.19
-- Dumped by pg_dump version 9.5.19

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: provinces; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.provinces (
    id character(2) DEFAULT NULL::bpchar NOT NULL,
    name character varying(255) DEFAULT NULL::character varying NOT NULL
);


ALTER TABLE public.provinces OWNER TO postgres;

--
-- Data for Name: provinces; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.provinces VALUES ('11', 'ACEH');
INSERT INTO public.provinces VALUES ('12', 'SUMATERA UTARA');
INSERT INTO public.provinces VALUES ('13', 'SUMATERA BARAT');
INSERT INTO public.provinces VALUES ('14', 'RIAU');
INSERT INTO public.provinces VALUES ('15', 'JAMBI');
INSERT INTO public.provinces VALUES ('16', 'SUMATERA SELATAN');
INSERT INTO public.provinces VALUES ('17', 'BENGKULU');
INSERT INTO public.provinces VALUES ('18', 'LAMPUNG');
INSERT INTO public.provinces VALUES ('19', 'KEPULAUAN BANGKA BELITUNG');
INSERT INTO public.provinces VALUES ('21', 'KEPULAUAN RIAU');
INSERT INTO public.provinces VALUES ('31', 'DKI JAKARTA');
INSERT INTO public.provinces VALUES ('32', 'JAWA BARAT');
INSERT INTO public.provinces VALUES ('33', 'JAWA TENGAH');
INSERT INTO public.provinces VALUES ('34', 'DI YOGYAKARTA');
INSERT INTO public.provinces VALUES ('35', 'JAWA TIMUR');
INSERT INTO public.provinces VALUES ('36', 'BANTEN');
INSERT INTO public.provinces VALUES ('51', 'BALI');
INSERT INTO public.provinces VALUES ('52', 'NUSA TENGGARA BARAT');
INSERT INTO public.provinces VALUES ('53', 'NUSA TENGGARA TIMUR');
INSERT INTO public.provinces VALUES ('61', 'KALIMANTAN BARAT');
INSERT INTO public.provinces VALUES ('62', 'KALIMANTAN TENGAH');
INSERT INTO public.provinces VALUES ('63', 'KALIMANTAN SELATAN');
INSERT INTO public.provinces VALUES ('64', 'KALIMANTAN TIMUR');
INSERT INTO public.provinces VALUES ('65', 'KALIMANTAN UTARA');
INSERT INTO public.provinces VALUES ('71', 'SULAWESI UTARA');
INSERT INTO public.provinces VALUES ('72', 'SULAWESI TENGAH');
INSERT INTO public.provinces VALUES ('73', 'SULAWESI SELATAN');
INSERT INTO public.provinces VALUES ('74', 'SULAWESI TENGGARA');
INSERT INTO public.provinces VALUES ('75', 'GORONTALO');
INSERT INTO public.provinces VALUES ('76', 'SULAWESI BARAT');
INSERT INTO public.provinces VALUES ('81', 'MALUKU');
INSERT INTO public.provinces VALUES ('82', 'MALUKU UTARA');
INSERT INTO public.provinces VALUES ('91', 'PAPUA BARAT');
INSERT INTO public.provinces VALUES ('94', 'PAPUA');


--
-- Name: provinces_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.provinces
    ADD CONSTRAINT provinces_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

