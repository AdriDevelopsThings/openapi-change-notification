FROM node:16-alpine as frontend-base
WORKDIR /app

# install pnpm
RUN yarn global add pnpm

FROM frontend-base as frontend-dependencies

COPY ./frontend/package.json ./frontend/pnpm-lock.yaml ./
RUN pnpm install

FROM frontend-dependencies as frontend-build

COPY ./frontend ./

RUN pnpm next build
RUN pnpm next export

FROM golang:alpine as backend-dependencies
WORKDIR /build

RUN apk --no-cache add build-base ca-certificates

COPY ./go.mod ./go.sum ./

RUN go mod download

FROM backend-dependencies as backend-build

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o main .

FROM busybox
WORKDIR /dist

ENV ENVIRONMENT=production
ENV REDIS_URL=
ENV SQLITE_PATH=/db.sql
ENV FRONTEND_STATIC_SERVE=static
ENV FRONTEND_URL=
ENV FRONTEND_EMAIL_VERIFICATION_PATH=
ENV FRONTEND_UNSUBSCRIBE_PATH=
ENV FRONTEND_UNSUBSCRIBE_VERIFICATION_PATH=
ENV SMTP_SERVER=
ENV SMTP_USERNAME=
ENV SMTP_PASSWORD=
ENV SMTP_FROM_ADDRESS=

COPY --from=backend-build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=backend-build /build/main /dist/main
COPY --from=frontend-build /app/out /dist/static


ENV LISTEN_HOST=0.0.0.0
ENV LISTEN_PORT=80
EXPOSE 80
VOLUME [ "/db.sql" ]
CMD ["/dist/main"]

