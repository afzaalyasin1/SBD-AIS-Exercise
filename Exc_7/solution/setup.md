# note commands:

# Commands Used to Setup Docker Swarm Cluster

# 1. Initialize Swarm
docker swarm init

# 2. Create Secrets
docker secret create postgres_user docker/postgres_user_secret
docker secret create postgres_password docker/postgres_password_secret
docker secret create s3_user docker/s3_user_secret
docker secret create s3_password docker/s3_password_secret

# 3. Deploy Stack
docker stack deploy -c docker-compose.yml mysystem

# 4. Check Services
docker service ls
docker service ps mysystem_orderservice
docker service ps mysystem_postgres
docker service ps mysystem_minio