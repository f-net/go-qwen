CREATE TABLE IF NOT EXISTS thread (
    id INTEGER PRIMARY KEY AUTOINCREMENT, -- 唯一标识符
    name TEXT NOT NULL,                   -- 会话名
    remote_id TEXT DEFAULT '' NOT NULL,   -- 远程id
    assistant_id INTEGER NOT NULL,        -- 助手id
    created_at INTEGER NOT NULL,          -- 创建时间
    updated_at INTEGER NOT NULL           -- 更新时间
);