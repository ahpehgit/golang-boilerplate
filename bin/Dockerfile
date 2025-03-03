# Build executable first, CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../bin  in src folder
# To build image, run docker build -f bin/Dockerfile -t acrospprod.azurecr.io/<repo name:tag> . from project root folder

FROM public.ecr.aws/docker/library/alpine

ARG APP_NAME=golang-boilerplate

ENV OSPRESERVATION_PACKAGE=release
ENV OSPRESERVATION_HOST=0.0.0.0
ENV OSPRESERVATION_PORT=80
ENV APPNAME=${APP_NAME}

WORKDIR /home

# Get timezone data
RUN apk update && apk add bash tzdata
ENV TZ="Asia/Singapore"

COPY ./src/config.yml /home
COPY ./bin/${APP_NAME} /home

EXPOSE $OSPRESERVATION_PORT

ENTRYPOINT ["/bin/sh", "-c", "/home/$APPNAME"]