.PHONY: up-redis
up-redis: ## Redisをローカルで起動します
	sudo ifconfig lo0 alias 127.0.100.1
	sudo ifconfig lo0 alias 127.0.100.2
	sudo ifconfig lo0 alias 127.0.100.3
	docker compose -f docker-compose.yaml up -d

.PHONY: down-redis
down-redis: ## ローカルで起動しているRedisを停止します
	docker compose -f docker-compose.yaml down
	sudo ifconfig lo0 inet 127.0.100.1 delete
	sudo ifconfig lo0 inet 127.0.100.2 delete
	sudo ifconfig lo0 inet 127.0.100.3 delete
