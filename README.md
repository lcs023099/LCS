### 高并发商城秒杀系统
#### 1.解决方案
a.秒杀系统和原来商城系统部署完全隔离，避免相互影响。

b.根据后端承载能力，进行限流，设计降级机制，防止全网挂掉。

c.平台化，所有操作在web平台实时生效。

d.使用Redis承载海量QPS。

#### 2.系统架构设计
<img src="https://ws1.sinaimg.cn/large/70ecf2fbly1fv3a3q4y4oj20ts0qstap.jpg" height=340 width=440>

#### 2.1系统设计及配置管理
<img src="https://ws1.sinaimg.cn/large/70ecf2fbly1fv3hypkbe0j21g20y4q6j.jpg" height=340 width=440>

#### 2.2接入层核心功能分析
<img src="https://ws1.sinaimg.cn/large/70ecf2fbly1fv3avx4w7qj20mi082aar.jpg" height=120 width=440>

#### 2.3秒杀逻辑层核心功能分析
<img src="https://ws1.sinaimg.cn/large/70ecf2fbly1fv3bisz40jj20t409875b.jpg" height=150 width=440>
