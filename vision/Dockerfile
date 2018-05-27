FROM openalpr/openalpr

# Install pip
RUN apt-get update && apt-get install -y \
    python3-pip

# Build python bindings
WORKDIR /srv/openalpr/src/bindings/python
RUN python3 /srv/openalpr/src/bindings/python/setup.py install

# Install pip packages
COPY requirements.txt /srv/openalpr/proj/
WORKDIR /srv/openalpr/proj
RUN pip3 install -r requirements.txt

# Copy project files
COPY . /srv/openalpr/proj/

# Run the terminal
ENTRYPOINT ["/bin/sh","-c"]
CMD ["python3 main.py"]