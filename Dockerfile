# Nginx
#
# VERSION               0.0.1

#FROM      ubuntu
#MAINTAINER Guillaume J. Charmes <guillaume@dotcloud.com>

# make sure the package repository is up to date
#RUN echo "deb http://archive.ubuntu.com/ubuntu precise main universe" > /etc/apt/sources.list
#RUN apt-get update

#RUN apt-get install -y inotify-tools nginx apache2 openssh-server
FROM ubuntu

RUN apt-get update && apt-get install -y php5-cli php5-dev php-pear wget build-essential
RUN wget http://www.xmailserver.org/libxdiff-0.23.tar.gz && tar xvzf libxdiff-0.23.tar.gz && cd libxdiff-0.23 && ./configure && make && make install
RUN pecl install xdiff-1.5.2
RUN echo extension=xdiff.so >> /etc/php5/cli/conf.d/xdiff.ini
RUN php -r "readfile('https://getcomposer.org/installer');" | php -- --install-dir=/bin --filename=composer

WORKDIR /app
ENV PORT 5000
CMD composer install && php bin/server.php
