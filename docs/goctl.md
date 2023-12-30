替换模板文件
`goctl model mysql -src={patterns} -dir={dir} -cache --home ./goctl`

生成swagger
`goctl api plugin -plugin goctl-swagger="swagger -filename user.json -host 127.0.0.1:8888 -basepath /api " -api user.api -dir ../docs`

运行swagger ui， 例如在user目录下执行
`docker run --rm -p 8083:8080 -e SWAGGER_JSON=/docs/user.json -v $PWD/docs:/docs swaggerapi/swagger-ui`

在dsl中生成api
`goctl api go -api *.api -dir ../api  --style=go_zero --home ../../../goctl`

在rpc internal下生成model
`goctl model mysql datasource -url="root:@tcp(47.106.141.3:3306)/q_user" -table="role"  -dir="./model/role" -cache=true --style=go_zero --home /root/backend/qinglv-backend/goctl`

在rpc目录下生成rpc服务
`goctl rpc protoc order.proto --go_out=. --go-grpc_out=. --zrpc_out=. --style=go_zero`

请求rpc服务
`grpcurl -d '{"nickname": "keith"}' -plaintext 127.0.0.1:8080 user.User.Register`




docker run -p 6379:6379 --name redis -v /usr/src/software/redis/redis.conf:/etc/redis/redis.conf  -v /usr/src/software/redis/data:/data -d redis redis-server /etc/redis/redis.conf --appendonly yes