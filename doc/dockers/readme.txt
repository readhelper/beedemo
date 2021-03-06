器：这里面什么才是最重要的？container monitoring，一项可以帮助你了解容器内发生什么以及让你获得某种支持的技术，随着Docker自身的monitoring 和 introspection技术API的成熟，已经有最新一波的容器监测技术利用到它们。

以下列举6个Docker相关的监测App和服务，从简单的开源工具到复杂的企业整体解决方案。

cAdvisor
谷歌的container introspection解决方案是cAdvisor，其本身也是在Docker容器内封装的实用工具，它能够快速的产生，运行容器的基本行为等有用信息。它部署和使用很简单，但是它只可以监视在同一个host上运行的其他容器，所以它对多节点部署不是太管用。

Prometheus.io
Soundcloud实验室出品，Prometheus.io是一个系统监控和报警的开源框架。它的强大之处在于它可以摄取来自一个巨大范围数据源的数据包括容器。更厉害的是，容器的数据来源并不局限cAdvisor、Kubernetes、CoreOS的Etcd。

Scout
Scout是一款监视服务，它并不是一个独立的开源项目。但是它比单独的cAdvisor涵盖了更多的领域——换句话说，它从跨多主机聚合数据，并且它有一个插件架构允许更广范围的集成。但它也是一个付费服务，10美元每主机。所以那些寻求一个完全免费的解决方案的开发者一定会失望。

DataDog
同样作为“container monitoring as a service”的是 DataDog。15美元每主机，相比Scout拥有更详细的和灵活的报表功能。

Logentries
许多Docker-monitoring的服务都是长时间的系统监测maven。Logentries是其中之一，该公司似乎已经感觉到这波容器浪潮。它们最近推出了一个免费的Docker容器日志服务，30天试用版本甚至还包括Logentries的其他付费服务，但容器日志服务本身仍然是免费的。小提醒：如果你运行Docker在一个受限制的环境中，例如Google Compute Engine，你需要运行在特许模式下运行Logentries容器。

Sematext
Sematext是一款云端on-prem多个集成监控解决方案，Sematext最近添加Docker监测在其SPM解决方案中。像其他监测方案一样，该监控是通过一个容器管理，通过事件回滚到各种各样的软件源，一个代理可以用来监视多个服务器。还包括对容器级事件的监测（停止、开始等等）和管理容器产生的日志。计划收费方案是3.5美分每服务小时。
