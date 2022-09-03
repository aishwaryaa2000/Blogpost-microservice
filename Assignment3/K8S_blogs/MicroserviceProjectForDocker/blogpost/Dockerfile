#docker file is used to create an image so that we can make a container

FROM node:18-alpine3.15  
#downloading image and creating a container

#RUN command inside the container
RUN mkdir -p /home/app/blogpost
#make a folder /home/app/blogpost inside container

COPY package.json /home/app/blogpost
#copies package.json from host to blogpost folder inside container

WORKDIR /home/app/blogpost
#go to the working directory to do npm install
RUN npm install

COPY  . /home/app/blogpost
#Copy whole folder contents into blogpost folder inside container

CMD [ "node","index.js" ]
#cmd will run commands after the container of the new image is created
