CREATE TABLE Collection (
    id INT PRIMARY KEY,
    title VARCHAR(255),
    year_founded INT,
    description TEXT
);

CREATE TABLE Department (
    id INT PRIMARY KEY,
    name VARCHAR(255),
    location_info TEXT
);

CREATE TABLE Exhibit (
    id INT PRIMARY KEY,
    title VARCHAR(255),
    author VARCHAR(255),
    creation_date DATE,
    description TEXT
);

CREATE TABLE Visitor (
    id INT PRIMARY KEY,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    date_of_birth DATE,
    phone_number VARCHAR(20)
);

CREATE TABLE Event (
    id INT PRIMARY KEY,
    title VARCHAR(255),
    description TEXT,
    start_date DATE,
    end_date DATE
);