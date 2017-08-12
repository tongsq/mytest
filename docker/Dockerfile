FROM centos:latest
ADD php-7.1.8.tar.gz /src/
RUN cd /src/php-7.1.8 \
&& yum -y install gcc \
&& yum -y install libxml2-devel \
&& yum -y install openssl openssl-devel libcurl-devel \
&& ./configure --prefix=/usr/local/php7 --with-pdo-mysql --with-curl --enable-bcmath --with-openssl --enable-mbstring --enable-opcache \
&& make && make install \
&& cp /usr/local/php7/bin/php /usr/local/bin/
