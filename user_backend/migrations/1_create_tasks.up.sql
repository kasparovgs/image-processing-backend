CREATE TABLE tasks (
    uuid UUID PRIMARY KEY,
    user_id TEXT NOT NULL,
    status TEXT NOT NULL,
    base64image TEXT,
    filter_name TEXT NOT NULL,
    filter_parameters JSONB NOT NULL DEFAULT '{}'::jsonb
);