CREATE TABLE students (
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password varchar(255) not null,
    groups int not null unique,
    role varchar(30) not null
);

CREATE TABLE masters (
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password varchar(255) not null,
    role varchar(30) not null
);

CREATE TABLE course (
    id serial not null unique,
    name varchar(255) not null,
    description varchar(255) not null,
    student_group int references students(groups) not null
);

CREATE TABLE teacher_course (
    id serial not null unique,
    teacher_id int references masters(id) on delete cascade not null,
    course_id int references course(id) on delete cascade not null
);

CREATE TABLE books (
    id serial not null unique,
    name varchar(255) not null,
    author varchar(255) not null
);

CREATE TABLE student_books (
    id serial not null unique,
    student_id int references students(id) on delete cascade not null,
    book_id int references books(id) on delete cascade not null
);