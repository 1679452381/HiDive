# 生成api业务代码 ， 进入"服务/cmd/api/desc"目录下，执行下面命令
# goctl api go -api *.api -dir ../  --style=goZero

# 生成rpc业务代码
#服务/cmd/rpc/pb"目录下，执行下面命令
#goctl rpc protoc *.proto --go_out=../ --go-grpc_out=../  --zrpc_out=../ --style=goZero