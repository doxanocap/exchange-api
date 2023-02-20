
docker build -t app-server .

cd ../client

docker build -t app-client

docker run --name some-postgres -e POSTGRES_PASSWORD=eldoseldos -d --rm postgres -v=.init.sql:/docker-entrypoint-initdb.d/init.sql
 
 docker run --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword -d postgres
