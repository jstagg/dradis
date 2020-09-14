echo off

rem docker network create -d bridge my-bridge-network

docker run --name dradis-fe -p 8888:8888 -d --net my-bridge-network jstagg/dradis-fe:latest

rem docker network connect my-bridge-network dradis-fe

rem -30-
