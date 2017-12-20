FROM node AS frontend-build

COPY ./web /app

RUN cd /app \
    && npm install \
    && npm run build


FROM golang:1.8 AS server-build

COPY ./server /go/src/app
COPY --from=frontend-build /app/build /go/src/app/build

RUN cd /go/src/app/ \
    && go get \
    && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o robolearn .

FROM alpine

WORKDIR /root/
COPY --from=server-build /go/src/app/robolearn .

CMD ["./robolearn"]

EXPOSE 9000
