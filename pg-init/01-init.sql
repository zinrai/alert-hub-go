CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS alerts (
    id SERIAL PRIMARY KEY,
    subject VARCHAR(200) NOT NULL,
    body TEXT NOT NULL,
    identifier VARCHAR(50) NOT NULL UNIQUE,
    urgency VARCHAR(10) NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'OPEN',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE FUNCTION encrypt_body(p_body TEXT, p_key TEXT) RETURNS TEXT AS $$
BEGIN
    RETURN encode(pgp_sym_encrypt(p_body, p_key), 'base64');
END;
$$ LANGUAGE plpgsql SECURITY DEFINER;

CREATE OR REPLACE FUNCTION decrypt_body(p_encrypted_body TEXT, p_key TEXT) RETURNS TEXT AS $$
BEGIN
    RETURN pgp_sym_decrypt(decode(p_encrypted_body, 'base64')::bytea, p_key);
END;
$$ LANGUAGE plpgsql SECURITY DEFINER;
