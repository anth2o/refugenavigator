FROM golang:1.24-alpine AS backend-builder
WORKDIR /app
COPY backend-go/go.mod backend-go/go.sum ./
RUN go mod download && \
    go mod verify
COPY backend-go/cmd/server/ ./cmd/server/
COPY backend-go/internal ./internal
RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/server/main.go

FROM node:22-alpine AS frontend-builder
WORKDIR /app
COPY frontend/package.json frontend/yarn.lock ./
RUN yarn install
COPY frontend/ ./
RUN yarn run build

FROM alpine/git:latest AS git-tag
WORKDIR /app
COPY .git ./.git
RUN git describe --tags --always > .git-tag

FROM alpine:3.22
COPY --from=backend-builder /app/main /app/backend-go/main
COPY --from=frontend-builder /app/dist /app/frontend/dist
COPY --from=git-tag /app/.git-tag /app/.git-tag
WORKDIR /app/backend-go
ENV GIN_MODE=release
ENV PORT=8080
CMD ["./main"]