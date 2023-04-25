-- pg_dump -U postgres -s recruitment 
--
-- PostgreSQL database dump
--

-- Dumped from database version 13.4 (Debian 13.4-1.pgdg100+1)
-- Dumped by pg_dump version 13.4 (Debian 13.4-1.pgdg100+1)

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
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


--
-- Name: applications_grade_enum; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.applications_grade_enum AS ENUM (
    '0',
    '1',
    '2',
    '3',
    '4'
);


ALTER TYPE public.applications_grade_enum OWNER TO postgres;

--
-- Name: applications_group_enum; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.applications_group_enum AS ENUM (
    'web',
    'lab',
    'ai',
    'game',
    'android',
    'ios',
    'design',
    'pm'
);


ALTER TYPE public.applications_group_enum OWNER TO postgres;

--
-- Name: applications_rank_enum; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.applications_rank_enum AS ENUM (
    '0',
    '1',
    '2',
    '3',
    '4'
);


ALTER TYPE public.applications_rank_enum OWNER TO postgres;

--
-- Name: applications_step_enum; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.applications_step_enum AS ENUM (
    '0',
    '1',
    '2',
    '3',
    '4',
    '5',
    '6',
    '7'
);


ALTER TYPE public.applications_step_enum OWNER TO postgres;

--
-- Name: candidates_gender_enum; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.candidates_gender_enum AS ENUM (
    '0',
    '1',
    '2'
);


ALTER TYPE public.candidates_gender_enum OWNER TO postgres;

--
-- Name: comments_evaluation_enum; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.comments_evaluation_enum AS ENUM (
    '0',
    '1',
    '2'
);


ALTER TYPE public.comments_evaluation_enum OWNER TO postgres;

--
-- Name: interviews_name_enum; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.interviews_name_enum AS ENUM (
    'web',
    'lab',
    'ai',
    'game',
    'android',
    'ios',
    'design',
    'pm',
    'unique'
);


ALTER TYPE public.interviews_name_enum OWNER TO postgres;

--
-- Name: interviews_period_enum; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.interviews_period_enum AS ENUM (
    '0',
    '1',
    '2'
);


ALTER TYPE public.interviews_period_enum OWNER TO postgres;

--
-- Name: members_gender_enum; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.members_gender_enum AS ENUM (
    '0',
    '1',
    '2'
);


ALTER TYPE public.members_gender_enum OWNER TO postgres;

--
-- Name: members_group_enum; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.members_group_enum AS ENUM (
    'web',
    'lab',
    'ai',
    'game',
    'android',
    'ios',
    'design',
    'pm'
);


ALTER TYPE public.members_group_enum OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: applications; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.applications (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    "createdAt" timestamp with time zone DEFAULT now() NOT NULL,
    "updatedAt" timestamp with time zone DEFAULT now() NOT NULL,
    grade public.applications_grade_enum NOT NULL,
    institute character varying NOT NULL,
    major character varying NOT NULL,
    rank public.applications_rank_enum NOT NULL,
    "group" public.applications_group_enum NOT NULL,
    intro character varying NOT NULL,
    "isQuick" boolean NOT NULL,
    referrer character varying,
    resume character varying,
    abandoned boolean DEFAULT false NOT NULL,
    rejected boolean DEFAULT false NOT NULL,
    step public.applications_step_enum DEFAULT '0'::public.applications_step_enum NOT NULL,
    "candidateId" uuid,
    "recruitmentId" uuid,
    "interviewAllocationsGroup" timestamp with time zone,
    "interviewAllocationsTeam" timestamp with time zone
);


ALTER TABLE public.applications OWNER TO postgres;

--
-- Name: candidates; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.candidates (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    "createdAt" timestamp with time zone DEFAULT now() NOT NULL,
    "updatedAt" timestamp with time zone DEFAULT now() NOT NULL,
    name character varying NOT NULL,
    phone character varying NOT NULL,
    mail character varying,
    gender public.candidates_gender_enum NOT NULL,
    "passwordSalt" character varying NOT NULL,
    "passwordHash" character varying NOT NULL
);


ALTER TABLE public.candidates OWNER TO postgres;

--
-- Name: comments; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.comments (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    "createdAt" timestamp with time zone DEFAULT now() NOT NULL,
    "updatedAt" timestamp with time zone DEFAULT now() NOT NULL,
    content character varying NOT NULL,
    evaluation public.comments_evaluation_enum NOT NULL,
    "applicationId" uuid,
    "memberId" uuid
);


ALTER TABLE public.comments OWNER TO postgres;

