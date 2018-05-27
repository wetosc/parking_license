docker build -t ocv .
docker run -it --privileged -P --rm --device=/dev/media0:/dev/media0 ocv