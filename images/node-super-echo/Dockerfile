FROM node:16
ENV NODE_ENV=production

WORKDIR /app
EXPOSE 8080

COPY ["package.json", "package-lock.json*", "./"]

RUN npm install --production

COPY . .

CMD [ "node", "server.js" ]
