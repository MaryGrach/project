-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

DROP TABLE IF EXISTS user_data CASCADE;
DROP TABLE IF EXISTS city CASCADE;
DROP TABLE IF EXISTS country CASCADE;
DROP TABLE IF EXISTS sight CASCADE;
DROP TABLE IF EXISTS journey CASCADE;
DROP TABLE IF EXISTS journey_sight CASCADE;
DROP TABLE IF EXISTS image_data CASCADE;
DROP TABLE IF EXISTS feedback CASCADE;
DROP TABLE IF EXISTS profile_data CASCADE;

CREATE TABLE country(
    id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY ,
    country text NOT NULL UNIQUE
);

CREATE TABLE city(
    id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY ,
    city text NOT NULL,
    country_id integer REFERENCES country(id)
);

CREATE TABLE user_data (
    id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY ,
    email text NOT NULL UNIQUE,
    passwrd text NOT NULL
);

CREATE TABLE profile_data (
    user_id integer REFERENCES user_data(id),
    username text UNIQUE,
    avatar text,
	bio text
);

CREATE TABLE sight(
    id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY ,
    rating float NOT NULL CHECK (rating > 0 AND rating <= 5),
    name text NOT NULL,
    description text,
    city_id integer REFERENCES city (id),
    country_id integer REFERENCES country (id),
	UNIQUE (name, city_id),
    latitude REAL,
    longitude REAL
);

CREATE TABLE image_data(
    id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY ,
    "path" text NOT NULL UNIQUE,
    sight_id integer REFERENCES sight(id)
);

CREATE TABLE journey(
    id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY ,
	name VARCHAR(255) NOT NULL UNIQUE,
    user_id integer REFERENCES user_data(id),
    description VARCHAR(255)
);

CREATE TABLE journey_sight(
    id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY ,
    journey_id integer REFERENCES journey(id),
    sight_id integer REFERENCES sight(id),
    priority integer NOT NULL,
	UNIQUE (journey_id, sight_id)
);

CREATE TABLE feedback(
    id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY ,
    user_id integer REFERENCES user_data(id),
    sight_id integer REFERENCES sight(id),
    rating integer NOT NULL CHECK (rating > 0 AND rating <= 5),
    feedback text NOT NULL
);

INSERT INTO country(country) VALUES
('Россия'),
('Беларусь'),
('Татарстан'),
('Крым'),
('Дагестан'),
('Абхазия');

INSERT INTO city (city, country_id) VALUES 
('Москва', 1),
('Вольск', 1),
('Тамбов', 1),
('Бахчисарай', 4),
('Евпатория', 1),
('Балаклава', 2),
('Казань', 3),
('Салта', 4),
('Мир', 2),
('Гудаута', 6),
('Дербент', 5),
('Нижний Новгород', 1),
('Ицари', 5),
('Сулакский каньон', 5);



