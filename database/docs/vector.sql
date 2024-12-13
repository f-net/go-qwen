CREATE TABLE IF NOT EXISTS vector_store (
    id INTEGER PRIMARY KEY AUTOINCREMENT, -- 唯一标识符
    vector_store_app_id TEXT DEFAULT NULL, -- 远程文件id
    object TEXT DEFAULT NULL,              -- 对象
    name TEXT DEFAULT NULL,                -- 名称
    description TEXT,                      -- 描述
    usage_bytes INTEGER DEFAULT NULL,      -- 使用空间
    file_total_bytes INTEGER NOT NULL,     -- 文件大小
    file_counts TEXT,                      -- 文件数量
    expires_after TEXT,                    -- 待定
    expires_at INTEGER DEFAULT NULL,       -- 待定
    assistant_amount INTEGER NOT NULL,     -- 待定
    file_amount INTEGER NOT NULL,          -- 待定
    created_at INTEGER NOT NULL,           -- 创建时间
    updated_at INTEGER NOT NULL            -- 更新时间
);