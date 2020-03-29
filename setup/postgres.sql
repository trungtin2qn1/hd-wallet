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

--Create table accounts
--Drop table accounts
drop table if exists accounts;
drop sequence if exists admin_users_id_seq;
CREATE SEQUENCE accounts_id_seq;

CREATE TABLE accounts (
    id bigint NOT NULL DEFAULT next_id('accounts_id_seq') primary key,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    private_key text,
    public_key text,
    mnemonic text
);

--Create table wallets
--Drop table wallets
drop table if exists wallets;
drop sequence if exists wallets_id_seq;
CREATE SEQUENCE wallets_id_seq;

CREATE TABLE wallets (
    id bigint NOT NULL DEFAULT next_id('wallets_id_seq') primary key,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    account_id bigint,
    foreign key (account_id) references accounts(id)
);

--Create table addresses
--Drop table addresses
drop table if exists addresses;
drop sequence if exists addresses_id_seq;
CREATE SEQUENCE addresses_id_seq;

CREATE TABLE addresses (
    id bigint NOT NULL DEFAULT next_id('addresses_id_seq') primary key,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    index int,
    currency int,
    address text,
    wallet_id bigint,
    foreign key (wallet_id) references wallets(id)
);