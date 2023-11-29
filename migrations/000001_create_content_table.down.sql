CREATE TABLE IF NOT EXISTS content (
    id              SERIAL PRIMARY KEY,
    copyright       TEXT,
    explanation     TEXT,
    hdurl           VARCHAR(150),
    media_type      VARCHAR(10),
    service_version VARCHAR(5),
    title           VARCHAR(70),
    url             VARCHAR(150),
    date            VARCHAR(10)
);
