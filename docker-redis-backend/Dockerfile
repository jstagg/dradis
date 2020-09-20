# Specify the node base image
#FROM redis:6
FROM bitnami/minideb:buster

# Maintainer
LABEL maintainer="jstagg@gmail.com"

# Explicitly expose the port(s)
EXPOSE 6379

# Ready the container
#RUN apt update && apt install -y dos2unix vim && apt clean
RUN install_packages redis dos2unix vim
RUN mkdir -p /data
RUN dos2unix /data/* && chmod 755 /data/load.sh && dos2unix /data/load.sh

# Install app dependencies
COPY data.* /data/
COPY load.sh /data/

# Run the process and the loader
CMD ["/data/load.sh"]
