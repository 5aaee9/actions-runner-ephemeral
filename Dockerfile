FROM golang:alpine AS builder

WORKDIR /app
COPY . /app/

ENV CGO_ENABLED=0

RUN go build -trimpath -ldflags="-s -w" -o actions-runner-ephemeral ./cmd/actions-runner-ephemeral

FROM scratch

COPY --from=builder /app/actions-runner-ephemeral /actions-runner-ephemeral

ENTRYPOINT [ "/actions-runner-ephemeral" ]