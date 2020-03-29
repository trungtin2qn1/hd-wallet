CREATE OR REPLACE FUNCTION next_id(OUT result bigint, seq text) AS $$
DECLARE
    our_epoch bigint := 1314220021721;
    seq_id bigint;
    now_millis bigint;
    shard_id int := 5;
BEGIN
    SELECT nextval(seq) % 1024 INTO seq_id;
    SELECT FLOOR(EXTRACT(EPOCH FROM clock_timestamp())) INTO now_millis;
    result := (now_millis - our_epoch)*1000 << 23;
    result := result | (shard_id <<10);
    result := result | (seq_id);
    
END;
    $$ LANGUAGE PLPGSQL;


CREATE OR REPLACE FUNCTION convertTVkdau (x text) RETURNS text AS
$$
DECLARE
 cdau text; kdau text; r text;
BEGIN
 cdau = 'áàảãạâấầẩẫậăắằẳẵặđéèẻẽẹêếềểễệíìỉĩịóòỏõọôốồổỗộơớờởỡợúùủũụưứừửữựýỳỷỹỵÁÀẢÃẠÂẤẦẨẪẬĂẮẰẲẴẶĐÉÈẺẼẸÊẾỀỂỄỆÍÌỈĨỊÓÒỎÕỌÔỐỒỔỖỘƠỚỜỞỠỢÚÙỦŨỤƯỨỪỬỮỰÝỲỶỸỴ';
 kdau = 'aaaaaaaaaaaaaaaaadeeeeeeeeeeeiiiiiooooooooooooooooouuuuuuuuuuuyyyyyaaaaaaaaaaaaaaaaadeeeeeeeeeeeiiiiiooooooooooooooooouuuuuuuuuuuyyyyy';
 r = x;
 FOR i IN 0..length(cdau)
 LOOP
 r = replace(r, substr(cdau,i,1), substr(kdau,i,1));
 END LOOP;
 RETURN r;
END;
$$ LANGUAGE plpgsql;
--Generate new func insta5

--Create table admin_users
--Drop table admin_users
drop table if exists admin_users;
drop sequence if exists admin_users_id_seq;
CREATE SEQUENCE admin_users_id_seq;

CREATE TABLE admin_users (
    id bigint NOT NULL DEFAULT next_id('admin_users_id_seq') primary key,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    email text,
    first_name text,
    last_name text,
    password text,
    username text,
    role text,
    last_login timestamp,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

--Create table movies
--Drop table movies
drop table if exists movies;
drop sequence if exists movies_id_seq;
CREATE SEQUENCE movies_id_seq;

CREATE TABLE movies (
    id bigint NOT NULL DEFAULT next_id('movies_id_seq') primary key,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    name text,
    description text,
    image text,
    trailer text,
    started_at timestamp,
    duration integer,
    rating float,
    views integer,
    manual_point integer,
    algorithm_point integer,
    type integer
);