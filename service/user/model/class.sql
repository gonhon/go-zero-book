CREATE TABLE `class` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `num` int(3) NOT NULL DEFAULT 0 COMMENT '人数',
    `name` varchar(255)  NOT NULL DEFAULT '' COMMENT '班级名称',
    `code` varchar(225)  NOT NULL DEFAULT '' COMMENT '班级编号',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `code_unique` (`code`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 ; 

--  goctl model mysql datasource -url="$datasource" -table="class" -c -dir .

-- goctl model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/go-zero-book" -table="class"  -dir="${modeldir}" -cache=false --style=goZero

-- goctl model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/go-zero-book" -table="class"  -dir="${modeldir}" -cache=false --style=goZero

