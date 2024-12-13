CREATE TABLE IF NOT EXISTS openai_file (
    id INTEGER PRIMARY KEY AUTOINCREMENT, -- 唯一标识符
    file_app_id TEXT NOT NULL,            -- 文件id
    file_name TEXT NOT NULL,              -- 文件名
    extension TEXT NOT NULL,              -- 文件格式
    file_bytes INTEGER NOT NULL,          -- 文件大小
    vector_store_id INTEGER,              -- 向量id
    status TEXT NOT NULL,                 -- 状态
    purpose TEXT,                         -- 待定
    created_at INTEGER NOT NULL,          -- 创建时间
    updated_at INTEGER NOT NULL           -- 更新时间
);