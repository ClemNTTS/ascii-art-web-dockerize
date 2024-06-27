#!/bin/bash
docker images

echo "////////////////////////////////////////////"

sudo docker run -d --rm -p 3000:3000 ascii
sudo docker ps -a
xdg-open http://localhost:3000/

