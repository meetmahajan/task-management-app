# Stage 1: Build
FROM node:14-alpine as build

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .

RUN npm run build --prod

# Stage 2: Serve
FROM nginx:alpine

COPY --from=build /app/dist/task-management-app /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]