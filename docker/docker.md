# Run gocotea with docker
Images:
- gocotea_container/ - image with installed gocotea
- ansible_host_container/ - image acting as an Ansible host

Build both images and enter theire names into docker-compose.yaml file. After this make:
```bash
docker compose up -d
```

After that enter into gocotea container (make docker ps to obtain gocotea container ID):
```bash
docker exec -it GOCOTEA_CONTAINER_ID /bin/bash
```

In the container make:
```bash
cd /home/ubuntu/gocotea_run
go mod tidy
go run main.go
```

After this the gocotea and gopython packages will be downloaded and the program will be launched.

