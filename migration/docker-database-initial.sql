create table accounts (
    id serial primary key,
    agencynumber varchar,
    accountnumber varchar
);

INSERT INTO accounts (agencynumber, accountnumber) VALUES
('0001', '123456'),
('0001', '456789');