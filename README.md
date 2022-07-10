# eids

基于雪花算法的ids生成器

### 性能分析
 - 18位的id,约 73w/s
 - 19位的id,约 600w/s

### 使用方式
- 可用client+server 配置ids服务
- 可以配置不同的服务id,集群化管理,也可用保证唯一性

### 说明
- 机器码不一样所产出的id不会相同