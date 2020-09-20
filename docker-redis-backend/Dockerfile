# Specify the node base image
#FROM redis:6
FROM bitnami/minideb:buster

# Maintainer
LABEL maintainer="jstagg@gmail.com"

# Explicitly expose the port(s)
EXPOSE 6379

# Ready the container
#RUN apt update && apt install -y dos2unix vim && apt clean
RUN install_packages redis dos2unix
RUN mkdir -p /data

# Install app dependencies
COPY data.* /data/
COPY load.sh /data/
RUN dos2unix /data/* && chmod 755 /data/load.sh && dos2unix /data/load.sh

# Run the process and the loader
CMD ["/data/load.sh"]

# build hint
# docker build -t jstagg/dradis-back:minideb .
# docker build -t dradis-back .