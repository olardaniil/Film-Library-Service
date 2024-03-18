CREATE TABLE actors (
    id SERIAL PRIMARY KEY,
    last_name VARCHAR(32),
    first_name VARCHAR(32),
    middle_name VARCHAR(32),
    gender_id INTEGER,
    date_birth DATE
);
INSERT INTO actors
(last_name, first_name, middle_name, gender_id, date_birth)
VALUES
    ('Старр', 'Мартин', '', 1, '1982-07-30'),
    ('Миддлдитч', 'Томас', '', 1, '1982-03-10'),
    ('О. Ян', 'Джимми', '', 1, '1987-06-11'),
    ('Диамантополос', 'Крис', '', 1, '1975-05-09'),
    ('Кузьмина', 'Ольга', 'Николаевна', 2, '1987-06-16');

CREATE TABLE movies (
    id SERIAL PRIMARY KEY,
    title VARCHAR(150) UNIQUE,
    description VARCHAR(1000),
    rating FLOAT CHECK (rating >= 0 AND rating <= 10),
    release_date DATE
);
INSERT INTO movies
(title, description, rating, release_date)
VALUES
    ('Кремниевая долина', 'История о группе гиков, готовящих к запуску собственные стартапы в высокотехнологичном центре Сан-Франциско. Главные герои сериала бесплатно проживают в доме местного миллионера, но взамен им придётся отдать по 10% прибыли от будущих проектов', 8.4, '2014-04-06'),
    ('Чебурашка', 'На небольшой приморский городок обрушивается дождь из апельсинов, а вместе с фруктами с неба падает неизвестный науке мохнатый непоседливый зверёк. Одержимое апельсинами животное оказывается в домике нелюдимого старика-садовника Геннадия, который из вредности решает оставить его жить у себя, так как местная богачка жаждет заполучить необычного зверя для своей избалованной внучки. Также эта коварная женщина, владелица кондитерской фабрики, пытается выведать секрет шоколада у хозяйки маленького магазинчика — дочери Геннадия, много лет обиженной на отца.', 7.3, '2023-01-01'),
    ('Гагарин. Первый в космосе', 'Фильм посвящён первым шагам человечества на пути освоения космоса и непосредственно судьбе первого космонавта Ю. А. Гагарина. Основной лейтмотив — борьба за право быть первым: соревнование в первом отряде космонавтов; конкуренция технологий в ракетостроении; противостояние сверхдержав — СССР и США. В первый отряд космонавтов отбирали из трёх тысяч лётчиков-истребителей по всей стране.', 7.0, '2013-01-01');

CREATE TABLE actors_movies (
    actor_id INTEGER,
    movie_id INTEGER,
    FOREIGN KEY (actor_id) REFERENCES actors(id),
    FOREIGN KEY (movie_id) REFERENCES movies(id)
);
INSERT INTO actors_movies
(actor_id, movie_id)
VALUES
    (1, 1),
    (2, 1),
    (3, 1),
    (4, 1),
    (5, 2),
    (5, 3);

CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username VARCHAR(255) UNIQUE NOT NULL,
                       password_hash VARCHAR(255) NOT NULL,
                       is_admin BOOLEAN DEFAULT FALSE
);
INSERT INTO users
(username, password_hash, is_admin)
VALUES
    ('user', '12dea96fec20593566ab75692c9949596833adc9', false),
    ('admin', 'd033e22ae348aeb5660fc2140aec35850c4da997', true);
