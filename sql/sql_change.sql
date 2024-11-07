alter table user add column password   varchar(128) NOT NULL;

alter table device add column device_unique_id   varchar(128) NOT NULL;
alter table device modify column`system_version` varchar(64) COLLATE utf8mb4_bin NOT NULL COMMENT '系统版本';


CREATE TABLE im_recent_conversation (
                                        conversation_type tinyint comment '1:对个人；2 room 3. group',
                                        owner_uid INT NOT NULL,
                                        target_id int not null,
                                        last_message_id INT NOT NULL,
                                        last_message_content varchar(1024) ,
                                        last_time TIMESTAMP NOT NULL,

                                        unread_cnt int not null default 0 comment '消息未读数',

                                        PRIMARY KEY (`conversation_type`,`owner_uid`,target_id)
);
