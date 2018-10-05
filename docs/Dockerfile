FROM node:8.11.4

RUN mkdir /app
WORKDIR /app/website

COPY website/package.json .
RUN npm install

COPY . /app

CMD ["npm", "start"]
