FROM node:carbon
RUN npm install -g webpack
WORKDIR /usr/src/app
COPY package*.json ./
RUN npm install
COPY . .
RUN webpack
EXPOSE 3000
CMD [ "npm", "start" ]