CREATE TABLE openai_file (
    id INT PRIMARY KEY AUTO_INCREMENT,
    file_app_id VARCHAR(255)  COMMENT '文件id',
    file_name VARCHAR(255)  COMMENT '文件名',
    extension VARCHAR(255)  COMMENT '文件格式',
    file_bytes BIGINT  COMMENT '文件大小',
    vector_store_id INT  COMMENT '向量id',
    status VARCHAR(255)  COMMENT '状态',
    purpose VARCHAR(255)  COMMENT '待定',
    created_at bigint NOT NULL COMMENT '创建时间',
    updated_at bigint NOT NULL COMMENT '更新时间'
);