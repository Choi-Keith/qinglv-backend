替换模板文件
`goctl model mysql -src={patterns} -dir={dir} -cache --home ./goctl`                                                  

生成swagger
`goctl api plugin -plugin goctl-swagger="swagger -filename user.json -host 127.0.0.1:8888 -basepath /api " -api user.api -dir ../docs`

运行swagger ui， 例如在user目录下执行
`docker run --rm -p 8083:8080 -e SWAGGER_JSON=/docs/user.json -v $PWD/docs:/docs swaggerapi/swagger-ui`

在dsl中生成api
`goctl api go -api *.api -dir ../api  --style=go_zero --home ../../../goctl`

在dsl中生成ts接口
`goctl api ts --api *.api --dir=../docs/`

在rpc internal下生成model
`goctl model mysql datasource -url="root:xxx@tcp(43.139.228.81:3306)/q_user" -table="user"  -dir="./model/user" -cache=true --style=go_zero --strict=true --home /home/ubuntu/backend/qinglv-backend/goctl/1.5.6`

在docker运行redis
`sudo docker run --restart=always -d -p 16379:6379 -v /home/ubuntu/redis/data:/data:rw  --name redis redis:latest --requirepass xxx`

在rpc目录下生成rpc服务
`goctl rpc protoc ./pb/user.proto --go_out=. --go-grpc_out=. --zrpc_out=. --style=go_zero -m`

请求rpc服务
`grpcurl -d '{"nickname": "keith"}' -plaintext 127.0.0.1:8080 user.User.Register`

   