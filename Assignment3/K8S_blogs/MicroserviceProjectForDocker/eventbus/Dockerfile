#docker file is used to create an image so that we can make a container

FROM node:18-alpine3.15  
#downloading image and creating a container

#RUN command inside the container
RUN mkdir -p /home/app/eventbus
# creats a eventbus folder inside container

COPY package.json /home/app/eventbus
#copies package.json from host to eventbus folder inside container

WORKDIR /home/app/eventbus
#go to the working directory to do npm install

RUN npm install

COPY . /home/app/eventbus

CMD [ "node","index.js" ]
#cmd will run commands after the container of the new image is created
