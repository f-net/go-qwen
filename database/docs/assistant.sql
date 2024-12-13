CREATE TABLE IF NOT EXISTS assistant (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL, -- '名称'
    instructions TEXT NOT NULL, -- '指令'
    assistant_app_id TEXT DEFAULT '' NOT NULL, -- '远程assistant_id'
    model TEXT DEFAULT '' NOT NULL, -- '模型'
    tools TEXT NOT NULL, -- '工具'
    tool_resources TEXT NOT NULL, -- '用具资源'
    remark TEXT NOT NULL, -- '备注'
    created_at INTEGER NOT NULL, -- '创建时间'
    updated_at INTEGER NOT NULL -- '更新时间'
);