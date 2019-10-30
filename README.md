# UserRegistrationMircoServiceGO
Example of micro Service using golang

create table userinfo.UserRegistration(
   id VARCHAR(40) NOT NULL,
   email VARCHAR(100) NOT NULL,
   name VARCHAR(40) NOT NULL,
   password VARCHAR(40) NOT NULL,
   mobile VARCHAR(40) NOT NULL,
   city VARCHAR(40) NOT NULL,
   PRIMARY KEY ( email )
);
