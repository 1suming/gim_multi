alter table user add column password   varchar(128) NOT NULL;

alter table device add column device_unique_id   varchar(128) NOT NULL;
alter table device modify column`system_version` varchar(64) COLLATE utf8mb4_bin NOT NULL COMMENT '系统版本';
