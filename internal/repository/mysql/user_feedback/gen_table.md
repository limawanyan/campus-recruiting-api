#### campus-recruiting.user_feedback 
用户反馈

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| ---- | ---- | ---- | ---- | ---- | ---- | ---- | ---- |
| 1 | id |  | int unsigned | PRI | NO | auto_increment |  |
| 2 | user_id | 反馈用户 | int |  | NO |  |  |
| 3 | content | 反馈内容 | varchar(255) |  | YES |  |  |
| 4 | tel | 联系电话 | varchar(11) |  | YES |  |  |
| 5 | type | 反馈类型(优化建议、Bug反馈) | int |  | NO |  |  |
| 6 | state | 状态（0 待处理、1已处理） | int |  | NO |  |  |
| 7 | reply |  | varchar(255) |  | YES |  |  |
| 8 | created_at |  | datetime |  | NO |  |  |
| 9 | updated_at |  | datetime |  | NO |  |  |
