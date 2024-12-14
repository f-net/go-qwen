CREATE TABLE IF NOT EXISTS message (
    id INTEGER PRIMARY KEY AUTOINCREMENT, -- '唯一标识符'
    question TEXT NOT NULL,               -- '提问'
    answer TEXT NOT NULL,                 -- '回答'
    remote_id TEXT DEFAULT '' NOT NULL,   -- '远程id'
    assistant_id INTEGER NOT NULL,        -- '助手id'
    thread_id INTEGER NOT NULL,           -- '会话id'
    created_at INTEGER NOT NULL,          -- '创建时间'
    updated_at INTEGER NOT NULL           -- '更新时间'
);