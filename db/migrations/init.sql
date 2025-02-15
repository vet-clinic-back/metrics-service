-- Создаем последовательность
CREATE SEQUENCE IF NOT EXISTS owner_id_seq START 1;
CREATE SEQUENCE IF NOT EXISTS veterinarian_id_seq START 1;
CREATE SEQUENCE IF NOT EXISTS medical_record_id_seq START 1;
CREATE SEQUENCE IF NOT EXISTS pet_id_seq START 1;
CREATE SEQUENCE IF NOT EXISTS medical_entry_id_seq START 1;
CREATE SEQUENCE IF NOT EXISTS device_id_seq START 1;

-- Создаем таблицу
CREATE TABLE IF NOT EXISTS owner (
    id INTEGER PRIMARY KEY DEFAULT nextval('owner_id_seq'),
    full_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    phone VARCHAR(20) NOT NULL,
    password_hash VARCHAR(255) NOT NULL
);

-- Добавляем комментарии
COMMENT ON COLUMN owner.id IS 'Идентификатор владельца';
COMMENT ON COLUMN owner.full_name IS 'ФИО владельца';
COMMENT ON COLUMN owner.email IS 'Электронная почта';
COMMENT ON COLUMN owner.phone IS 'Телефон владельца';
COMMENT ON COLUMN owner.password_hash IS 'Пароль (Hash)';


CREATE TABLE IF NOT EXISTS veterinarian (
    id INTEGER PRIMARY KEY DEFAULT nextval('veterinarian_id_seq'),
    full_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    phone VARCHAR(20) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    position VARCHAR(100),
    clinic_number VARCHAR(20)
);

COMMENT ON COLUMN veterinarian.id IS 'Идентификатор ветеринара';
COMMENT ON COLUMN veterinarian.full_name IS 'ФИО ветеринара';
COMMENT ON COLUMN veterinarian.email IS 'Электронная почта';
COMMENT ON COLUMN veterinarian.phone IS 'Телефон ветеринара';
COMMENT ON COLUMN veterinarian.password_hash IS 'Пароль (Hash)';
COMMENT ON COLUMN veterinarian.position IS 'Должность ветеринара';
COMMENT ON COLUMN veterinarian.clinic_number IS 'Номер поликлиники';


CREATE TABLE IF NOT EXISTS pet (
    id INTEGER PRIMARY KEY DEFAULT nextval('pet_id_seq'),
    animal_type VARCHAR(50) NOT NULL,
    name VARCHAR(100) NOT NULL,
    gender VARCHAR(10) CHECK (gender IN ('Male', 'Female')),
    age INTEGER NOT NULL,
    weight NUMERIC(5, 2),
    condition TEXT,
    behavior TEXT,
    research_status TEXT
);

COMMENT ON COLUMN pet.id IS 'Идентификатор питомца';
COMMENT ON COLUMN pet.animal_type IS 'Вид животного';
COMMENT ON COLUMN pet.name IS 'Имя питомца';
COMMENT ON COLUMN pet.gender IS 'Пол питомца';
COMMENT ON COLUMN pet.age IS 'Возраст питомца';
COMMENT ON COLUMN pet.weight IS 'Вес питомца';
COMMENT ON COLUMN pet.condition IS 'Состояние питомца';
COMMENT ON COLUMN pet.behavior IS 'Поведение питомца';
COMMENT ON COLUMN pet.research_status IS 'Статус исследования';


CREATE TABLE IF NOT EXISTS medical_record (
    id INTEGER PRIMARY KEY DEFAULT nextval('medical_record_id_seq'),
    veterinarian_id INTEGER REFERENCES veterinarian(id),
    owner_id INTEGER REFERENCES owner(id),
    pet_id INTEGER REFERENCES pet(id)
);

COMMENT ON COLUMN medical_record.id IS 'Идентификатор медкарты';
COMMENT ON COLUMN medical_record.veterinarian_id IS 'Ссылка на ветеринара';
COMMENT ON COLUMN medical_record.owner_id IS 'Ссылка на владельца';
COMMENT ON COLUMN medical_record.pet_id IS 'Ссылка на питомца';



CREATE TABLE IF NOT EXISTS device (
    id INTEGER PRIMARY KEY DEFAULT nextval('device_id_seq'),
    unique_number VARCHAR(100),
    status VARCHAR(50) NOT NULL
);

COMMENT ON COLUMN device.id IS 'Идентификатор устройства';
COMMENT ON COLUMN device.unique_number IS 'Уникальный номер устройства';
COMMENT ON COLUMN device.status IS 'Статус устройства';



CREATE TABLE IF NOT EXISTS medical_entry (
    id INTEGER PRIMARY KEY DEFAULT nextval('medical_entry_id_seq'),
    entry_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    description TEXT,
    disease TEXT,
    vaccinations TEXT,
    recommendation TEXT,
    medical_record_id INTEGER REFERENCES medical_record(id),
    device_number INTEGER REFERENCES device(id),
    veterinarian_id INTEGER REFERENCES veterinarian(id)
);


COMMENT ON COLUMN medical_entry.id IS 'Идентификатор записи в медкарте';
COMMENT ON COLUMN medical_entry.entry_date IS 'Дата и время записи';
COMMENT ON COLUMN medical_entry.description IS 'Описание записи (возможно, XML)';
COMMENT ON COLUMN medical_entry.disease IS 'Заболевание';
COMMENT ON COLUMN medical_entry.vaccinations IS 'Вакцинации';
COMMENT ON COLUMN medical_entry.recommendation IS 'Рекомендации';
COMMENT ON COLUMN medical_entry.medical_record_id IS 'Ссылка на медкарту';
COMMENT ON COLUMN medical_entry.device_number IS 'Номер устройства';
COMMENT ON COLUMN medical_entry.veterinarian_id IS 'Идентификатор ветеринара';

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

COMMENT ON TABLE sensor_data IS 'Данные с медицинских сенсоров';
COMMENT ON COLUMN sensor_data.value1_load_cell IS 'Значение с нагрузочной ячейки 1';
COMMENT ON COLUMN sensor_data.value2_load_cell IS 'Значение с нагрузочной ячейки 2';
COMMENT ON COLUMN sensor_data.voltage1_load_cell IS 'Напряжение с нагрузочной ячейки 1';
COMMENT ON COLUMN sensor_data.voltage2_load_cell_real IS 'Напряжение с нагрузочной ячейки 2 (реальное)';
COMMENT ON COLUMN sensor_data.voltage2_load_cell_imag IS 'Напряжение с нагрузочной ячейки 2 (имагинарное)';
COMMENT ON COLUMN sensor_data.temperature IS 'Температура';
COMMENT ON COLUMN sensor_data.value_pulse IS 'Значение пульса';
COMMENT ON COLUMN sensor_data.voltage_pulse IS 'Напряжение пульса';
COMMENT ON COLUMN sensor_data.value_muscle_activity IS 'Значение активности мышц';
COMMENT ON COLUMN sensor_data.voltage_muscle_activity IS 'Напряжение активности мышц';