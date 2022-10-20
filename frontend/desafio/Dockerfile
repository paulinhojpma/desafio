FROM node:latest
WORKDIR /frontend/desafio/

COPY . ./
EXPOSE 3000
# building the app
RUN npm i
RUN npm run build


# Running the app
CMD [ "npm", "start"]