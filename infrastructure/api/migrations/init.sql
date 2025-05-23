-- Создание таблицы пользователей
CREATE TABLE IF NOT EXISTS "user" (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Создание таблицы типов растений
CREATE TABLE IF NOT EXISTS plant_type (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL,
    description TEXT,
    watering_frequency INTEGER NOT NULL, -- в днях
    fertilizing_frequency INTEGER NOT NULL -- в днях
);

-- Создание таблицы растений
CREATE TABLE IF NOT EXISTS plants (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    species VARCHAR(255) NOT NULL,
    watering_frequency INTEGER NOT NULL,
    last_watered TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- Создание таблицы уведомлений
CREATE TABLE IF NOT EXISTS notifications (
    id SERIAL PRIMARY KEY,
    plant_id INTEGER REFERENCES plants(id) ON DELETE CASCADE,
    type VARCHAR(20) NOT NULL, -- 'watering' или 'fertilizing'
    message TEXT NOT NULL,
    is_read BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Создание индексов
CREATE INDEX IF NOT EXISTS idx_notifications_plant_id ON notifications(plant_id);
CREATE INDEX IF NOT EXISTS idx_notifications_created_at ON notifications(created_at);

-- Добавление тестовых данных
INSERT INTO plant_type (name, description, watering_frequency, fertilizing_frequency)
VALUES 
    ('Кактус', 'Не требует частого полива', 14, 30),
    ('Фикус', 'Любит умеренный полив', 7, 14),
    ('Орхидея', 'Требует особого ухода', 5, 21)
ON CONFLICT (name) DO NOTHING; 