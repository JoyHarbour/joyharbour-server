#
# server build
#

FROM golang:1.24-alpine AS build-server

# install packages

# de-escalate to build more safely
RUN adduser -D builder
RUN mkdir -p /app
RUN chown builder:builder /app
USER builder

# copy project meta & download dependencies
WORKDIR /app/src
COPY --chown=builder:builder server/go.mod server/go.sum ./
RUN go mod download

# build program
COPY --chown=builder:builder server .
RUN mkdir /app/bin
RUN go build -o /app/bin/joyharbour-server cmd/main.go

#
# ui build
#

FROM node:20-alpine AS build-ui

# install packages

# de-escalate to build more safely
RUN adduser -D builder
USER builder

# copy project meta & download dependencies
COPY --chown=builder:builder ui/package.json ui/package-lock.json /app/
WORKDIR /app
RUN npm install

# build frontend
COPY --chown=builder:builder ui /app
RUN npm run build

#
# final container
#

FROM alpine:3.21

# de-escalate
RUN adduser -D joyharbour
RUN mkdir /app /app/home
RUN chown -R joyharbour:joyharbour /app
USER joyharbour

# copy files from bits
COPY --from=build-server --chown=root:root /app/bin /app/bin
COPY --from=build-ui --chown=root:root /app/dist /app/home/www

EXPOSE 8000

ENV JOYHARBOUR_HOME=/app/home
CMD [ "/app/bin/joyharbour-server" ]
