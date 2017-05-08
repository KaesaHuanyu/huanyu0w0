Zookeeper Cluster
=================

Zookeeper cluster for swarm. Compatible with DCE 2.0 and later only, for DCE 1.x, please use zookeeper-cluster-legacy.

测试zookeeper

打开一个zookeeper01 的控制台，新建一个文件 zk-test.sh，写入下面内容
```
#!/bin/bash
zk1="zkCli.sh "
zk2="zkCli.sh -server zookeeper02:2181"
zk3="zkCli.sh -server zookeeper03:2181"
check_zk_cluter(){
$zk1 delete /test >/dev/null 2>&1
$zk1 delete /test2 >/dev/null 2>&1
$zk1 delete /test3 >/dev/null 2>&1

key=1
echo "add /test=${key} with zk1: "
$zk1 create /test ${key}  2>&1 |grep  "Created"
echo "get /test from all:"
$zk1 get /test 2>&1 |grep "^${key}$"
$zk2 get /test 2>&1 |grep "^${key}$"
$zk3 get /test 2>&1 |grep "^${key}$"

key=2
echo "add /test2=${key} with zk2: "
$zk2 create /test2 ${key} 2>&1 |grep  "Created"
echo "get /test2 from all:"
$zk1 get /test2 2>&1 |grep "^${key}$"
$zk2 get /test2 2>&1 |grep "^${key}$"
$zk3 get /test2 2>&1 |grep "^${key}$"

key=3
echo "add /test3=${key} with zk1: "
$zk1 create /test3 ${key} 2>&1 |grep  "Created"
echo "get /test3 from all:"
$zk1 get /test3 2>&1 |grep "^${key}$"
$zk2 get /test3 2>&1 |grep "^${key}$"
$zk3 get /test3 2>&1 |grep "^${key}$"
}
check_zk_reboot(){
key=3
$zk1 ls / 2>&1 |grep -v "INFO"
echo "get /test3 from all:"
$zk1 get /test3 2>&1 |grep "^${key}$"
$zk2 get /test3 2>&1 |grep "^${key}$"
$zk3 get /test3 2>&1 |grep "^${key}$"
}


check_kafka(){
    #http://blog.csdn.net/xw_classmate/article/details/53264303
    kafka-topics.sh --zookeeper zookeeper01:2181,zookeeper02:2181,zookeeper03:2181 --create  --partitions 10 --replication-factor 3 --topic daocloud
    kafka-simple-consumer-shell.sh --broker-list "broker01:9092,broker02:9092,broker03:9092" --print-offsets --topic daocloud &
    echo "you can write some thing ,it will print again by consumer"
    kafka-console-producer.sh --topic daocloud --broker-list "broker01:9092,broker02:9092,broker03:9092"

}

check_zk_cluter

```
然后通过`bash zk-test.sh`执行，可以看到测试结果（通过往每一个zk写入一个值并检查是否能够在每一个zk上都能查询到）


## Kafka
Kafka is a distributed streaming platform.


检查kafka是否正常
部署后，打开kafka-manager,点击 Cluster ->  Add Cluster,
输入：
Cluster Name: kafka，
Cluster Zookeeper Hosts: zookeeper01:2181,zookeeper02:2181,zookeeper03:2181
Kafka Version: 0.9.0.1
点击save，然后点击 "Go to cluster view.",查看集群各种状态，（默认启动了3个 Brokers）


测试kafka 信息发送接收：
1.接收消息端。打开任意一个kafka容器的控制台，之行下面脚本，

```
kafka-topics.sh --zookeeper zookeeper01:2181,zookeeper02:2181,zookeeper03:2181 --create  --partitions 10 --replication-factor 3 --topic daocloud
kafka-console-consumer.sh  --zookeeper "zookeeper01:2181,zookeeper02:2181,zookeeper03:2181"  --topic daocloud --from-beginning
```

2.发送消息端。打开一个kafka容器窗口，之行下面的脚本，并输入hello，然后回车，可以在之前打开的窗口看到输入的信息
```
kafka-console-producer.sh --topic daocloud --broker-list "broker01:9092,broker02:9092,broker03:9092"
```

