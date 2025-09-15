FROM golang:1.23.0-alpine AS build

WORKDIR /build

COPY image_processor/go.mod image_processor/go.sum ./

COPY pkg/ ../pkg/

COPY user_backend/go.mod ../user_backend/go.mod
COPY user_backend/domain ../user_backend/domain

RUN go mod download

COPY image_processor/ ./

RUN apk add --no-cache make

RUN go build -o app cmd/app/main.go

FROM alpine AS runner

WORKDIR app

# RUN apk add --no-cache curl

COPY --from=build /build/app ./app
COPY --from=build /build/cmd/app/config/config.yml ./config.yml

CMD ["/app/app", "--config=/app/config.yml"]
