FROM       scratch
MAINTAINER Constantine Vassil <constantine@mobiledatabooks.com>
#
ENV ELASTICSEARCH_URL elasticsearch_url
ENV NAMESPACE namespace
ENV NODE_NAME node_name
#
ADD        tc-helloworld-go-ws-logging-elasticsearch tc-helloworld-go-ws-logging-elasticsearch
EXPOSE     8080
ENTRYPOINT ["/tc-helloworld-go-ws-logging-elasticsearch"]
