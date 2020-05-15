通过kubectl，执行下面的命令在Kubernetes集群中部署Oracle数据库。
```
$ kubectl create -f rabbitmq.yaml --namespace=kube-public
```

在部署完成后，通过下面的命令可以查看RabbitMQ暴露的端口：
```
$ kubectl get svc --namespace=kube-public
```