--
-- Name: interview_selections; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.interview_selections (
    "applicationsId" uuid NOT NULL,
    "interviewsId" uuid NOT NULL
);


ALTER TABLE public.interview_selections OWNER TO postgres;

--
-- Name: interviews; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.interviews (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    "createdAt" timestamp with time zone DEFAULT now() NOT NULL,
    "updatedAt" timestamp with time zone DEFAULT now() NOT NULL,
    date timestamp with time zone NOT NULL,
    period public.interviews_period_enum NOT NULL,
    name public.interviews_name_enum NOT NULL,
    "slotNumber" integer NOT NULL,
    "recruitmentId" uuid
);


ALTER TABLE public.interviews OWNER TO postgres;

--
-- Name: members; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.members (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    "createdAt" timestamp with time zone DEFAULT now() NOT NULL,
    "updatedAt" timestamp with time zone DEFAULT now() NOT NULL,
    name character varying NOT NULL,
    phone character varying NOT NULL,
    mail character varying,
    gender public.members_gender_enum NOT NULL,
    "weChatID" character varying NOT NULL,
    "joinTime" character varying NOT NULL,
    "isCaptain" boolean DEFAULT false NOT NULL,
    "isAdmin" boolean DEFAULT false NOT NULL,
    "group" public.members_group_enum NOT NULL,
    avatar character varying,
    "passwordSalt" character varying NOT NULL,
    "passwordHash" character varying NOT NULL
);


ALTER TABLE public.members OWNER TO postgres;

--
-- Name: recruitments; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.recruitments (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    "createdAt" timestamp with time zone DEFAULT now() NOT NULL,
    "updatedAt" timestamp with time zone DEFAULT now() NOT NULL,
    name character varying NOT NULL,
    beginning timestamp with time zone NOT NULL,
    deadline timestamp with time zone NOT NULL,
    "end" timestamp with time zone NOT NULL,
    statistics jsonb
);


ALTER TABLE public.recruitments OWNER TO postgres;

--
-- Name: candidates PK_140681296bf033ab1eb95288abb; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.candidates
    ADD CONSTRAINT "PK_140681296bf033ab1eb95288abb" PRIMARY KEY (id);


--
-- Name: members PK_28b53062261b996d9c99fa12404; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.members
    ADD CONSTRAINT "PK_28b53062261b996d9c99fa12404" PRIMARY KEY (id);


--
-- Name: recruitments PK_4e63272ea2bc161c08ba2257e87; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.recruitments
    ADD CONSTRAINT "PK_4e63272ea2bc161c08ba2257e87" PRIMARY KEY (id);


--
-- Name: comments PK_8bf68bc960f2b69e818bdb90dcb; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comments
    ADD CONSTRAINT "PK_8bf68bc960f2b69e818bdb90dcb" PRIMARY KEY (id);


--
-- Name: applications PK_938c0a27255637bde919591888f; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.applications
    ADD CONSTRAINT "PK_938c0a27255637bde919591888f" PRIMARY KEY (id);


--
-- Name: interview_selections PK_bd70f0ab2cfd6fa9792b9f026c5; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.interview_selections
    ADD CONSTRAINT "PK_bd70f0ab2cfd6fa9792b9f026c5" PRIMARY KEY ("applicationsId", "interviewsId");


--
-- Name: interviews PK_fd41af1f96d698fa33c2f070f47; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.interviews
    ADD CONSTRAINT "PK_fd41af1f96d698fa33c2f070f47" PRIMARY KEY (id);


--
-- Name: members UQ_061b00c27959b553b9199892cc5; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.members
    ADD CONSTRAINT "UQ_061b00c27959b553b9199892cc5" UNIQUE (mail);


--
-- Name: members UQ_090c60f7851c98387ce3e1995ef; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.members
    ADD CONSTRAINT "UQ_090c60f7851c98387ce3e1995ef" UNIQUE (phone);


--
-- Name: applications UQ_25f4545ed14dc1aa07fddf666c9; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.applications
    ADD CONSTRAINT "UQ_25f4545ed14dc1aa07fddf666c9" UNIQUE ("candidateId", "recruitmentId");


--
-- Name: interviews UQ_2dce828e693cf75351c7e6c8e40; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.interviews
    ADD CONSTRAINT "UQ_2dce828e693cf75351c7e6c8e40" UNIQUE (date, period, name, "recruitmentId");


--
-- Name: members UQ_4a680e426ebcc7b922d91fbc6ab; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.members
    ADD CONSTRAINT "UQ_4a680e426ebcc7b922d91fbc6ab" UNIQUE ("weChatID");


