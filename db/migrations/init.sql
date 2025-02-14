-- db/migrations/init.sql
CREATE TABLE IF NOT EXISTS sensor_data (
    id SERIAL PRIMARY KEY,
    value1_load_cell FLOAT NOT NULL,
    value2_load_cell FLOAT NOT NULL,
    voltage1_load_cell FLOAT NOT NULL,
    voltage2_load_cell_real FLOAT NOT NULL,
    voltage2_load_cell_imag FLOAT NOT NULL,
    temperature FLOAT NOT NULL,
    value_pulse FLOAT NOT NULL,
    voltage_pulse FLOAT NOT NULL,
    value_muscle_activity FLOAT NOT NULL,
    voltage_muscle_activity FLOAT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);