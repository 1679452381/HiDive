## 视频网站后台

## api
        /user
        登录
        注册
        登出
        查看用户信息
        关注用户
        取消关注用户
        /video
        获取视频列表
        获取热门视频
        根据关键字搜索视频
        查看视频
        用户上传视频
        用户删除视频
        点赞视频
        取消赞视频
        点踩视频
        取消点踩视频
        /comment
        游客获取视频评论
        用户发表视频评论
        用户发表评论的评论
        用户删除评论
        /barrage 弹幕
        游客获取视频弹幕
        用户发表视频弹幕
        用户删除弹幕
        /chat 

## 项目结构 
* app 
  * 所有业务代码 api rpc以及mq（消息队列 延迟队列 定时任务）
* common 
  * 通用组件 error middleware interceptor tool ctxdata
* data 
  * 项目以来中间件（mysql es redis grafana等）产生的数据 此目录放在git忽略文件中 不提交
* deploy 配置文件
* doc 文档
* modd.conf 热加载 每次更新代码 自动重新加载