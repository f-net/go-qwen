CREATE TABLE `assistant_vector_store` (
  assistant_id bigint NOT NULL  COMMENT '助手id',
  vector_store_id bigint NOT NULL  COMMENT '存储id'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;