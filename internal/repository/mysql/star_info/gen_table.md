#### campus-recruiting.star_info 
收藏信息

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| ---- | ---- | ---- | ---- | ---- | ---- | ---- | ---- |
| 1 | id | 记录ID | int unsigned | PRI | NO | auto_increment |  |
| 2 | created_at | 创建时间 | datetime |  | NO |  |  |
| 3 | updated_at | 更新时间 | datetime |  | NO |  |  |
| 4 | from_uid | 收藏用户 | int |  | YES |  |  |
| 5 | type | 收藏类型（面经、面题） | int |  | YES |  |  |
| 6 | topic_id | 收藏主题 | int |  | YES |  |  |
