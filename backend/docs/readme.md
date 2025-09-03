# 新着情報 read function requirements

## Required dataset
- id : int // table's primary key
- 事故発生日 : Date (ex: 2024年10月10日)
- 登録原因 : String　（ex: 作業員の男性が死亡）
- 投稿登録日 : Date (ex: 9月3日)
- 住所 : String (ex: 静岡県浜松市中央区湖東町)

## SQL for creating table 
```
create table regist_info (
    id SERIAL PRIMARY KEY,
    accident_date varchar(255),
    accident_cause varchar(255),
    registartion_date varchar(255), 
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