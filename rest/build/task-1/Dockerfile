FROM golang as build-env

ADD . /
RUN go mod download && go build -o /build/task-1/task1 /cmd/task-1/main.go


FROM gcr.io/distroless/base as task1-image

COPY --from=build-env build/task-1/task1 /build/task-1/task1
COPY --from=build-env tools/operations.log /tools/operations.log
COPY --from=build-env tools/responsetime.log /tools/responsetime.log
COPY --from=build-env tools/data.csv /tools/data.csv

CMD ["/build/task-1/task1"]
EXPOSE 8080 