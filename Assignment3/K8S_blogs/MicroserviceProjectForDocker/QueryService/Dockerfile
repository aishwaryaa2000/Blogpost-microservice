#docker file is used to create an image so that we can make a container

FROM node:18-alpine3.15  
#downloading image and creating a container

#RUN command inside the container
RUN mkdir -p /home/app/QueryService
# ./Assignment2/MicroserviceProjectForDocker/blogpost/package.json
COPY package.json /home/app/QueryService
#copies files from host machine inside your container
#copies package.json from host to app folder inside container

WORKDIR /home/app/QueryService
#go to the working directory to do npm install

RUN npm install

COPY . /home/app/QueryService

CMD [ "node","index.js" ]
#cmd will run commands after the container of the new image is created
