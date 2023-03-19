# syntax=docker/dockerfile:1

# 1. Dev
FROM node:19-alpine3.16 AS dev

WORKDIR /app

COPY package.json /app/package.json
COPY package-lock.json /app/package-lock.json

RUN npm ci

COPY . /app

ENV CI=true
ENV PORT=80

CMD [ "npm", "start" ]

# 2. Build
FROM dev AS build

ENV NODE_ENV=production

RUN npm run build

# 3. Prod
FROM nginx:1.23.3-alpine

COPY --from=build /app/nginx.conf /etc/nginx/conf.d/default.conf

WORKDIR /usr/share/nginx/html

RUN rm -rf ./*

COPY --from=build /app/build .

ENTRYPOINT ["nginx", "-g", "daemon off;"]
