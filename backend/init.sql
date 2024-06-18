GRANT ALL PRIVILEGES ON DATABASE zxcdatabase TO zxcuser;

\c zxcdatabase;

GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO zxcuser;

CREATE TABLE bin(
    bin INT PRIMARY KEY,
    bank VARCHAR(100)
);

