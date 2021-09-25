start:
	foreman start

up: SERVICES=icecast postgres liquidsoap
up:
	docker-compose --env-file .env.dev build ${SERVICES}
	docker-compose --env-file .env.dev up -d ${SERVICES}

deploy: export DOCKER_HOST=ssh://ubuntu@stream.djfullmoon.com
deploy:
	echo ${DOCKER_HOST}
	docker-compose --env-file .env.prod build
	docker-compose --env-file .env.prod up -d icecast liquidsoap
	docker-compose --env-file .env.prod up -d app
	docker-compose --env-file .env.prod up -d worker

psql:
	psql postgres://djfm:djfm@localhost:5432/djfullmoon_development
