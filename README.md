# gaia-examples
Gaia Examples

## consul 启动
```shell
consul agent -server -bootstrap -advertise 127.0.0.1 -data-dir ./data -ui
```


## nacos 启动
```shell
sh startup.sh -m standalone
```
## nacos 关闭
```shell
sh shutdown.sh
```