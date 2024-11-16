CREATE TABLE IF NOT EXISTS notify
(
    id       BIGSERIAL PRIMARY KEY,
    user_id  BIGINT       NOT NULL,
    status   TEXT      NOT NULL,
    notification TEXT        NOT NULL,
    created_at timestamp NOT NULL,
    expired_at timestamp,
);