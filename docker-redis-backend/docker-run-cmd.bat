@echo off

docker network create -d bridge my-bridge-network

docker run --name dradis-back -p 6379:6379 -d dradis-back:latest

docker network connect my-bridge-network dradis-back

rem -30-
