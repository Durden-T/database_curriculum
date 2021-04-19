CREATE TABLE `tbmrodata` (
  `start_time` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `serving_sector` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `interfering_sector` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `lte_sc_rsrp` float DEFAULT NULL,
  `lte_nc_rsrp` float DEFAULT NULL,
  `lte_nc_earfcn` float DEFAULT NULL,
  `lte_nc_pci` float DEFAULT NULL,
  KEY `tbmrodata_serving_sector_index` (`serving_sector`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci



