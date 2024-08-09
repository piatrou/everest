FROM merik/pnpm:8 as ui

WORKDIR /everest
COPY . .

WORKDIR /everest/ui

RUN pnpm install && \
  EVEREST_OUT_DIR=/everest/public/dist/ pnpm build


FROM golang:1.22-alpine as build

WORKDIR /everest

COPY . .
COPY --from=ui /everest/public /everest/public

RUN apk add -U --no-cache ca-certificates make
RUN make build

FROM scratch

WORKDIR /

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /everest/bin/everest /everest-api

EXPOSE 8080

ENTRYPOINT ["/everest-api"]

