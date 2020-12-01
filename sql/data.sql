--
-- PostgreSQL database dump
--

-- Dumped from database version 12.5 (Ubuntu 12.5-0ubuntu0.20.04.1)
-- Dumped by pg_dump version 12.5 (Ubuntu 12.5-0ubuntu0.20.04.1)

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
-- Data for Name: customer; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.customer (id, name, phone_number, password, created_at, flag_status) FROM stdin;
18becdfc-a170-404e-ae77-ae2399054e4b	reno syahputra	081231651890	$2a$10$nxIGqRzeONGaQrUeR0c95.xnj4v4LMh11WAqzL0TW8BXotjpUnM.C	2020-12-01 16:22:42.157834+07	0
a0c63e7d-880a-45d7-ad1c-b756b6855d1e	rikka	081231651888	$2a$10$3M9x9yTBh5CvyXovAHDY.eJiXHbaGk3oLqKgJ4PBpadqOsdCz/CqC	2020-12-01 16:22:53.264076+07	0
\.


--
-- Data for Name: owner; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.owner (id, name, email, password, created_at, flag_status) FROM stdin;
04ca286c-98e5-4d09-8e85-290f27111273	reno syahputra	reno@gmail.com	$2a$10$Rq1OqpsUw2MZQZji5jbo2OEkRgeVl0iibZPGXdme9NqS6BqhiUJoC	2020-12-01 16:10:02.423665+07	0
89852144-d958-4223-86a0-5300c1b27c25	rikka takahasi	rikka@gmail.com	$2a$10$0BOVlfDXfRqZm8B158iyd.O0ORt96qh6M4mq48J7n2UNNXt17z3bG	2020-12-01 16:10:16.933132+07	0
\.


--
-- Data for Name: store; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.store (id, owner_id, name, description, image_url, latitude, longitude, created_at, flag_status) FROM stdin;
fbe9330e-18d5-4633-8df1-ca5cc408f531	89852144-d958-4223-86a0-5300c1b27c25	Dolan Cafe	kafe untuk nongkrong bareng temen	https://dolanyok.com/wp-content/uploads/2019/12/Silol-Kopi-Eatery.jpg	-7.814995	110.35595	2020-12-01 16:43:23.742714+07	0
be5ca5ab-1e36-4bb9-9af0-bbb36c7c8ab4	89852144-d958-4223-86a0-5300c1b27c25	Miku Cafe	kafe untuk para wibu	https://dolanyok.com/wp-content/uploads/2019/12/Silol-Kopi-Eatery.jpg	-7.814995	110.35595	2020-12-01 16:43:54.463236+07	0
\.


--
-- Data for Name: seats; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.seats (id, store_id, name, description, "position", price, created_at, flag_status) FROM stdin;
42fed3d7-8f9d-4d8f-81ed-92c19f5b2020	fbe9330e-18d5-4633-8df1-ca5cc408f531	Kursi depan Kafe jogja	kursi bagian depan dari kafe untuk nongkrong bareng temen	1	20000	2020-12-01 22:14:29.032545+07	0
c0fbc833-163f-4f8f-80a8-dc63912ca403	fbe9330e-18d5-4633-8df1-ca5cc408f531	Kursi tengah Kafe jogja	kursi bagian engah dari kafe untuk nongkrong bareng temen	2	20000	2020-12-01 22:14:48.318453+07	0
\.


--
-- Data for Name: booking; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.booking (id, customer_id, seats_id, price, total, duration_from, duration_to, payment_status, created_at, flag_status) FROM stdin;
6c04ae8e-cdd6-425b-88f8-85ff36ed4a11	a0c63e7d-880a-45d7-ad1c-b756b6855d1e	42fed3d7-8f9d-4d8f-81ed-92c19f5b2020	20000	20000	2020-09-29 08:00:00+07	2020-09-29 09:00:00+07	1	2020-12-01 22:34:37.590842+07	0
d8121c41-20be-48b3-8d0b-fc7f56085aca	a0c63e7d-880a-45d7-ad1c-b756b6855d1e	c0fbc833-163f-4f8f-80a8-dc63912ca403	30000	30000	2020-09-29 10:00:00+07	2020-09-29 11:00:00+07	1	2020-12-01 22:35:53.045563+07	0
\.


--
-- Data for Name: product; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.product (id, store_id, name, description, image_url, price, created_at, flag_status) FROM stdin;
67f3545e-325f-4343-94ef-0beea2c7a9db	fbe9330e-18d5-4633-8df1-ca5cc408f531	Espresso coklat 1	kopi paling populer dari kafe untuk nongkrong bareng temen	/img/cafe_product.jpg	14000	2020-12-01 22:19:15.555827+07	0
c1062030-c275-489b-9ec2-75a63ea232a9	fbe9330e-18d5-4633-8df1-ca5cc408f531	Espresso coklat 2	kopi paling populer dari kafe untuk nongkrong bareng temen	/img/cafe_product.jpg	14000	2020-12-01 22:19:18.948329+07	0
\.


--
-- Data for Name: booking_detail; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.booking_detail (id, booking_id, product_id, price, quantity, sub_total, created_at, flag_status) FROM stdin;
343e4d9c-a493-4f7b-92cd-8a2d7662ef81	6c04ae8e-cdd6-425b-88f8-85ff36ed4a11	c1062030-c275-489b-9ec2-75a63ea232a9	14000	2	28000	2020-12-01 22:40:10.7722+07	0
\.


--
-- Data for Name: transaction; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.transaction (id, booking_id, customer_id, total, payment_type, payment_status, payment_order_id, payment_id, payment_time, approval_code, bank_name, va, cstore_code, cstore_name, created_at, flag_status) FROM stdin;
521efd36-a396-44ac-801b-8a5233ac86ab	6c04ae8e-cdd6-425b-88f8-85ff36ed4a11	a0c63e7d-880a-45d7-ad1c-b756b6855d1e	20000	0	1									2020-12-01 22:46:14.942402+07	0
c421a9bd-b921-4bf6-80f7-9a51211fc9ac	d8121c41-20be-48b3-8d0b-fc7f56085aca	a0c63e7d-880a-45d7-ad1c-b756b6855d1e	20000	0	1									2020-12-01 23:06:22.085259+07	0
\.


--
-- PostgreSQL database dump complete
--

