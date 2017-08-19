// Copyright 2017 Mobile Data Books, LLC. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// dep init
import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"gopkg.in/olivere/elastic.v5"
	"gopkg.in/sohlich/elogrus.v2"
	"net/http"
	"os"
	"time"
)

type viewHandler_helloHandler struct {
	// Logger is the log.Logger instance used to log messages with the Logger middleware
	Logger *logrus.Logger
}

func (v *viewHandler_helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t0 := time.Now()
	t1 := time.Now()

	t1 = time.Now()
	t01 := t1.Sub(t0)
	tStr1 := fmt.Sprintf("Hello World from Go in minimal Docker container (4.28MB) - tc-helloworld-go-ws-logging-elasticsearch - v.1.0, it took %v to run", t01)
	//fmt.Printf("\n%s\n", tStr1)
	v.Logger.Info(tStr1)

	fmt.Fprintln(w, tStr1)

}

// http://elasticsearch-logging:9200
func main() {

	envNAMESPACE := os.Getenv("NAMESPACE")
	envNODE_NAME := os.Getenv("NODE_NAME")
	envELASTICSEARCH_URL := os.Getenv("ELASTICSEARCH_URL")

	log := logrus.New()
	log.Println("NAMESPACE:" + envNAMESPACE)
	log.Println("NODE_NAME:" + envNODE_NAME)
	log.Println("elastic.SetURL:" + envELASTICSEARCH_URL)

	elastic.SetSniff(false)
	client, err := elastic.NewClient(elastic.SetURL(envELASTICSEARCH_URL)) // http://10.96.185.105:9200 10.0.1.73:30138  http://10.32.0.8:9200
	if err != nil {
		log.Panic(err)
	}
	hook, err := elogrus.NewElasticHook(client, "localhost", logrus.DebugLevel, "tc-helloworld-go-ws-logging-elasticsearch-log")
	if err != nil {
		log.Panic(err)
	}
	log.Hooks.Add(hook)

	http.Handle("/", &viewHandler_helloHandler{
		Logger: log,
	})
	log.Println("tc-helloworld-go-ws-logging-elasticsearch: started, serving at 8080")

	//fmt.Println("Started, serving at 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

// curl 'http://10.0.1.73:30138/_nodes/http?pretty'

// [vagrant@tc-centos-master-hatc2 ~]$ curl 'http://10.96.185.105:9200/_nodes/http?pretty'
// {
//   "_nodes" : {
//     "total" : 2,
//     "successful" : 2,
//     "failed" : 0
//   },
//   "cluster_name" : "kubernetes-logging",
//   "nodes" : {
//     "f4fAn_-cSfK39ljvXv2ZZg" : {
//       "name" : "elasticsearch-logging-1",
//       "transport_address" : "10.32.0.9:9300",
//       "host" : "10.32.0.9",
//       "ip" : "10.32.0.9",
//       "version" : "5.5.1",
//       "build_hash" : "19c13d0",
//       "roles" : [
//         "master",
//         "data",
//         "ingest"
//       ],
//       "attributes" : {
//         "ml.enabled" : "true"
//       },
//       "http" : {
//         "bound_address" : [
//           "[::]:9200"
//         ],
//         "publish_address" : "10.32.0.9:9200",
//         "max_content_length_in_bytes" : 104857600
//       }
//     },
//     "KE7r0yjSSKu__z8t9IgGEg" : {
//       "name" : "elasticsearch-logging-0",
//       "transport_address" : "10.32.0.8:9300",
//       "host" : "10.32.0.8",
//       "ip" : "10.32.0.8",
//       "version" : "5.5.1",
//       "build_hash" : "19c13d0",
//       "roles" : [
//         "master",
//         "data",
//         "ingest"
//       ],
//       "attributes" : {
//         "ml.enabled" : "true"
//       },
//       "http" : {
//         "bound_address" : [
//           "[::]:9200"
//         ],
//         "publish_address" : "10.32.0.8:9200",
//         "max_content_length_in_bytes" : 104857600
//       }
//     }
//   }
// }