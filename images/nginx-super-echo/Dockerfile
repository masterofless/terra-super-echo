FROM node:17

WORKDIR /app
COPY package*.json /app
COPY src/*.js /app/src/

RUN npm install
RUN npx babel src --out-file terra-super-echo.js --presets=@babel/env,@babel/preset-react

FROM nginx:latest
WORKDIR /usr/share/nginx/html

COPY index.html /usr/share/nginx/html/index.html
COPY --from=0 /app/*.js /usr/share/nginx/html/
