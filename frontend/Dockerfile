FROM node:10-alpine

ENV BACKEND_ADDRESS $BACKEND_ADDRESS

RUN mkdir app

COPY . /app

WORKDIR /app

RUN npm install \
    && npm run build

EXPOSE 8080

FROM nginx:stable-alpine

COPY --from=0 /app/dist /usr/share/nginx/html
