use securitydb;

INSERT INTO tb_role (role, status) VALUES ('ROLE_ADMIN','CREATED');
INSERT INTO tb_role (role, status) VALUES ('ROLE_USER','CREATED');

INSERT INTO `tb_user` (user_name, password, email, status,created_user, created_date, updated_user, updated_date,origin) VALUES ('edumar111','$2a$10$J11dsG20VhVJJYFapTNbxOtFgWyBOgTFKdRrsYENFMFI4VmYSJMh.','edumar111@gmail.com','1','SYSTEM','2020-09-05','SYSTEM','2020-09-05','0');

INSERT INTO `tb_user` (user_name, password, email, status, created_user, created_date, updated_user, updated_date,origin) VALUES ('admin','$2a$10$J11dsG20VhVJJYFapTNbxOtFgWyBOgTFKdRrsYENFMFI4VmYSJMh.','edumar111@gmail.com','1','SYSTEM','2020-09-05','SYSTEM','2020-09-05','0');

INSERT INTO `tb_user_rol` (id_user, role) VALUES (1,'ROLE_USER');
INSERT INTO `tb_user_rol` (id_user, role) VALUES (2,'ROLE_ADMIN');
INSERT INTO `tb_user_rol` (id_user, role) VALUES (2,'ROLE_USER');