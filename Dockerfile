FROM debian:bookworm-slim AS downloader

RUN apt-get update && \
    apt-get install -y curl

WORKDIR /models
RUN curl -L -O https://huggingface.co/CompVis/stable-diffusion-v-1-4-original/resolve/main/sd-v1-4.ckpt

FROM golang:1.24-bookworm AS builder

RUN apt-get update && \
    apt-get install -y cmake git

WORKDIR /src
COPY . .

RUN cd stable-diffusion.cpp && rm -rf build

RUN go generate
RUN go build main.go

FROM debian:bookworm-slim

WORKDIR /app
COPY --from=builder /src/main .
COPY --from=downloader /models/sd-v1-4.ckpt /app/models/sd-v1-4.ckpt

CMD ["/app/main"]