FROM golang as build-env

ADD . /
RUN go mod download && go build -o /build/task-2/task2 /cmd/task-2/main.go


FROM gcr.io/distroless/base as task2-image

COPY --from=build-env build/task-2/task2 /build/task-2/task2
COPY --from=build-env tools/operations.log /tools/operations.log
COPY --from=build-env tools/responsetime.log /tools/responsetime.log

CMD ["/build/task-2/task2"]
EXPOSE 8080 