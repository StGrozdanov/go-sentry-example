CREATE TABLE IF NOT EXISTS courses
(
    id        SERIAL PRIMARY KEY,
    title     varchar(100) NOT NULL,
    sub_title varchar(100) NOT NULL,
    image_url varchar(250) NOT NULL
);

CREATE TABLE IF NOT EXISTS frequent_questions
(
    id       SERIAL PRIMARY KEY,
    question VARCHAR(200) NOT NULL UNIQUE,
    answer   TEXT         NOT NULL
);

INSERT INTO frequent_questions
    (question, answer)
VALUES ('Will I figure out how to workout those Sentry queries?',
        'There is a possibility. But dont do this on your own. Call the support :D'),
       ('How much time would the solution require?',
        'Between a day and a year.');

INSERT INTO courses(title, sub_title, image_url)
VALUES ('Example course', 'Example sub title of the course', 'https://image-example.com'),
       ('Example course 2', 'Example sub title of the course', 'https://image-example.com'),
       ('Example course 3', 'Example sub title of the course', 'https://image-example.com');

