# syntax=docker/dockerfile:1
# Etap 1: Obraz bazowy Scratch z aplikacją Go
FROM alpine AS builder

# Wersja aplikacji GO może zostać zmieniona w argumencie
ARG APP_VERSION=1.22.2

# Instalacja potrzebnych pakietów w tym klienta ssh i git
RUN apk add --no-cache openssh-client git url tar gcc musl-dev

# Pobranie klucza publicznego z github.com
RUN mkdir -p -m 0600 ~/.ssh && ssh-keyscan github.com >> ~/.ssh/known_hosts

# Pobranie i instalacja Go
RUN curl -L https://go.dev/dl/go${APP_VERSION}.linux-amd64.tar.gz -o go.tar.gz \
    && tar -C /usr/local -xzf go.tar.gz \
    && rm go.tar.gz

# Ścieżka do Go
ENV PATH="/usr/local/go/bin:${PATH}"

WORKDIR /app

# Klonowanie repozytorium przez ssh (w nim zawarty jest plik main.go)
RUN --mount=type=ssh git clone git@github.com:Extremewars/PAwChO-lab6.git .

# Kompilacja aplikacji
RUN go build -ldflags "-X main.version=${APP_VERSION}" -o main main.go

# Etap 2: Obraz końcowy nginx
FROM nginx:alpine AS final

WORKDIR /app

RUN apk add --no-cache curl

COPY --from=builder /app/main .

COPY nginx.conf /etc/nginx/nginx.conf

EXPOSE 80

HEALTHCHECK --interval=10s --timeout=3s --start-period=5s \
  CMD curl -f http://localhost/ || exit 1

CMD ["/bin/sh", "-c", "/app/main & sleep 2 && nginx -g 'daemon off;'"]
