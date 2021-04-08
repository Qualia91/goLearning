FROM golang:1.14-alpine AS build

# this is a multistage build. It uses the first block to build the go app in one container, 
# then puiblished it to another container with just the runtime in. This is done because
# the build container needs the golang base image to build, whereas the released go binary
# is small and selc contained, meaning the relase container can be allot smaller.
WORKDIR /src/
COPY main.go go.* /src/
RUN CGO_ENABLED=0 go build -o /bin/simplegodocker

FROM scratch
COPY --from=build /bin/simplegodocker /bin/simplegodocker
ENTRYPOINT ["/bin/simplegodocker"]