PGDMP                         z            mcp_budget_app    14.2    14.2     ?           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            ?           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            ?           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            ?           1262    16395    mcp_budget_app    DATABASE     r   CREATE DATABASE mcp_budget_app WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'English_United States.1252';
    DROP DATABASE mcp_budget_app;
                postgres    false            ?            1259    16401    tblm_expenses    TABLE     ?   CREATE TABLE public.tblm_expenses (
    id integer NOT NULL,
    tanggal_ex date NOT NULL,
    jumlah_ex integer NOT NULL,
    catatan_ex character varying(255) NOT NULL,
    created_date date NOT NULL,
    updated_date date NOT NULL
);
 !   DROP TABLE public.tblm_expenses;
       public         heap    postgres    false            ?            1259    16407    tblm_expenses_id_seq    SEQUENCE     ?   ALTER TABLE public.tblm_expenses ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.tblm_expenses_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);
            public          postgres    false    210            ?            1259    16396    tblm_income    TABLE     ?   CREATE TABLE public.tblm_income (
    id integer NOT NULL,
    tanggal_in date NOT NULL,
    jumlah_in integer NOT NULL,
    catatan_in character varying(255) NOT NULL,
    created_date date NOT NULL,
    updated_date date NOT NULL
);
    DROP TABLE public.tblm_income;
       public         heap    postgres    false            ?            1259    16406    tblm_income_id_seq    SEQUENCE     ?   ALTER TABLE public.tblm_income ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.tblm_income_id_seq
    START WITH 3
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);
            public          postgres    false    209            ?          0    16401    tblm_expenses 
   TABLE DATA           j   COPY public.tblm_expenses (id, tanggal_ex, jumlah_ex, catatan_ex, created_date, updated_date) FROM stdin;
    public          postgres    false    210   z       ?          0    16396    tblm_income 
   TABLE DATA           h   COPY public.tblm_income (id, tanggal_in, jumlah_in, catatan_in, created_date, updated_date) FROM stdin;
    public          postgres    false    209   ?       ?           0    0    tblm_expenses_id_seq    SEQUENCE SET     B   SELECT pg_catalog.setval('public.tblm_expenses_id_seq', 4, true);
          public          postgres    false    212            ?           0    0    tblm_income_id_seq    SEQUENCE SET     @   SELECT pg_catalog.setval('public.tblm_income_id_seq', 4, true);
          public          postgres    false    211            d           2606    16405     tblm_expenses tblm_expenses_pkey 
   CONSTRAINT     ^   ALTER TABLE ONLY public.tblm_expenses
    ADD CONSTRAINT tblm_expenses_pkey PRIMARY KEY (id);
 J   ALTER TABLE ONLY public.tblm_expenses DROP CONSTRAINT tblm_expenses_pkey;
       public            postgres    false    210            b           2606    16400    tblm_income tblm_income_pkey 
   CONSTRAINT     Z   ALTER TABLE ONLY public.tblm_income
    ADD CONSTRAINT tblm_income_pkey PRIMARY KEY (id);
 F   ALTER TABLE ONLY public.tblm_income DROP CONSTRAINT tblm_income_pkey;
       public            postgres    false    209            ?   M   x?3?4000?#NNN###]]C$&?Q???NCS?N??Ē??TLL8@ssp?????? ???      ?   q   x?3?4202?50?5??42600??M?N?S?K,?THLI?KGQgr!8F??& ?!??Y?9??&?1?9 ???)???p	C$&?	?c??q????'fe?P???? ?*     