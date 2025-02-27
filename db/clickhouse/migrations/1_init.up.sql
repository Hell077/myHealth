CREATE TABLE IF NOT EXISTS sugar_log
(
    id            UUID DEFAULT generateUUIDv4(),  -- Уникальный идентификатор записи
    user_id       UInt64,                           -- ID пользователя (из таблицы пользователей)
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

CREATE TABLE IF NOT EXISTS insulin_log
(
    id UUID DEFAULT generateUUIDv4(),
    user_id UInt64,
    insulinType Enum8('long'=2,'short'=1),
    unit UInt8,
    name String,
    created_at DATETIME DEFAULT now()
)
    engine = MergeTree
        ORDER BY id;

CREATE TABLE if not exists food_log
(
    id UUID DEFAULT generateUUIDv4(),
    user_id UInt64,
    weight Float32,
    meal_time DateTime,
    food_name String,
    carbs Float32,
    protein Float32,
    fat Float32,
    created_at DateTime DEFAULT now()
)
    ENGINE = MergeTree()
        ORDER BY (user_id, meal_time);



