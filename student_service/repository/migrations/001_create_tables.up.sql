CREATE TABLE groups (
    id UUID NOT NULL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    main_teacher_id UUID NOT NULL
);

CREATE TABLE students (
    id UUID NOT NULL PRIMARY KEY,
    first_name VARCHAR(30) NOT NULL,
    last_name VARCHAR(40) NOT NULL,
    email VARCHAR(50) NOT NULL UNIQUE,
    phone_number VARCHAR(13) NOT NULL UNIQUE,
    level INT NOT NULL,
    password VARCHAR(10) NOT NULL UNIQUE,
    group_id UUID NOT NULL REFERENCES groups(id)
);