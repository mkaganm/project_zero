CREATE TABLE verifications (
                              id SERIAL PRIMARY KEY,
                              user_id INTEGER REFERENCES users(id),
                              verification_code_hash VARCHAR(255) NOT NULL,
                              created_at TIMESTAMP DEFAULT NOW() NOT NULL
);
