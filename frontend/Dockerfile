FROM node:latest as build-stage
WORKDIR /app
COPY package*.json ./
RUN node --max-old-space-size=250 `which npm` install
COPY ./ .
RUN node --max-old-space-size=250 `which npm` run build

FROM nginx as production-stage
RUN mkdir /app
COPY --from=build-stage /app/dist /app
COPY nginx.conf /etc/nginx/nginx.conf
CMD ["nginx", "-g", "daemon off;"]