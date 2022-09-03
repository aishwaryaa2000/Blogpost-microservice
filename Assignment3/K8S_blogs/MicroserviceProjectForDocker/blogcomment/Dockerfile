#docker file is used to create an image so that we can make a container

FROM node:18-alpine3.15  
#downloading image and creating a container

#RUN command inside the container
RUN mkdir -p /home/app/blogcomment
#make a folder /home/app/blogcomment inside container

COPY package.json /home/app/blogcomment
#copies package.json from host to blogcomment folder inside container

WORKDIR /home/app/blogcomment
#go to the working directory to do npm install
RUN npm install

COPY . /home/app/blogcomment
#Copy whole folder contents into blogcomment folder inside container

 CMD [ "node","index.js" ]
#cmd will run commands after the container of the new image is created
