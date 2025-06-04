-- Aluno: Walisson Fagundes Santana

-- Tabela professores:
CREATE TABLE professor (
    id SERIAL PRIMARY KEY NOT NULL,
    nome VARCHAR(10) NOT NULL,
    profissão VARCHAR(50) NOT NULL,
    ativo BOOLEAN NOT NULL
);

-- Tabela materia: 
CREATE TABLE materia (
    idMateria SERIAL PRIMARY KEY NOT NULL,
    nome VARCHAR(100) NOT NULL,
    cargaHoraria TIME NOT NULL,
    idProfessor INT NOT NULL,
    FOREIGN KEY (idProfessor) REFERENCES professor(id) ON DELETE CASCADE
);

-- Q01)
ALTER TABLE professor
    ADD COLUMN cpf VARCHAR(11) UNIQUE;

-- Q02)
ALTER TABLE professor
    ALTER COLUMN nome TYPE VARCHAR(100);

-- Q03)
ALTER TABLE materia
    DROP COLUMN cargaHoraria;

-- Q04)
INSERT INTO professor (nome, profissão, ativo, cpf) VALUES
('Wevton', 'Ciência da Computação', TRUE, '89087665517'),
('João', 'Ciência da Computação', TRUE, '76489653456'),
('José', 'Ciência da Computação', TRUE, '97965434569'),
('Lombo', 'Ciência da Computação', TRUE, '09827654309');

-- Q05)
UPDATE professor
    SET nome = 'Aquiles'
    WHERE id = 3;

UPDATE professor
    SET nome = 'Bruno'
    WHERE id = 2;

-- Q06)
INSERT INTO materia (nome, idProfessor) VALUES
('Algoritmos e Estrutura de Dados', 1),
('POO', 2),
('Banco de Dados', 3),
('Programação Web', 4);

-- Q07)
DELETE FROM materia
    WHERE idMateria = 4;

-- Q08)
SELECT * FROM professor
    WHERE ativo = TRUE;

-- Q09)
SELECT materia.*
FROM materia
JOIN professor ON materia.idProfessor = professor.id
WHERE professor.nome = 'Wevton'
ORDER BY materia.nome;

-- Q10)
SELECT * FROM professor ORDER BY id;
SELECT * FROM materia ORDER BY idMateria;
