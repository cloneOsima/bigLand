[sqlc site](https://docs.sqlc.dev/en/latest/index.html)

- repository layer 에서 query를 직접 사용하다보니 함수가 너무 길어지고 읽기가 힘들어지는 문제가 생김 
- query 를 별도의 sql 파일로 저장하고 sqlc 를 사용하여 해당 쿼리와 db 스키마를 활용한 구조체/함수를 생성

## Install & Use
```
$> go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
$> sqlc init     // sqlc.yaml 파일 생성
$> mockery          //  mock객체 생성 

```
## sqlc.yaml 변경
```
version: "2"
cloud:
    organization: ""
    project: ""
    hostname: ""
servers: []
sql:                
  - engine: "postgresql"        
    queries: "./internal/repositories/query.sql"
    schema: "./internal/repositories/schema.sql"
    gen:
      go:                       
        package: "sqlc"                                 // 생성된 sqlc 파일이 사용할 패키지명 제시
        out: "./internal/sqlc"                          // sqlc 파일 생성 위치 
        sql_package: "pgx/v5"                           
        overrides:                                      // sqlc 에서 pgtype으로 만드는 값들을 go 타입으로 override
        - db_type: "uuid"
          nullable: false
          go_type:
            import: "github.com/google/uuid"
            type: "UUID"
        - db_type: "pg_catalog.text"
          go_type:
            type: "string"
          nullable: true
        - db_type: "text"
          go_type:
            type: "string"
          nullable: true
        - db_type: "pg_catalog.timestamptz"
          go_type:
            type: "time.Time"
          nullable: true
        - db_type: "pg_catalog.timestamptz"
          go_type:
            type: "time.Time"
        - column: "posts.location"
          go_type: 
            type: "[]byte"
        - column: "comments.post_id"
          go_type:
            import: "github.com/google/uuid"
            type: "UUID"
        - column: "comments.author_id"
          go_type:
            import: "github.com/google/uuid"
            type: "UUID"
overrides:
    go: null
plugins: []
rules: []
options: {}

```
이외는 초기값에 수정안함

## sqlc struct vs models package 
- sqlc 가 생성하는 struct를 기본적으로 사용하되, 특정 정보만이 필요한 기능에 대해서는 service layer 에서 매핑하도록 구현 예정

