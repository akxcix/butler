# this is not the right way, but its a temp solution
docker build -t butler:latest .
docker run -d -p 8080:8080 -v $(pwd)/config.yml:/app/config.yml butler:latest
