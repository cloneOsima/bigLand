# 新着情報 read function requirements

## Required dataset

#### post dataset
- id : int // table's primary key
- 事故発生日 : Date (ex: 2024年10月10日)
- 登録原因 : String　（ex: 作業員の男性が死亡）
- 投稿登録日 : Date (ex: 9月3日)
- 住所 : String (ex: 静岡県浜松市中央区湖東町)

#### post list dataset
- id : int // table's primary key
- 投稿登録日 : Date (ex: 9月3日)
- 住所 : String (ex: 静岡県浜松市中央区湖東町)

## DDL for creating table 
```
create table post (
    id SERIAL PRIMARY KEY,
    accident_date varchar(255),
    accident_cause varchar(255),
    created_at varchar(255), 
    registed_addres varchar(255)
)
```

## Databasedriver (pq vs pgx)
- https://blog.dushyanth.in/understanding-database-drivers-in-go-the-role-of-libpq-and-pgx-in-postgresql-connections
- pgx
    - newer and is actively maintained.
    - tends to be faster and more performant than lib/pq because it has optimizations and a better understanding of PostgreSQL’s internals.


## Environmantal variable library (godotenv vs viper)
- https://enigmacamp.medium.com/difference-between-godotenv-vs-viper-in-golang-22a8d20f9835
- godotenv 
    - faster than viper 
    - no need to complex formats(only need a .env)


## Database connection pool 
- Supabase basically uses IPv6
    - Error occurred when using direct connection (Not compatible with IPv4)
        - on the server, IPv6 ping test failed
    - Use transaction pooler (Compatible with IPv4)
        - when server is started, InitDB() creates a new pool
        - when server is terminated, CloseDB() deletes the current pool 
- https://pkg.go.dev/github.com/jackc/pgx/v5/pgxpool
- https://github.com/jackc/pgx/wiki/Getting-started-with-pgx
- https://medium.com/@neelkanthsingh.jr/understanding-database-connection-pools-and-the-pgx-library-in-go-3087f3c5a0c
- https://medium.com/@lhc1990/solving-supabase-ipv6-connection-issues-the-complete-developers-guide-96f8481f42c1    // recommend to read it!