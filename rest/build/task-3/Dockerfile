FROM golang as build-env

ADD . /
RUN go mod download && go build -o /build/task-3/task3 /cmd/task-3/main.go


FROM gcr.io/distroless/base as task3-image

COPY --from=build-env build/task-3/task3 /build/task-3/task3
COPY --from=build-env tools/operations.log /tools/operations.log
COPY --from=build-env tools/responsetime.log /tools/responsetime.log

CMD ["/build/task-3/task3"]
EXPOSE 8080 