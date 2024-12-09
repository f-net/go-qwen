CREATE TABLE `assistant` (
    id bigint NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name varchar(255) NOT NULL COMMENT '名称',
    instructions text NOT NULL COMMENT '指令',
    description text NOT NULL COMMENT '描述',
    assistant_app_id varchar(255) DEFAULT '' NOT NULL COMMENT '远程assistant_id',
    model varchar(255) DEFAULT '' NOT NULL COMMENT '模型',
    tools longtext NOT NULL COMMENT '工具',
    tool_resources longtext NOT NULL COMMENT '用具资源',
    remark text default '' NOT NULL COMMENT '备注',
    created_at bigint NOT NULL COMMENT '创建时间',
    updated_at bigint NOT NULL COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;