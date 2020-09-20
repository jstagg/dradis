# Specify the front end
FROM golang:1.15-alpine AS build

# with minideb
#FROM bitnami/minideb:buster AS build
#RUN install_packages golang git

WORKDIR /src/
COPY main.go go.* /src/
RUN CGO_ENABLED=0 go build -o /bin/dradis

FROM scratch
COPY --from=build /bin/dradis /bin/dradis
ENTRYPOINT ["/bin/dradis"]

# build hint
# docker build -t jstagg/repo:dradis-front .