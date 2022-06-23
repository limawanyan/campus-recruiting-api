#### campus-recruiting.like_info 
点赞信息表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| ---- | ---- | ---- | ---- | ---- | ---- | ---- | ---- |
| 1 | id | 记录ID | int unsigned | PRI | NO | auto_increment |  |
| 2 | created_at | 创建时间 | datetime |  | NO |  |  |
| 3 | updated_at | 更新时间 | datetime |  | NO |  |  |
| 4 | topic_id | 主题ID | int |  | YES |  |  |
| 5 | type | 点赞类型（讨论帖/面题/讨论帖评论/面题评论/薪资可信度） | int |  | NO |  |  |
| 6 | from_uid | 来自哪个用户 | int |  | NO |  |  |
| 7 | state | 状态，点赞是否有效(0 无效 1 有效) | int |  | NO |  | 1 |
