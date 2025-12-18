-- UP
CREATE TABLE IF NOT EXISTS pets (
    id INT PRIMARY KEY,
    name TEXT NOT NULL,
    date_of_birth DATE,
    breed TEXT,
    sex TEXT,
    microchip_id TEXT,
    status TEXT,
    litter_name TEXT,
    slug TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
