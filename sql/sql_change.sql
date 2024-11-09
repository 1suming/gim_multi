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
alter table message add column target_id  bigint  NOT NULL default 0 comment '目标对象id';
alter table message add  column sender_id bigint  NOT NULL default 0 comment '消息发送者';


CREATE TABLE `chatroom`
(
    `id`           bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `name`         varchar(50)   NOT NULL COMMENT '群组名称',
    `avatar_url`   varchar(255)  NOT NULL COMMENT '群组头像',
    `introduction` varchar(255)  NOT NULL COMMENT '群组简介',
    `max_user_num`     int(11) NOT NULL DEFAULT '0' COMMENT '群组人数',
    `extra`        varchar(1024) NOT NULL COMMENT '附加属性',
    `create_time`  datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`  datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='聊天室';