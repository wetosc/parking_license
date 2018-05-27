#!/bin/sh
docker build -t vision .
docker run -it --net mynet --name vision --rm vision