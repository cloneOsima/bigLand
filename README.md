# binLand

React, Golang 기반의 웹사이트트

## 🚀 시작하기

### 환경 설정

```bash
# 의존성 설치


# 환경 파일 복사 및 설정
cp .env.example .env.local
# .env.local 파일을 편집하여 필요한 설정값을 입력하세요
```

### 개발 서버 실행

```bash
# 로컬 환경 (기본)


# 개발 환경


# 운영 환경 모드

```

### 서버만 실행

```bash
# 로컬 환경


# 개발 환경


# 운영 환경

```

### 정적 파일 서버만 실행

```bash

```

## 📁 프로젝트 구조

```
backend
├── cmd/
│   └── main.go            # 라우터 초기화, 서버 진입점
├── config/                # 앱 설정 정의 ( 환경 변수, 데이터베이스 설정 등 ) 
│   └── cofnig.go          
├── handlers/              # http request 정의 패키지 (controller)
│   └── api.go             
├── models                 # Database 용 model 및 DTO 정의 패키지
│   └── models.go   
├── models                 # Database 용 model 및 DTO 정의 패키지
│   └── models.go  
├── mw                     # http 리퀘스트 대응 미들웨어 관련 패키지 ( 인증, 로깅, cors, 에러 핸드링 등 ) 
│   └── mw.go  
├── repositories           # DB 접속 및 CRUD 로직 관련 패키지 
│   └── repo.go
├── server                 # Gin router 및 서버 초기화 관련 패키지 
│   └── router.go
├── services               # BL 정의 패키지 
│   └── service.go  
└── utils.json             # 공용 함수 정의 패키지
│   └── service.go
```
```
├── pages/
│   ├── index.html          # 홈 페이지
│   ├── sample1.html        # 데이터 관리 페이지
│   └── sample2.html        # 지도 분석 페이지
├── server/
│   └── server.js           # Express 서버
├── static/
│   ├── css/               # 스타일시트
│   │   ├── app.css
│   │   ├── components.css
│   │   └── layout.css
│   └── js/
│       ├── api/           # API 클라이언트
│       ├── common/        # 공통 유틸리티
│       ├── components/    # Vue 컴포넌트
│       ├── pages/         # 페이지 컴포넌트
│       └── utils/         # 유틸리티 함수
├── .env.example           # 환경 설정 템플릿
├── .env.local             # 로컬 환경 설정
├── .env.development       # 개발 환경 설정
├── .env.production        # 운영 환경 설정
└── nodemon.json           # Nodemon 설정
```

## 🌍 환경별 설정

### Local Environment (.env.local)
- 로컬 개발 환경
- 디버그 모드 활성화
- 목 데이터 사용
- 상세 로깅

### Development Environment (.env.development)
- 개발 서버 환경
- API 문서 활성화
- 개발용 데이터베이스 연결
- 보통 수준의 로깅

### Production Environment (.env.production)
- 운영 서버 환경
- 보안 강화
- 최소 로깅
- 성능 최적화

## 🔧 주요 기능

### 헤더/푸터/사이드바 레이아웃


### API 통신
- Ajax 기반 통신
- 환경별 설정 지원
- 에러 핸들링

### 지도 기능
- OpenStreetMap 지원
- OpenLayers 라이브러리
- 마커 및 벡터 레이어 관리
- 다양한 지도 스타일 지원
- 좌표계 변환 및 투영법 지원

### 데이터 관리
- CRUD 기능
- 페이지네이션
- 필터링 및 검색

## 🛠 개발 도구

### Nodemon 설정


### 환경 변수 관리


## 📡 API 엔드포인트

- `GET /api/data` - 데이터 조회
- `POST /api/data` - 데이터 생성
- `GET /api/health` - 서버 상태 확인
- `GET /api/env` - 환경 정보 (디버그 모드 시)

## 🎯 사용 기술

- **Frontend**: React (CDN), OpenLayers
- **Backend**: Golang
- **Map**: OpenStreetMap
- **Tools**: GIN

## 🐳 Docker 배포

### Docker 빌드 및 실행

```bash
# Docker 이미지 빌드
docker build -t n .

# Docker 컨테이너 실행
docker run -d -p 3000:3000 --name n n
```

### Docker Compose 사용

```bash
# 운영 환경 실행
docker-compose up -d

# 개발 환경 실행 (소스 코드 변경 감지)
docker-compose -f docker-compose.dev.yml up -d

# 로그 확인
docker-compose logs -f

# 컨테이너 중지
docker-compose down
```

### Docker 이미지 관리

```bash
# 이미지 목록 확인
docker images

# 컨테이너 상태 확인
docker ps

# 컨테이너 로그 확인
docker logs name

# 컨테이너 내부 접속
docker exec -it image sh
```

### 환경 변수 설정

Docker 실행 시 환경 변수를 설정할 수 있습니다:

```bash
docker run -d \
  -p p:p \
  -e NODE_ENV=production \
  -e PORT= \
  --name  \
  
```

### 데이터 볼륨 마운트

데이터 지속성을 위해 볼륨을 마운트할 수 있습니다:

```bash
docker run -d \
  -p p:p \
  -v $(pwd)/data:/app/data \
  --name name \
  name
```


## 📝 라이센스

MIT License
