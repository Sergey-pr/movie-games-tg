#first stage
FROM golang:alpine as builder
WORKDIR /MultiStage
RUN mkdir bin
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/main ./main.go


#second stage
FROM alpine:latest
WORKDIR /root/
ENV TZ UTC
COPY --from=builder /MultiStage/bin/main /root/
COPY --from=builder /card_files/ /root/card_files
CMD ["./main"]
