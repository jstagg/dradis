echo off

docker network create -d bridge my-bridge-network

docker run --name dradis -p 6379:6379 -d jstagg/dradis:1.0

docker network connect my-bridge-network dradis

rem -30-
