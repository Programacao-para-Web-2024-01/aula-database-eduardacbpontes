IF NOT EXISTS (SELECT * FROM sys.databases WHERE name = 'web')
BEGIN
    CREATE DATABASE web;
END
GO

USE web;
GO

IF OBJECT_ID('students', 'U') IS NULL
BEGIN
    CREATE TABLE students
    (
        id    INT IDENTITY(1,1) PRIMARY KEY,
        name  NVARCHAR(255) NOT NULL,
        age   INT NOT NULL,
        email NVARCHAR(255) NOT NULL,
        phone NVARCHAR(11) NOT NULL
    );
END
GO

IF OBJECT_ID('subjects', 'U') IS NULL
BEGIN
    CREATE TABLE subjects
    (
        id       INT IDENTITY(1,1) PRIMARY KEY,
        name     NVARCHAR(255) NOT NULL,
        workload INT NOT NULL
    );
END
GO

IF OBJECT_ID('students_subjects', 'U') IS NULL
BEGIN
    CREATE TABLE students_subjects
    (
        student_id INT FOREIGN KEY REFERENCES students(id),
        subject_id INT FOREIGN KEY REFERENCES subjects(id),
        CONSTRAINT PK_students_subjects PRIMARY KEY (student_id, subject_id)
    );
END
GO

IF OBJECT_ID('professors', 'U') IS NULL
BEGIN
    CREATE TABLE professors
    (
        id    INT IDENTITY(1,1) PRIMARY KEY,
        name  NVARCHAR(255) NOT NULL,
        age   INT NOT NULL,
        email NVARCHAR(255) NOT NULL,
        phone NVARCHAR(11) NOT NULL
    );
END
GO

IF OBJECT_ID('professors_subjects', 'U') IS NULL
BEGIN
    CREATE TABLE professors_subjects
    (
        professor_id INT FOREIGN KEY REFERENCES professors(id),
        subject_id   INT FOREIGN KEY REFERENCES subjects(id),
        CONSTRAINT PK_professors_subjects UNIQUE (professor_id, subject_id)
    );
END
GO

IF OBJECT_ID('users', 'U') IS NULL
BEGIN
    CREATE TABLE users
    (
        id       INT IDENTITY(1,1) PRIMARY KEY,
        username NVARCHAR(255) NOT NULL UNIQUE,
        password NVARCHAR(255) NOT NULL
    );
END
GO
IF NOT EXISTS (SELECT * FROM sys.databases WHERE name = 'web')
BEGIN
    CREATE DATABASE web;
END
GO

USE web;
GO

IF OBJECT_ID('students', 'U') IS NULL
BEGIN
    CREATE TABLE students
    (
        id    INT IDENTITY(1,1) PRIMARY KEY,
        name  NVARCHAR(255) NOT NULL,
        age   INT NOT NULL,
        email NVARCHAR(255) NOT NULL,
        phone NVARCHAR(11) NOT NULL
    );
END
GO

IF OBJECT_ID('subjects', 'U') IS NULL
BEGIN
    CREATE TABLE subjects
    (
        id       INT IDENTITY(1,1) PRIMARY KEY,
        name     NVARCHAR(255) NOT NULL,
        workload INT NOT NULL
    );
END
GO

IF OBJECT_ID('students_subjects', 'U') IS NULL
BEGIN
    CREATE TABLE students_subjects
    (
        student_id INT FOREIGN KEY REFERENCES students(id),
        subject_id INT FOREIGN KEY REFERENCES subjects(id),
        CONSTRAINT PK_students_subjects PRIMARY KEY (student_id, subject_id)
    );
END
GO

IF OBJECT_ID('professors', 'U') IS NULL
BEGIN
    CREATE TABLE professors
    (
        id    INT IDENTITY(1,1) PRIMARY KEY,
        name  NVARCHAR(255) NOT NULL,
        age   INT NOT NULL,
        email NVARCHAR(255) NOT NULL,
        phone NVARCHAR(11) NOT NULL
    );
END
GO

IF OBJECT_ID('professors_subjects', 'U') IS NULL
BEGIN
    CREATE TABLE professors_subjects
    (
        professor_id INT FOREIGN KEY REFERENCES professors(id),
        subject_id   INT FOREIGN KEY REFERENCES subjects(id),
        CONSTRAINT PK_professors_subjects UNIQUE (professor_id, subject_id)
    );
END
GO

IF OBJECT_ID('users', 'U') IS NULL
BEGIN
    CREATE TABLE users
    (
        id       INT IDENTITY(1,1) PRIMARY KEY,
        username NVARCHAR(255) NOT NULL UNIQUE,
        password NVARCHAR(255) NOT NULL
    );
END
GO
