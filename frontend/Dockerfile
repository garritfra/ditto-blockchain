FROM node:10-alpine

RUN mkdir app

COPY . /app

WORKDIR /app

RUN npm install \
    && npm run build

EXPOSE 8080

FROM nginx:stable-alpine

COPY dist /usr/share/nginx/html