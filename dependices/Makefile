include ../thirdparty/maintain/Makefile

dependices: docker-compose.yaml ## 安装依赖
	docker-compose -p $(NAME_SPACE) up -d

dependices_down: docker-compose.yaml ## 依赖卸载
	docker-compose -p $(NAME_SPACE) down || true