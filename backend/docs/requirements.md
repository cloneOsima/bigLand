# Required functions
- get post list (50 limit)  : fin
- get post info             : fin
- create a new post         : fin 
- comment
- delete post (admin)       
- delete comment (admin)
- create a new account      : in progress
- user login

## DDL for creating table 
```
CREATE TABLE users (
    user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    last_login_at TIMESTAMP WITH TIME ZONE,
    is_active BOOLEAN DEFAULT TRUE
);

CREATE TABLE posts (
    post_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    content TEXT NOT NULL,
    incident_date DATE NOT NULL,
    posted_date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    latitude DOUBLE PRECISION NOT NULL,
    longtitude DOUBLE PRECISION NOT NULL,
    address_text TEXT,
    location GEOMETRY(Point, 4326) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE
);

CREATE TABLE comments (
    comment_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    post_id UUID REFERENCES posts(post_id) ON DELETE CASCADE,
    author_id UUID REFERENCES users(user_id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    posted_date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    is_active BOOLEAN DEFAULT TRUE
);
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

# 新着情報 read function requirements
- GET Function: return list of posts (limit 50)
- Get 50 post list ordered by posted date

## Required dataset
#### post list dataset
dataset for list(recently added post) 
- post_id : 투고 id 
- posted_date : 투고일
- address_text : 사고지 주소

# Post read function requirements
- GET function: return detailed information about the post
- Get postID from http request + Search database by using postID + Return result 
- UUID validation check is needed

## Required dataset
#### single post dataset
dataset for a single post(detailed information)
- post_id : 투고 id
- content : 투고 내용
- incident_date : 사고 발생일
- posted_date : 투고일
- latitude : 위도
- longitude : 경도 
- address_text : 사고지 주소
- location : gis 데이터 
- is_active : 활성화/비활성화 플래그

# Post 投稿 function requirementes
- POST function : create new post
- Insert new row at post table by using given input value 
- Input value validation check is needed

## Required dataset
#### post dataset
- content : 투고 내용
- incident_date : 사고 발생일
- latitude : 위도
- longitude : 경도 
- address_text : 사고지 주소

# ユーザー 加入 function requirementes
- POST function : create new user account
- Insert new row at post table by using given input value 
- Input value validation check is needed

## Required dataset
#### post dataset
- content : 투고 내용
- incident_date : 사고 발생일
- latitude : 위도
- longitude : 경도 
- address_text : 사고지 주소

# Viewport 

## Required dataset

#### viewport dataset
dataset for viewport
- post_id : 투고 id 
- latitude : 위도
- longitude : 경도 
- location : gis 데이터 