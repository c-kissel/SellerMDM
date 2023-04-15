-- Create contact table
CREATE TABLE contacts (
    id UUID PRIMARY KEY,
    firstname varchar(255),
    lastname varchar(255),
    email varchar(255),
    mobile varchar(255),
    messenger varchar(255),
    created_at timestamp with time zone not null,
    updated_at timestamp with time zone not null
);

-- Create sellers table
CREATE TABLE sellers (
    id UUID PRIMARY KEY,
    name varchar(255) not null,
    ogrn varchar(255),
    inn varchar(255),
    city varchar(255),
    imagenames text[],
    created_at timestamp with time zone not null,
    updated_at timestamp with time zone not null
);

-- Create reviews table
CREATE TABLE reviews (
    id UUID PRIMARY KEY,
    seller UUID references sellers(id) not null,
    contact UUID references contacts(id) not null,
    rating int,
    commentary text,
    parent UUID references reviews(id)
);

-- Create role mapping table
CREATE TABLE roles (
    id UUID PRIMARY KEY,
    seller UUID references sellers(id) not null,
    contact UUID references contacts(id) not null,
    role text
);