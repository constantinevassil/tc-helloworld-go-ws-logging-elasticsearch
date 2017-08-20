# tc-helloworld-go-ws-logging-elasticsearch

## Deploy 

```bash
sudo kubectl apply -f https://raw.githubusercontent.com/topconnector/tc-helloworld-go-ws-logging-elasticsearch/master/tc-helloworld-go-ws-logging-elasticsearch-deployment.yaml
```

## Create service

```bash
sudo kubectl apply -f https://raw.githubusercontent.com/topconnector/tc-helloworld-go-ws-logging-elasticsearch/master/tc-helloworld-go-ws-logging-elasticsearch-svc.yaml
```

## get IP address, port

```bash
[vagrant@tc-centos-master-hatc2 ~]$ kubectl get svc
NAME                                        CLUSTER-IP     EXTERNAL-IP   PORT(S)          AGE
kubernetes                                  10.96.0.1      <none>        443/TCP          2d
tc-helloworld-go-ws-logging-elasticsearch   10.103.40.46   <nodes>       1010:31031/TCP   5m
```

## Call service at: 10.103.40.46:1010

```bash
[vagrant@tc-centos-master-hatc2 ~]$ curl 10.103.40.46:1010
Hello World from Go in minimal Docker container (4.28MB) - tc-helloworld-go-ws-loggin
```

```bash
[vagrant@tc-centos-master-hatc2 ~]$ kubectl get svc
NAMESPACE     NAME                                        CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE
...
kube-system   elasticsearch-logging                       10.103.188.191   <nodes>       9200:32168/TCP   2d
```

## Calling Kibana from host
```bash
curl 192.168.0.199:31173
OR
inside VM
[vagrant@tc-centos-master-hatc2 ~]$ curl 10.98.145.221:5601
<script>var hashRoute = '/app/kibana';
var defaultRoute = '/app/kibana';

var hash = window.location.hash;
if (hash.length) {
  window.location = hashRoute + hash;
} else {
  window.location = defaultRoute;
}
```

1. Open 192.168.0.199:31173 in browser
1. Create index tc-helloworld-go-ws-logging-elasticsearch-log

## Explore logs in Kibana:

```bash
GET _search
{
  "query": {
    "match_all": {}
  }
}
```


```bash
{
    "_index": "tc-helloworld-go-ws-logging-elasticsearch-log",
    "_type": "log",
    "_id": "AV34C1u8-QzKVnwpofos",
    "_score": 1,
    "_source": {
      "Host": "localhost",
      "Timestamp": "2017-08-19T01:11:20.245549586Z",
      "Message": "tc-helloworld-go-ws-logging-elasticsearch: started, serving at 8080",
      "Data": {},
      "Level": "INFO"
    }
},
```



