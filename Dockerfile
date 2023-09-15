FROM golang:1.21 as builder
WORKDIR /app/
COPY *.go /app/
RUN go build -o /app/out/main *.go

FROM scratch
COPY --from=builder /app/out/main /app/main
CMD ["/app/main"]