## demo关于kubebuilder的使用测试
包括如下内容：
1. 模版创建deployment文件
2. 系统资源监听 owns方法与自定义队列

## 1. 初始化项目 创建API
> kubebuilder 3.3.0  go1.17.11
```shell
kubebuilder init --domain demo.io --repo github.com/ylinyang/kubebuilder-demo
kubebuilder create api --group test --version v1 --kind App
make manifests  # 生成crd文件
make install    # 安装到k8s集群
```


