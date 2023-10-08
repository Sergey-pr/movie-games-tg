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
RUN mkdir /root/card_files
RUN mkdir /root/bot_files
COPY --from=builder /MultiStage/card_files/ /root/card_files/
COPY --from=builder /MultiStage/bot_files/ /root/bot_files/
CMD ["./main"]
