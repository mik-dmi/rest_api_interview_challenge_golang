CREATE TABLE IF NOT EXISTS properties (
    name VARCHAR(255) NOT NULL,
    units VARCHAR[] NOT NULL,
    PRIMARY KEY (name)
);