@echo off

rem docker network create -d bridge my-bridge-network

docker run --name dradis-front -p 8888:8888 -d --net my-bridge-network dradis-front:latest

docker network connect my-bridge-network dradis-front

rem -30-
