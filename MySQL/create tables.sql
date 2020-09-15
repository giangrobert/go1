Use test_db;
create table if not exists agenda(
    id bigint unsigned not null auto_increment,
    nombre varchar(255) not null,
    direccion varchar(255) not null,
    correo_electronico varchar(255) not null,
    primary key(id)
);


DROP TABLE IF EXISTS `employee`;
CREATE TABLE `employee` (
  `id` int(6) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL,
  `city` varchar(30) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;