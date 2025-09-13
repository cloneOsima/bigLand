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
require-template-schema-exists: true
template: testify
template-schema: '{{.Template}}.schema.json'
packages:
  github.com/cloneOsima/bigLand/backend:
    config:
      all: true
```
이외는 초기값에 수정안함