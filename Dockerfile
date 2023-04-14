FROM golang:alpine AS build
RUN apk --no-cache add gcc g++ make git
WORKDIR /go/src/app

COPY go.mod .
RUN go mod download

COPY . .
# RUN go mod init app_name
RUN go mod tidy
RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/app ./cmd/app/main.go


FROM alpine:3.17
# RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin    
COPY --from=build /go/src/app/bin /go/bin
COPY --from=build /go/src/app/cmd/app/config.yaml /go/bin
COPY --from=build /go/src/app/.env /go/bin/.env
EXPOSE 8088
WORKDIR /go/bin    
ENTRYPOINT ./app