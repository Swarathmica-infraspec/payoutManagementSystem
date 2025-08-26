CREATE TABLE IF NOT EXISTS payees (
    id SERIAL PRIMARY KEY,
    beneficiary_name TEXT NOT NULL,
    beneficiary_code TEXT NOT NULL UNIQUE,
    account_number INTEGER NOT NULL,
    ifsc_code VARCHAR(11) NOT NULL,
    bank_name VARCHAR(50) NOT NULL,
    email TEXT NOT NULL UNIQUE,
    mobile VARCHAR(10) NOT NULL,
    payee_category TEXT
);
