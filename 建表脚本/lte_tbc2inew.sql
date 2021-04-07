CREATE TABLE `tbc2inew` (
  `serving_sector` varchar(255) NOT NULL,
  `interfering_sector` varchar(255) NOT NULL,
  `c2i_mean` float DEFAULT NULL,
  `c2i_std` float DEFAULT NULL,
  `prb_c2i9` float DEFAULT NULL,
  `prb_abs6` float DEFAULT NULL,
  PRIMARY KEY (`serving_sector`,`interfering_sector`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci


create
    definer = root@localhost procedure gen_tbc2i_new(IN minimum int)
begin
    truncate tbc2inew;
	INSERT INTO lte.tbc2inew(serving_sector, interfering_sector, c2i_mean, c2i_std)
	select serving_sector,interfering_sector,avg(lte_sc_rsrp-lte_nc_rsrp),stddev(lte_sc_rsrp-lte_nc_rsrp)
	from lte.tbmrodata
	group by serving_sector,interfering_sector
	having count(lte_sc_rsrp)>=minimum;
end;


