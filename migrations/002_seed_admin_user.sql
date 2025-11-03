-- +migrate Up
INSERT INTO users (username, password, created_at, created_by)
VALUES (
    'admin',
    '$2a$10$pFdaPqglFANKAy2F4Izrv.o.V0qZkewcL.cNqPGsPZRCzRKRgcTUu', -- This is the hashed value of 'admin'
    CURRENT_TIMESTAMP,
    'system'
) ON CONFLICT (username) DO NOTHING;

-- +migrate Down
DELETE FROM users WHERE username = 'admin';