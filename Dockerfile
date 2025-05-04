# Go 빌드용 이미지
FROM golang:1.24-alpine AS builder

# 작업 디렉토리 설정
WORKDIR /app

# Go 모듈 복사 및 다운로드
COPY go.mod go.sum ./
RUN go mod download

# 소스 복사
COPY . .

# Go 프로그램 빌드
RUN go build -o main .

# 실행용 이미지
FROM alpine:latest

# 인증서 루트 설치 (Gin이 HTTPS 요청 등을 처리할 때 필요)
RUN apk --no-cache add ca-certificates

# 작업 디렉토리 설정
WORKDIR /root/

# 빌드한 바이너리 복사
COPY --from=builder /app/main .

# 포트 오픈
EXPOSE 8080

# 실행 명령
CMD ["./main"]
