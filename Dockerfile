FROM       scratch
MAINTAINER Constantine Vassil <constantine@mobiledatabooks.com>
#
ENV ELASTICSEARCH_URL elasticsearch_url
ENV ELASTICSEARCH_LOG_INDEX_NAME name
ENV NAMESPACE namespace
ENV NODE_NAME node_name

#
ADD        tc-helloworld-go-ws-logging-elasticsearch tc-helloworld-go-ws-logging-elasticsearch
EXPOSE     1010
ENTRYPOINT ["/tc-helloworld-go-ws-logging-elasticsearch"]
