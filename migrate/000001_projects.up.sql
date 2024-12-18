CREATE TABLE IF NOT EXISTS projects (
             guid UUID PRIMARY KEY NOT NULL,
             alias VARCHAR(255) NOT NULL,
             name VARCHAR(255) NOT NULL,
             info TEXT,
             created_at TIMESTAMP NOT NULL,
             updated_at TIMESTAMP NOT NULL,
             deleted_at TIMESTAMP
);
