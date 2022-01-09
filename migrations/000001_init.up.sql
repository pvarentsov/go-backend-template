CREATE TABLE users(
    user_id        BIGSERIAL                      ,
    firstname      VARCHAR (50)           NOT NULL,
    lastname       VARCHAR (50)                   ,
    email          VARCHAR (100)  UNIQUE  NOT NULL,
    password       VARCHAR (100)          NOT NULL,

    PRIMARY KEY (user_id)
);
