FROM ubuntu:18.04

RUN sed -i 's|http://archive.ubuntu.com/ubuntu/|https://mirror.kakao.com/ubuntu/|g' /etc/apt/sources.list && \
    sed -i 's|http://security.ubuntu.com/ubuntu/|https://mirror.kakao.com/ubuntu/|g' /etc/apt/sources.list

RUN echo 'Acquire::https::mirror.kakao.com::Verify-Peer "false"; \
    Acquire::https::mirror.kakao.com::Verify-Host "false";' > /etc/apt/apt.conf.d/99insecure
    
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates

RUN rm -f /etc/apt/apt.conf.d/99insecure

RUN apt-get update && apt-get -y install \
    libbz2-dev libaio-dev bison zlib1g-dev \
    libsnappy-dev libgflags-dev libreadline6-dev \
    libncurses5-dev libssl-dev liblz4-dev \
    libboost-dev gdb git libzstd-dev libcap-dev \
    build-essential cmake make g++-8 gcc-8 \
    && apt-get clean

WORKDIR /usr/local/src

RUN git clone https://github.com/facebook/zstd.git && \
    cd zstd && \
    make -j$(nproc) && \
    make install && \
    cd ..

COPY mysql-5.6.35 /usr/local/src/mysql-5.6

RUN mkdir -p /usr/local/src/mysql-5.6/build
WORKDIR /usr/local/src/mysql-5.6/build

RUN git submodule init && git submodule update

RUN cmake ../ \
    -DCMAKE_BUILD_TYPE=RelWithDebInfo \
    -DWITH_SSL=system \
    -DWITH_ZLIB=bundled \
    -DMYSQL_MAINTAINER_MODE=0 \
    -DENABLED_LOCAL_INFILE=1 \
    -DENABLE_DTRACE=0 \
    -DENABLE_DOWNLOADS=1 \
    -DCMAKE_C_COMPILER=gcc-8 \
    -DCMAKE_CXX_COMPILER=g++-8 \
    -DCMAKE_INSTALL_PREFIX=/usr/local/mysql \
    -DCMAKE_CXX_FLAGS="-march=native -std=c++17" \
    -DWITH_ZSTD=system

RUN make -j$(nproc)\
    && make install

RUN groupadd mysql && useradd -r -g mysql mysql

WORKDIR /usr/local/mysql
RUN chown -R mysql:mysql . && \
    scripts/mysql_install_db --user=mysql && \
    chown -R root . && \
    chown -R mysql data

RUN cp support-files/my-default.cnf /etc/my.cnf
RUN echo "[mysqld]\n\
    rocksdb\n\
    default-storage-engine=rocksdb\n\
    skip-innodb\n\
    default-tmp-storage-engine=MyISAM\n\
    collation-server=latin1_bin\n\
    log-bin\n\
    binlog-format=ROW" >> /etc/my.cnf

RUN echo "PATH=$PATH:$HOME/bin:/usr/local/mysql/bin" >> ~/.bashrc && \
    source ~/.bashrc

RUN chmod 755 support-files/mysql.server && \
    cp support-files/mysql.server /etc/init.d/mysqld && \
    systemctl unmask mysqld.service && \
    systemctl enable mysqld.service

EXPOSE 3306

CMD ["mysqld", "--user=mysql"]
