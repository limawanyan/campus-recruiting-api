#### campus-recruiting.notice 
通知

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| ---- | ---- | ---- | ---- | ---- | ---- | ---- | ---- |
| 1 | id | 记录ID | int unsigned | PRI | NO | auto_increment |  |
| 2 | created_at | 创建时间 | datetime |  | NO |  |  |
| 3 | updated_at | 更新时间 | datetime |  | NO |  |  |
| 4 | event_type | 事件类型（点赞、评论、回复） | int |  | YES |  |  |
| 5 | content_type | 内容类型（面题、面经、面经评论、面题评论） | int |  | YES |  |  |
| 6 | type | 通知类型（用户、系统） | int |  | YES |  |  |
| 7 | trigger_user_id | 触发人ID | int |  | NO |  |  |
| 8 | trigger_user_nickname | 触发人昵称 | varchar(60) |  | NO |  |  |
| 9 | trigger_user_photo | 触发人头像 | varchar(255) |  | NO |  |  |
| 10 | subject_id | 主题ID(面题、面经ID) | int |  | NO |  |  |
| 11 | parent_comment_id | 根评论ID | int |  | YES |  |  |
| 12 | comment_id | 评论ID | int |  | YES |  |  |
| 13 | target_id | 接收人ID | int |  | NO |  |  |
| 14 | title | 标题 | varchar(255) |  | YES |  |  |
| 15 | content | 内容 | text |  | YES |  |  |
| 16 | uri | 内容链接 | varchar(255) |  | YES |  |  |
| 17 | is_read | 是否已读(0 未读  1 已读) | int |  | NO |  | 0 |
