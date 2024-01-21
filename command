// docker build
docker build -t sample_project:local .

//infra up
docker-compose -f infra/docker-compose.yaml up -d

keploy record -c "docker-compose up" --containerName "sample_project_app"

keploy test -c "docker-compose up" --containerName "sample_project_app" --delay 10
