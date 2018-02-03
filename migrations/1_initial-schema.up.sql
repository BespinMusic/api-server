Create Table songs (
    id serial primary key,
    cd text not null,
    nm text,
    --ref int references another.table,
    created timestamptz default now()
);
Create Index on songs (cd);