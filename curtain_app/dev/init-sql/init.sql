/*
CREATE USER 'curtain_app_user'@localhost IDENTIFIED BY 'curtain_app_PW#0';
GRANT ALL ON curtain_app.* TO curtain_app_user@localhost;
*/
CREATE DATABASE curtain_app;
USE curtain_app;
CREATE TABLE wake_up_log (id int not null auto_increment primary key, date timestamp not null, msg text);
