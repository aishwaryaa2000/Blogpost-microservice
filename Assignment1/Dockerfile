#Docker file is used to create an image so that we can make a container

FROM ubuntu/apache2:latest 
#downloading base image ubuntu/apache2 and creating a container.latest is the tag here

#RUN command inside the container
RUN service apache2 start
#Running the command to start apache inside container

COPY ./mypage.html  /var/www/html/
# Copying mypage.html from local machine to var/www/html inside the container


#NOTE -> We can also take the image - FROM httpd:2.4 -The Apache HTTP Server Project

