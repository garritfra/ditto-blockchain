FROM node:10-alpine
ENV BACKEND_ADDRESS=http://217.182.0.182:42000/

RUN mkdir app

COPY . /app

WORKDIR /app

RUN npm install \
    && npm run build

EXPOSE 8080

FROM nginx:stable-alpine

COPY --from=0 /app/dist /usr/share/nginx/html
