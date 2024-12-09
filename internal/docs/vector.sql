CREATE TABLE `vector_store` (
    id bigint NOT NULL AUTO_INCREMENT PRIMARY KEY,
    vector_store_app_id varchar(255) DEFAULT NULL  COMMENT '远程文件id',
    object varchar(255) DEFAULT NULL  COMMENT '对象',
    name varchar(255) DEFAULT NULL  COMMENT '名称',
    description text  COMMENT '描述',
    usage_bytes int DEFAULT NULL  COMMENT '使用空间',
    file_total_bytes bigint NOT NULL  COMMENT '文件大小',
    file_counts longtext  COMMENT '文件数量',
    expires_after longtext  COMMENT '待定',
    expires_at int DEFAULT NULL  COMMENT '待定',
    assistant_amount bigint NOT NULL  COMMENT '待定',
    file_amount bigint NOT NULL  COMMENT '待定',
    created_at bigint NOT NULL ,
    updated_at bigint NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;