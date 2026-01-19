# 创建项目目录
mkdir go-grpc-service && cd go-grpc-service

# 初始化 Go 模块
go mod init king/go-grpc-service

# 创建项目结构
mkdir -p api/proto/v1 internal/{server,client,service,repository} cmd/{server,client} config pkg/{middleware,utils}

# 安装grpc核心依赖
go get google.golang.org/grpc
go get google.golang.org/protobuf

# Go 插件
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest