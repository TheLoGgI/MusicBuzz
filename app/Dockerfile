FROM node:22-alpine AS build

WORKDIR /app

COPY *.json ./
COPY *.ts ./
COPY *.js ./
COPY index.html ./
COPY /src ./src
COPY /public ./public

RUN yarn install
RUN yarn build

FROM nginx:1.27.1


WORKDIR /etc/nginx

RUN groupadd -r nginx-group && useradd -r -g nginx-group nginx-user
COPY --chown=nginx-group:nginx-user nginx.conf ./nginx.conf
COPY --chown=nginx-group:nginx-user --from=build /app/dist ./html

# Ensure the directory exists and change ownership
RUN mkdir -p /var/cache/nginx && chown -R nginx-user:nginx-group /var/cache/nginx/
RUN mkdir -p /var/run && chown -R nginx-user:nginx-group /var/run/

# Set permissions to allow nginx-user to open files in /var/run
RUN chmod -R 755 /var/run/

USER nginx-user

EXPOSE 8080