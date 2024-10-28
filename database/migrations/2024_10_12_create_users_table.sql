CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,        -- ID único do usuário
    username VARCHAR(255) NOT NULL,           -- Nome de usuário (não nulo)
    password VARCHAR(255) NOT NULL,           -- Senha do usuário (não nulo)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Data de criação
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- Data de atualização
    created_by INT DEFAULT NULL,              -- ID do usuário que criou (nulo por padrão)
    updated_by INT DEFAULT NULL,              -- ID do usuário que atualizou (nulo por padrão)
    status ENUM('active', 'inactive') NOT NULL DEFAULT 'active' -- Status do usuário, pode ser 'active' ou 'inactive'
);
