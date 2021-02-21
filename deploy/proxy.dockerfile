FROM nginx:1.19-alpine
COPY ./proxy/nginx-dev.conf /etc/nginx/conf.d/default.conf
# COPY ./nginx-prd.conf /etc/nginx/conf.d/default.conf