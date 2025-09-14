[mockery site](https://github.com/vektra/mockery)

## Install & Use
```
$> go install github.com/vektra/mockery/v3@v3.5.4
$> mockery init     // .mockery.yml 파일 생성
$> mockery          //  mock객체 생성 

$> go get github.com/stretchr/testify
$> go get github.com/stretchr/testify/mock
```
## .mockery.yml 변경
```
all: false
dir: './internal/mocks/{{.SrcPackageName}}' // 생성한 mock객체 저장 위치
filename: 'Mock_{{.SrcPackageName}}.go'     // mock 파일명
force-file-write: true
formatter: goimports
include-auto-generated: false
log-level: info
structname: 'Mock{{.InterfaceName}}'
recursive: true                   // 하위 디렉토리 인터페이스의 목 객체 생성
exclude-subpkg-regex: ["handlers"]  // handler mock 객체 생성 제외
require-template-schema-exists: true
template: testify
template-schema: '{{.Template}}.schema.json'
packages:
  github.com/cloneOsima/bigLand/backend:
    config:
      all: true
```
이외는 초기값에 수정안함

## repo test
- 트랜잭션 등의 로직 없이 단순 CRUD 라면 단위 테스트는 크게 의미가 없다고 생각함 
- 실제 DB랑 연결해서 단위테스트 진행하는건 DB부하가 큼
- repo 테스트는 나중에 총합 테스트에서 DB와 연동을 볼 것