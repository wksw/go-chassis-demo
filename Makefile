update: .gitmodules ## 更新本地代码
	git submodule foreach --recursive git reset --hard 
	git submodule foreach --recursive git clean -fdx
	git submodule init
	git submodule update
	git submodule update --remote
	# git submodule foreach  --recursive 'tag="$$(git config -f $$toplevel/.gitmodules submodule.$$name.tag)";[ -n $$tag ] && git reset --hard  $$tag || echo "this module has no tag"'

build: main.go Dockerfile version ## 编译成docker镜像
	# 拉取基础编译镜像
	$(DOCKER_PULL) $(DOCKERHUBPULLHOST)/$(BUILDIMAGE) || true
	# 编译
	$(DOCKER_RUN) -w /go/src/$(GITREPO)/$(PROJECT)/$(SERVICENAME) \
		-v $$(pwd):/go/src/$(GITREPO)/$(PROJECT)/$(SERVICENAME) \
		$(DOCKERHUBPUSHHOST)/$(BUILDIMAGE) make linux
	# 打成docker镜像文件
	$(DOCKER_BUILD) --build-arg BINARY_FILE=$(SERVICE_NAME_LINUX) -t $(DOCKERHUBPUSHHOST)/$(NAME_SPACE)/$(SERVICENAME):$(VERSION) .
	# 推送到docker镜像仓库
	$(DOCKER_PUSH) $(DOCKERHUBPUSHHOST)/$(NAME_SPACE)/$(SERVICENAME):$(VERSION) || true


run: main.go ## 本机运行

release: ## 打包

linux: protos main.go ## 编译成linux环境下运行文件
	GOOS=linux $(GO_BUILD) -o ./bin/$(SERVICE_NAME_LINUX) .
mac: protos main.go ## 编译成mac环境下运行文件
	GOOS=darwin $(GO_BUILD) -o ./bin/$(SERVICE_NAME_MAC) .
windows: protos main.go ## 编译成windows环境下运行文件
	GOOS=windows $(GO_BUILD) -o ./bin/$(SERVICE_NAME_WIN) .

test: ## 运行测试用例

onebox: docker-compose.yaml version ## onebox环境部署
	export SERVICE_IMAGE=$(DOCKERHUBPULLHOST)/$(NAME_SPACE)/$(SERVICENAME):$(VERSION) && \
	docker-compose -p $(NAME_SPACE) up -d

onebox_down: docker-compose.yaml version ## onebox环境卸载
	docker-compose -p $(NAME_SPACE) down || true

protos: ## 协议文件编译
	make -C thirdparty/protos-repo build
 
onebox_restart: onebox_down onebox ## onebox环境重启

load_images: images ## 加载docker镜像

k8s: manifests version ## k8s环境部署

k8s_down: manifests version ## k8s环境卸载

k8s_upgrade: manifests version ## k8s环境升级

k8s_rollback: manifests version ## k8s环境回滚
