FROM node:10-alpine as node

WORKDIR /usr/src/app

COPY package*.json ./

RUN npm install

ADD . .

RUN npm run build

FROM abiosoft/caddy

COPY --from=node /usr/src/app/dist/go-wiki app/
COPY ./Caddyfile /etc/Caddyfile
WORKDIR app

EXPOSE 80 443 2015
CMD ["--conf", "/etc/Caddyfile", "--log", "stdout"]