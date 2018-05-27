#!/bin/sh
docker build -t mock_gate .
docker run -it --rm -p 9000:8080 mock_gate
