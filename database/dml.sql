SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL,ALLOW_INVALID_DATES';

CREATE SCHEMA IF NOT EXISTS `securitydb` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci ;
USE `securitydb` ;

-- -----------------------------------------------------
-- Table `securitydb`.`tb_user`
-- -----------------------------------------------------
CREATE  TABLE IF NOT EXISTS `securitydb`.`tb_user` (
  `id_user` BIGINT(20) NOT NULL AUTO_INCREMENT ,
  `user_name` VARCHAR(60) NOT NULL ,
  `password` VARCHAR(60) NULL ,
  `email` VARCHAR(60) NULL  ,
  `created_user` BIGINT(20) NULL ,
  `created_date` DATE NULL ,
  `updated_user` BIGINT(20) NULL ,
  `updated_date` DATE NULL ,
  `origin` CHAR(1) NULL COMMENT '0=mail\n1=facebook' ,
  `status` CHAR(1) NULL COMMENT '1=create\n0=delete\n' ,
  `image` VARCHAR(255) NULL ,
  `fisrt_name` VARCHAR(255) NULL ,
  `last_name` VARCHAR(255) NULL ,
  PRIMARY KEY (`id_user`) )
ENGINE = InnoDB;

CREATE UNIQUE INDEX `user_name_UNIQUE` ON `securitydb`.`tb_user` (`user_name` ASC) ;


-- -----------------------------------------------------
-- Table `securitydb`.`tb_role`
-- -----------------------------------------------------
CREATE  TABLE IF NOT EXISTS `securitydb`.`tb_role` (
  `role` VARCHAR(15) NOT NULL ,
  `status` VARCHAR(45) NULL ,
  `created_user` VARCHAR(45) NULL ,
  `created_date` VARCHAR(45) NULL ,
  `updated_user` VARCHAR(45) NULL ,
  `updated_date` VARCHAR(45) NULL ,
  PRIMARY KEY (`role`) )
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `securitydb`.`tb_module`
-- -----------------------------------------------------
CREATE  TABLE IF NOT EXISTS `securitydb`.`tb_module` (
  `id_module` INT NOT NULL AUTO_INCREMENT ,
  `name` VARCHAR(45) NULL ,
  PRIMARY KEY (`id_module`) )
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `securitydb`.`tb_resource`
-- -----------------------------------------------------
CREATE  TABLE IF NOT EXISTS `securitydb`.`tb_resource` (
  `id_resource` BIGINT(20) NOT NULL AUTO_INCREMENT ,
  `parent` BIGINT(20) NULL ,
  `name` VARCHAR(45) NULL ,
  `label` VARCHAR(45) NULL ,
  `order` INT NULL ,
  `url` VARCHAR(255) NULL ,
  `show` CHAR(1) NULL ,
  `created_user` BIGINT(20) NULL ,
  `created_date` DATE NULL ,
  `updated_user` BIGINT(20) NULL ,
  `updated_date` DATE NULL ,
  `id_module` INT NOT NULL ,
  PRIMARY KEY (`id_resource`) ,
  CONSTRAINT `fk_tb_resource_tb_module1`
    FOREIGN KEY (`id_module` )
    REFERENCES `securitydb`.`tb_module` (`id_module` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

CREATE INDEX `fk_tb_resource_tb_module1_idx` ON `securitydb`.`tb_resource` (`id_module` ASC) ;


-- -----------------------------------------------------
-- Table `securitydb`.`tb_perfil`
-- -----------------------------------------------------
CREATE  TABLE IF NOT EXISTS `securitydb`.`tb_perfil` (
  `id_resource` BIGINT(20) NOT NULL ,
  `role` VARCHAR(15) NOT NULL ,
  PRIMARY KEY (`id_resource`, `role`) ,
  CONSTRAINT `fk_tb_perfil_tb_resource1`
    FOREIGN KEY (`id_resource` )
    REFERENCES `securitydb`.`tb_resource` (`id_resource` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_tb_perfil_tb_role1`
    FOREIGN KEY (`role` )
    REFERENCES `securitydb`.`tb_role` (`role` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

CREATE INDEX `fk_tb_perfil_tb_role1_idx` ON `securitydb`.`tb_perfil` (`role` ASC) ;


-- -----------------------------------------------------
-- Table `securitydb`.`tb_user_rol`
-- -----------------------------------------------------
CREATE  TABLE IF NOT EXISTS `securitydb`.`tb_user_rol` (
  `role` VARCHAR(15) NOT NULL ,
  `id_user` BIGINT(20) NOT NULL ,
  PRIMARY KEY (`role`, `id_user`) ,
  CONSTRAINT `fk_tb_user_rol_tb_role`
    FOREIGN KEY (`role` )
    REFERENCES `securitydb`.`tb_role` (`role` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_tb_user_rol_tb_user1`
    FOREIGN KEY (`id_user` )
    REFERENCES `securitydb`.`tb_user` (`id_user` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

CREATE INDEX `fk_tb_user_rol_tb_user1_idx` ON `securitydb`.`tb_user_rol` (`id_user` ASC) ;

USE `securitydb` ;

CREATE USER 'admin' IDENTIFIED BY 'Zion@fast';

GRANT ALL ON `securitydb`.* TO 'admin';
GRANT SELECT, INSERT, TRIGGER ON TABLE `securitydb`.* TO 'admin';
GRANT SELECT, INSERT, TRIGGER, UPDATE, DELETE ON TABLE `securitydb`.* TO 'admin';

SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;