INSERT INTO sight(name, description, rating, city_id, country_id, latitude, longitude) VALUES 
    (
        'У дяди Вани',	
        'Ресторан с видом на Сталинскую высотку.', 
        4,
        1,
        1,
        55.768329,
        37.597468
    ),
	(
		'Государственный музей изобразительных искусств имени А.С. Пушкина',
		'Музей.',
        5,
		1,
		1,
        55.747277,
        37.605194
	),
	(
		'МГТУ им. Н. Э. Баумана',
		'Хороший вуз.',
        5,
		1,
		1,
        55.766471,
        37.683446
	),
	(
		'Вкусно - и точка',
		'Неплохое кафе, вызывает гастрит.',
        2,
		1,
		1,
        55.771585,
        37.68173
	),
	(
		'Краеведческий музей',
		'Один из самых больших провинциальных музеев краеведческого профиля.',
        4,
		2,
		1,
        53.148541,
        48.456204
	),
	(
		'Спасо-Преображенский кафедральный собор',
		'Спасо-Преображенский кафедральный собор расположен в центре города и является первым каменным храмом Тамбова и старейшим в Тамбовской обл.',
        5,
		3,
		1,
        52.727371,
        41.45911
	),
	(
		'Мирский замок',
		'Памятник архитектуры, внесён в список Всемирного наследия ЮНЕСКО.',
        4,
		9,
		2,
        53.900162,
        27.551518
	),
	(
		'Чуфут-Кале',
		'Пещерный город в Крыму. Топ.',
        5,
		4,
		4,
        44.733255,
        33.934201
	),
	(
		'Сасык-Сиваш',
		'Розовое озеро. Оно реально розовое.',
        4,
		5,
		4,
        45.181475,
        33.576833
	),
	(
		'Крепость Чембело',
		'Остатки крепости.',
        2,
		6,
        4,
        44.512136,
        33.598321
	),
	(
		'Мечеть Кул Шариф',
		'Главная джума-мечеть республики Татарстан и города Казани.',
        1,
		7,
		3,
        55.798399,
        49.105147
	),
	(
		'Салтинский Подземный Водопад',
		'Единственный в России подземный водопад.',
        4,
		8,
        5,
        42.380378,
        47.042068
	),
	(
        'Озеро Рица',	
        'Рица — горное озеро ледниково-тектонического происхождения на Западном Кавказе, в Гудаутском районе Абхазии', 
        5,
        10,
        6,
        43.48013,
        40.542047
    ),
	(
        'Архитектурный комплекс Цитадель Нарын-Кала',	
        'Древняя дербентская крепость, возведённая по повелению персидского правителя Хосрова I Ануширвана в VI веке, включена ЮНЕСКО в Список Всемирного наследия.', 
        3,
        11,
        5,
        42.05534,
        48.276883
    ),
	( 
        'Сторожевые башни Северного Кавказа',	
        'Хорошо сохранившиеся родовые башни XIV–XVI веков, которые выполняли роль жилища и защиты от врагов.', 
        3,
        13,
        5,
        42.086155,
        47.603964
    ),
	( 
        'Сулакский каньон',	
        'У истоков реки Сулак берёт начало уникальный каньон. Давным-давно бурная река расколола гору, разделив Салатавский и Гимринский хребты.', 
        4,
        14,
        5,
        43.017452,
        46.824505
    ),
	(
        'Стрелка Волги и Оки',	
        'Место, где реки Ока и Волга, сливаясь, образуют живописный треугольный мыс, называют Стрелкой. Это природная достопримечательность Нижнего Новгорода.', 
        5,
        12,
        1,
        56.335336,
        43.967053
    ),
	(
        'Чкаловская лестница',	
        'Чкаловская лестница - один из символов Нижнего Новгорода. Между Верхневолжской и Нижневолжской набережными находится интересная нижегородская достопримечательность, которая видна даже на космических снимках', 
        4,
        12,
        1,
        56.330091,
        44.008924
    ),
	(
        'Нижегородский Кремль',	
        'Нижегородский кремль – древняя крепость и одновременно главная историческая достопримечательность Нижнего Новгорода', 
        4,
        12,
        1,
        56.328437,
        44.003111
    );


INSERT INTO image_data(path, sight_id) VALUES 
('public/1.jpg', 1),
('public/2.jpg', 2),
('public/3.jpg', 3),
('public/4.jpg', 4),
('public/5.jpg', 5),
('public/6.jpg', 6),
('public/7.jpg', 7),
('public/8.jpg', 8),
('public/9.jpg', 9),
('public/10.jpg', 10),
('public/11.jpg', 11),
('public/12.jpg', 12),
('public/13.jpg', 13),
('public/14.jpg', 14),
('public/15.jpg', 15),
('public/16.jpg', 16),
('public/17.jpg', 17),
('public/18.jpg', 18);


CREATE OR REPLACE FUNCTION create_profile()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO profile_data (user_id, username, bio, avatar)
    VALUES (NEW.id, NEW.email, '', '');
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER create_profile_trigger
AFTER INSERT ON user_data
FOR EACH ROW
EXECUTE FUNCTION create_profile();


-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
