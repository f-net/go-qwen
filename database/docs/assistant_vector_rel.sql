CREATE TABLE IF NOT EXISTS assistant_vector_store (
      assistant_id INTEGER NOT NULL, -- '助手id'
      vector_store_id INTEGER NOT NULL, -- '存储id'
      PRIMARY KEY (assistant_id, vector_store_id) -- 如果需要组合主键
);