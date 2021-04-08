FROM golang:alpine as build

WORKDIR /src/
COPY . /src/
RUN CGO_ENABLED=0 go build -o /bin/messagingapp

FROM scratch
COPY --from=build /bin/messagingapp /bin/messagingapp
COPY --from=build /src/htmltemplates /bin/htmltemplates
COPY --from=build /src/resources /bin/resources
WORKDIR /bin/
EXPOSE 8080
ENTRYPOINT ["messagingapp"]