# Builder layer.
FROM node:13.2.0 as build
WORKDIR /app
COPY . ./
RUN npm i
RUN npm run-script build

# Runner layer.
FROM nginx:alpine
COPY --from=build /app/build /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
