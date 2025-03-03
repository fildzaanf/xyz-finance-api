CREATE DATABASE xyz_finance;

USE xyz_finance;
-- Tabel Users
CREATE TABLE users (
    id VARCHAR(36) PRIMARY KEY,
    email VARCHAR(255) NOT NULL,  
    password VARCHAR(255) NOT NULL,
    nik VARCHAR(50) NOT NULL UNIQUE,   
    full_name VARCHAR(255) NOT NULL,
    legal_name VARCHAR(255) NOT NULL,
    birth_place VARCHAR(100) NOT NULL,
    birth_date DATE NOT NULL,
    ktp_photo TEXT NOT NULL,
    selfie_photo TEXT NOT NULL,
    salary BIGINT NOT NULL,
    role ENUM('user') DEFAULT 'user',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Tabel Loans
CREATE TABLE loans (
    id VARCHAR(36) PRIMARY KEY,
    user_id VARCHAR(36) NOT NULL,
    tenor INT NOT NULL,
    limit_amount INT NOT NULL,
    status ENUM('valid', 'invalid') DEFAULT 'valid',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX idx_loans_user_id ON loans(user_id);

-- Tabel Transactions
CREATE TABLE transactions (
    id VARCHAR(36) PRIMARY KEY,
    loan_id VARCHAR(36) NOT NULL,
    asset_name VARCHAR(100) NOT NULL,
    total_amount INT NOT NULL,
    tenor INT NOT NULL,
    otr_price INT NOT NULL,
    admin_fee INT NOT NULL,
    interest INT NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (loan_id) REFERENCES loans(id) ON DELETE CASCADE
);

CREATE INDEX idx_transactions_loan_id ON transactions(loan_id);

-- Tabel Installments
CREATE TABLE installments (
    id VARCHAR(36) PRIMARY KEY,
    transaction_id VARCHAR(36) NOT NULL,
    installment_number INT NOT NULL,
    amount INT NOT NULL,
    due_date TIMESTAMP,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (transaction_id) REFERENCES transactions(id) ON DELETE CASCADE
);

CREATE INDEX idx_installments_transaction_id ON installments(transaction_id);

-- Tabel Payments
CREATE TABLE payments (
    id VARCHAR(36) PRIMARY KEY,
    installment_id VARCHAR(36) NOT NULL,
    gross_amount INT NOT NULL,
    status ENUM('deny', 'success', 'cancel', 'expire', 'pending') DEFAULT 'pending',
    payment_url TEXT,
    token TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (installment_id) REFERENCES installments(id) ON DELETE CASCADE
);

CREATE INDEX idx_payments_installment_id ON payments(installment_id);