--
-- Name: candidates UQ_6821529f6e7a5ee9f35f3840ee7; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.candidates
    ADD CONSTRAINT "UQ_6821529f6e7a5ee9f35f3840ee7" UNIQUE ("passwordHash");


--
-- Name: recruitments UQ_71625b46a5a6db8f4ce0f735c09; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.recruitments
    ADD CONSTRAINT "UQ_71625b46a5a6db8f4ce0f735c09" UNIQUE (name);


--
-- Name: candidates UQ_a0efe7a0921ca16f5ad25588b94; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.candidates
    ADD CONSTRAINT "UQ_a0efe7a0921ca16f5ad25588b94" UNIQUE (phone);


--
-- Name: members UQ_a6449b66e3b3395da27e055adfa; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.members
    ADD CONSTRAINT "UQ_a6449b66e3b3395da27e055adfa" UNIQUE ("passwordHash");


--
-- Name: candidates UQ_ddae5e5a9d1b9f21e2bb0676bdb; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.candidates
    ADD CONSTRAINT "UQ_ddae5e5a9d1b9f21e2bb0676bdb" UNIQUE (mail);


--
-- Name: IDX_0af6f67fecc8fc0a246622b59d; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "IDX_0af6f67fecc8fc0a246622b59d" ON public.recruitments USING btree ("updatedAt");


--
-- Name: IDX_1b8dfc096390728dc39f96d652; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "IDX_1b8dfc096390728dc39f96d652" ON public.interview_selections USING btree ("applicationsId");


--
-- Name: IDX_3aa4c8c63867c680ea06368b2f; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "IDX_3aa4c8c63867c680ea06368b2f" ON public.members USING btree ("updatedAt");


--
-- Name: IDX_4a2de01dd822709aef7b262094; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "IDX_4a2de01dd822709aef7b262094" ON public.comments USING btree ("updatedAt");


--
-- Name: IDX_58922f98620290bccd835b5497; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "IDX_58922f98620290bccd835b5497" ON public.applications USING btree ("updatedAt");


--
-- Name: IDX_5be83717ca7a7a0032b649696f; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "IDX_5be83717ca7a7a0032b649696f" ON public.candidates USING btree ("updatedAt");


--
-- Name: IDX_d3ad23adb2c1949d730a714479; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "IDX_d3ad23adb2c1949d730a714479" ON public.interview_selections USING btree ("interviewsId");


--
-- Name: IDX_eff752629806b7f834e6392ea8; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "IDX_eff752629806b7f834e6392ea8" ON public.interviews USING btree ("updatedAt");


--
-- Name: interview_selections FK_1b8dfc096390728dc39f96d652c; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.interview_selections
    ADD CONSTRAINT "FK_1b8dfc096390728dc39f96d652c" FOREIGN KEY ("applicationsId") REFERENCES public.applications(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: comments FK_225e629ae6b8d8e593a86b8298a; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comments
    ADD CONSTRAINT "FK_225e629ae6b8d8e593a86b8298a" FOREIGN KEY ("applicationId") REFERENCES public.applications(id) ON DELETE CASCADE;


--
-- Name: comments FK_343ff2dc7da292bd5a2b2183b4e; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comments
    ADD CONSTRAINT "FK_343ff2dc7da292bd5a2b2183b4e" FOREIGN KEY ("memberId") REFERENCES public.members(id) ON DELETE CASCADE;


--
-- Name: interviews FK_55e5f3b7320f0d8035c817e16ae; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.interviews
    ADD CONSTRAINT "FK_55e5f3b7320f0d8035c817e16ae" FOREIGN KEY ("recruitmentId") REFERENCES public.recruitments(id) ON DELETE CASCADE;


--
-- Name: applications FK_a34254e3f2b3d20f07f8dbd6322; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.applications
    ADD CONSTRAINT "FK_a34254e3f2b3d20f07f8dbd6322" FOREIGN KEY ("candidateId") REFERENCES public.candidates(id) ON DELETE CASCADE;


--
-- Name: interview_selections FK_d3ad23adb2c1949d730a714479b; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.interview_selections
    ADD CONSTRAINT "FK_d3ad23adb2c1949d730a714479b" FOREIGN KEY ("interviewsId") REFERENCES public.interviews(id);


--
-- Name: applications FK_dedab7441186221a819fb51feca; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.applications
    ADD CONSTRAINT "FK_dedab7441186221a819fb51feca" FOREIGN KEY ("recruitmentId") REFERENCES public.recruitments(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--
