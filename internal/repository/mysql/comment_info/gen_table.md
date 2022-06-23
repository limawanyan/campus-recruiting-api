#### campus-recruiting.comment_info 
评论信息

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| ---- | ---- | ---- | ---- | ---- | ---- | ---- | ---- |
| 1 | id | 记录ID | int unsigned | PRI | NO | auto_increment |  |
| 2 | created_at | 创建时间 | datetime |  | NO |  |  |
| 3 | updated_at | 更新时间 | datetime |  | NO |  |  |
| 4 | topic_id | 主题ID | int |  | YES |  |  |
| 5 | type | 评论类型（讨论帖/面题/薪资爆料） | int |  | YES |  |  |
| 6 | parent_id | 父级评论ID | int |  | YES |  | 0 |
| 7 | from_user_id | 评论人 | int |  | YES |  |  |
| 8 | to_user_id | 被评论人 | int |  | YES |  |  |
| 9 | from_nickname | 评论人昵称 | varchar(60) |  | YES |  |  |
| 10 | from_head | 评论人头像 | varchar(255) |  | YES |  |  |
| 11 | to_nickname | 被评论人昵称 | varchar(60) |  | YES |  |  |
| 12 | content | 评论内容 | text |  | YES |  |  |
| 13 | like_num | 点赞数量 | int |  | YES |  |  |
| 14 | is_deleted | 是否逻辑删除（0 否 1是） | int |  | YES |  | 0 |
