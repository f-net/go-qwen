CREATE TABLE `assistant_thread` (
    id bigint NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name varchar(255) NOT NULL COMMENT '会话名称',
    assistant_id  bigint NOT NULL COMMENT '助手id',
    remote_id varchar(255) NOT NULL COMMENT '远程id',
    created_at bigint NOT NULL COMMENT '创建时间',
    updated_at bigint NOT NULL COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;