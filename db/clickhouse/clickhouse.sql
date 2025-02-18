CREATE TABLE IF NOT EXISTS sugarLevel
(
    id            UUID DEFAULT generateUUIDv4(),  -- Уникальный идентификатор записи
    user_id       UUID,                           -- ID пользователя (из таблицы пользователей)
    measurement_time DateTime DEFAULT now(),      -- Время измерения сахара в крови
    sugar_value   Float32,                        -- Уровень сахара в крови (ммоль/л)
    meal_time     Enum8('before_meal' = 1, 'after_meal' = 2, 'random' = 3),  -- Время измерения (до еды, после еды, случайное)
    created_at    DateTime DEFAULT now()          -- Дата создания записи
)
    ENGINE = MergeTree()
        ORDER BY (user_id, measurement_time);


CREATE TABLE IF NOT EXISTS users
(
    id         UUID DEFAULT generateUUIDv4(),  -- Уникальный идентификатор пользователя
    login      String,                         -- Логин пользователя
    full_name  String,                         -- Полное имя
    age        UInt8,                          -- Возраст (может быть полезен для анализа)
    gender     Enum8('male' = 1, 'female' = 2, 'other' = 3),  -- Пол пользователя
    created_at DateTime DEFAULT now()          -- Дата регистрации
)
    ENGINE = MergeTree()
        ORDER BY id;
