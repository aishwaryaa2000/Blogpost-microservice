#docker file is used to create an image so that we can make a container

FROM node:18-alpine3.15  
#downloading image and creating a container

#RUN command inside the container
RUN mkdir -p /home/app/frontend

COPY package.json /home/app/frontend
#copies package.json from host to frontend folder inside container

WORKDIR /home/app/frontend
#go to the working directory to do npm install

RUN npm install

COPY . /home/app/frontend

WORKDIR /home/app/frontend/src

# here we cannot use node index.js because inside package.json of frontend folder there is different scripts:
# whereas in other folders,the package.json only had index.js written
CMD [ "npm","start" ]
#cmd will run commands after the container of the new image is created
