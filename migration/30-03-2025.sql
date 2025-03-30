ALTER TABLE users
ADD CONSTRAINT email_format_check CHECK (email LIKE '%@%.%'),
ADD CONSTRAINT unique_email UNIQUE (email);
