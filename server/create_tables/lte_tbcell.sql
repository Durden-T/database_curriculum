CREATE TABLE `tbcell` (
  `city` varchar(255) DEFAULT NULL,
  `sector_id` varchar(255) NOT NULL,
  `sector_name` varchar(255) NOT NULL,
  `enodeb_id` int NOT NULL,
  `enodeb_name` varchar(255) DEFAULT NULL,
  `earfcn` int NOT NULL,
  `pci` int DEFAULT NULL,
  `pss` int DEFAULT NULL,
  `sss` int DEFAULT NULL,
  `tac` int DEFAULT NULL,
  `vendor` varchar(255) DEFAULT NULL,
  `longitude` float NOT NULL,
  `latitude` float NOT NULL,
  `style` varchar(255) DEFAULT NULL,
  `azimuth` float NOT NULL,
  `height` float DEFAULT NULL,
  `electtilt` float DEFAULT NULL,
  `mechtilt` float DEFAULT NULL,
  `totletilt` float DEFAULT NULL,
  PRIMARY KEY (`sector_id`),
  KEY `tbcell_enodeb_id_index` (`enodeb_id`),
  KEY `tbcell_enodeb_name_index` (`enodeb_name`),
  KEY `tbcell_sector_name_index` (`sector_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci


create definer = root@localhost trigger lte.cal_pss_and_totletilt
    before insert
    on lte.tbcell
    for each row
begin
    set new.pss = new.pci % 3;
    set new.totletilt = new.electtilt + new.mechtilt;
end;


load data infile 'D:/tbcell.csv' ignore into table lte.tbcell
    fields terminated by ',' optionally enclosed by '"'
    lines terminated by '\r\n'
    ignore 1 lines
