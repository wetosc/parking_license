#!/bin/sh
docker build -t mock_gate .
docker run -it --rm --net mynet --name mockgate -p 9000:8080 mock_gate
