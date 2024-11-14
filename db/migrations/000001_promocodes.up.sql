CREATE TABLE reward (
    id serial PRIMARY KEY,
    title VARCHAR (255) NOT NULL,
    description VARCHAR (255)
);

CREATE TABLE promocode (
    id serial PRIMARY KEY,
    promocode VARCHAR (255),
    reward_id integer REFERENCES reward (id) ON DELETE CASCADE,
    expires timestamp without time zone DEFAULT NULL,
    max_uses integer NOT NULL,
    remain_uses integer NOT NULL
);


CREATE TABLE rewards (
    id serial PRIMARY KEY,
    user_id integer,
    promocode_id integer REFERENCES promocode (id) ON DELETE CASCADE,
    "timestamp" timestamp without time zone NOT NULL
);
