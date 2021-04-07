CREATE TABLE `tbc2i3` (
                          `sector_a` varchar(255) NOT NULL,
                          `sector_b` varchar(255) NOT NULL,
                          `sector_c` varchar(255) NOT NULL,
                          UNIQUE KEY `tbc2i3_pk` (`sector_a`,`sector_b`,`sector_c`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

create
    definer = root@localhost procedure triple_analysis(IN rate float)
begin
    truncate tbc2i3;
    insert into tbc2i3
    with t as
             (
                 select serving_sector,interfering_sector
                 from tbc2inew
                 where prb_abs6>rate
                 union
                 select interfering_sector,serving_sector
                 from tbc2inew
                 where prb_abs6>rate
             )
    select distinct a.serving_sector,a.interfering_sector,b.serving_sector
    from t as a,t as b
    where a.serving_sector=b.interfering_sector
      and a.serving_sector<>b.serving_sector
      and a.serving_sector>a.interfering_sector
      and a.interfering_sector>b.serving_sector;
end;

