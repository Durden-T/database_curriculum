CREATE TABLE `tbkpi` (
  `start_time` varchar(255) NOT NULL,
  `enodeb_name` varchar(255) DEFAULT NULL,
  `sector_description` varchar(255) DEFAULT NULL,
  `sector_name` varchar(255) NOT NULL,
  `rcc_conn_succ` int DEFAULT NULL,
  `rcc_conn_att` int DEFAULT NULL,
  `rcc_conn_rate` decimal(7,4) DEFAULT NULL,
  PRIMARY KEY (`sector_name`,`start_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci


create definer = root@localhost trigger lte.kpi_cal_rcc_conn_rate_and_format
    before insert
    on lte.tbkpi
    for each row
begin
    if new.rcc_conn_att != 0 then
        set new.rcc_conn_rate = new.rcc_conn_succ / new.rcc_conn_att;
    end if;
    if instr(new.start_time, '/') = 3 then
        set @year = substr(new.start_time, 7, 4);
        set @month = left(new.start_time, 2);
        set @day = substr(new.start_time, 4, 2);
        set new.start_time = concat(@year, '-', @month, '-', @day, right(new.start_time, 9));
    end if;
end;

load data infile 'D:/tbkpi.csv' ignore into table lte.tbkpi
    fields terminated by ',' optionally enclosed by '"'
    lines terminated by '\r\n'
    ignore 1 lines
