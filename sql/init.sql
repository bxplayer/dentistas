CREATE TABLE IF NOT EXISTS  dentista (
  id INT NOT NULL AUTO_INCREMENT,
  apellido VARCHAR(100) NOT NULL,
  nombre VARCHAR(100) NOT NULL,
  matricula VARCHAR(50) NOT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS  paciente (
  id INT NOT NULL AUTO_INCREMENT,
  nombre VARCHAR(50) NOT NULL,
  apellido VARCHAR(50) NOT NULL,
  domicilio VARCHAR(100) NOT NULL,
  DNI VARCHAR(20) NOT NULL,
  fecha_alta DATE NOT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS  turno (
  id INT NOT NULL AUTO_INCREMENT,
  descripcion TEXT,
  fecha_hora DATETIME NOT NULL,
  paciente_id INT NOT NULL,
  dentista_id INT NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (paciente_id) REFERENCES paciente(id),
  FOREIGN KEY (dentista_id) REFERENCES dentista(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


INSERT INTO dentista (apellido, nombre, matricula)
VALUES
('Gonzalez', 'Juan', 'A123'),
('Perez', 'Ana', 'B456'),
('Rodriguez', 'Carlos', 'C789');

INSERT INTO paciente (nombre, apellido, domicilio, DNI, fecha_alta)
VALUES
('Maria', 'Fernandez', 'Calle Falsa 123', '12345678', '2023-01-01'),
('Pedro', 'Lopez', 'Avenida Siempreviva 456', '23456789', '2023-02-01'),
('Laura', 'Garcia', 'Boulevard de los Sueños Rotos 789', '34567890', '2023-03-01');

INSERT INTO turno (descripcion, fecha_hora, paciente_id, dentista_id)
VALUES
('Control de rutina', '2023-10-01 10:00:00', 1, 1),
('Extracción muela', '2023-10-02 11:00:00', 2, 2),
('Limpieza dental', '2023-10-03 12:00:00', 3, 3);