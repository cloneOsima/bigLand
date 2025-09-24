[sqlc site](https://docs.sqlc.dev/en/latest/index.html)

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
  - engine: "postgresql"        // slqc 관련 쿼리/스키마 파일 위치 지정 
    queries: "./internal/repositories/query.sql"
    schema: "./internal/repositories/schema.sql"
    gen:
      go:                       // sqlc 사용할 언어 및 산출물 저장 위치 지정
        package: "sqlc"             
        out: "./internal/sqlc"
        sql_package: "pgx/v5"
overrides:
    go: null
plugins: []
rules: []
options: {}
```
이외는 초기값에 수정안함

## sqlc struct vs models package 
- sqlc 가 생성하는 struct를 기본적으로 사용하되, 특정 정보만이 필요한 기능에 대해서는 service layer 에서 매핑하도록 구현 예정

