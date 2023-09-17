DROP TABLE IF EXISTS pacientes;

CREATE TABLE pacientes (
                           id         	INT PRIMARY KEY AUTO_INCREMENT,
                           nombre        VARCHAR(128) NOT NULL,
                           apellido      VARCHAR(128) NOT NULL,
                           domicilio     VARCHAR(128) NOT NULL,
                           dni           VARCHAR(128) NOT NULL DEFAULT FALSE,
                           fecha_de_alta VARCHAR(128) NOT NULL
);

INSERT INTO pacientes (nombre, apellido, domicilio, dni, fecha_de_alta)
VALUES ('Pedrito', 'Gonzalez', 'Maipu 6785', '45.890.236', '2020-01-01');

INSERT INTO pacientes (nombre, apellido, domicilio, dni, fecha_de_alta)
VALUES ('Manuelito', 'Ramirez', 'Ind. 2344', '39.836.654', '2021-07-12');

DROP TABLE IF EXISTS odontologos;

CREATE TABLE odontologos (
                             id            INT PRIMARY KEY AUTO_INCREMENT,
                             nombre        VARCHAR(128) NOT NULL,
                             apellido      VARCHAR(128) NOT NULL,
                             matricula     VARCHAR(128) NOT NULL
);

INSERT INTO odontologos (nombre, apellido, matricula)
VALUES ('George', 'Sanchez', 'B385934');

INSERT INTO odontologos (nombre, apellido, matricula)
VALUES ('William', 'Rolando', 'A789496');

DROP TABLE IF EXISTS turnos;

CREATE TABLE turnos (
                        id            SERIAL PRIMARY KEY AUTO_INCREMENT,
                        paciente_id   INT NOT NULL,
                        odontologo_id INT NOT NULL,
                        FOREIGN KEY (paciente_id) REFERENCES pacientes(id),
                        FOREIGN KEY (odontologo_id) REFERENCES odontologos(id),
                        fecha_hora    VARCHAR(128) NOT NULL,
                        descripcion   VARCHAR(200) NOT NULL
);

INSERT INTO turnos (paciente_id, odontologo_id, fecha_hora, descripcion)
VALUES (1, 2, '2023-10-23', 'Limpieza de caries');