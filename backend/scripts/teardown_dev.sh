set -e

docker compose --env-file .env -f services/docker-compose.yml down --rmi all