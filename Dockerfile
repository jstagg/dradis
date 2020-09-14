FROM golang:1.15-alpine AS build

WORKDIR /src/
COPY main.go go.* /src/
RUN CGO_ENABLED=0 go build -o /bin/dradis

FROM scratch
COPY --from=build /bin/dradis /bin/dradis
ENTRYPOINT ["/bin/dradis"]
