FROM golang:1.23.0-alpine AS build

WORKDIR /build

COPY user_backend/go.mod user_backend/go.sum ./

COPY pkg/ ../pkg/

RUN go mod download

COPY user_backend/ ./

RUN apk add --no-cache make

RUN go build -o app cmd/app/main.go

FROM alpine AS runner

WORKDIR app

# RUN apk add --no-cache curl

COPY --from=build /build/app ./app
COPY --from=build /build/cmd/app/config/config.yml ./config.yml

CMD ["/app/app", "--config=/app/config.yml"]
