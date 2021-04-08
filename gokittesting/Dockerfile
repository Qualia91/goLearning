FROM golang:alpine as build

WORKDIR /src/
COPY . /src/
RUN CGO_ENABLED=0 go build -o /bin/gokittesting

FROM scratch
COPY --from=build /bin/gokittesting /bin/gokittesting
WORKDIR /bin/
ENTRYPOINT ["gokittesting"]