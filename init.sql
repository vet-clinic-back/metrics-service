-- migrations/init.sql
CREATE TABLE IF NOT EXISTS metrics (
    id SERIAL PRIMARY KEY,
    temperature INT NOT NULL,
    muscle_activity INT NOT NULL,
    chest_expansion1 INT NOT NULL,
    chest_expansion2 INT NOT NULL,
    pulse FLOAT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_metrics_created_at ON metrics(created_at);