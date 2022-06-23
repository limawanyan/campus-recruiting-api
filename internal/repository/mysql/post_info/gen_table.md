#### campus-recruiting.post_info 
平台帖子(面经/面题)

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| ---- | ---- | ---- | ---- | ---- | ---- | ---- | ---- |
| 1 | id | 记录ID | int unsigned | PRI | NO | auto_increment |  |
| 2 | title | 标题 | varchar(255) |  | NO |  |  |
| 3 | content | 内容 | text |  | NO |  |  |
| 4 | parent_type | 父级板块类型（面经/面题） | int |  | NO |  |  |
| 5 | sub_type | 子板块类型 | int |  | NO |  |  |
| 6 | created_at | 创建时间 | datetime |  | NO |  |  |
| 7 | updated_at | 更新时间 | datetime |  | NO |  |  |
| 8 | from_user_id | 发帖用户ID | int |  | NO |  |  |
| 9 | sort_weight | 热门排序权重（发帖时间戳+点赞数量*N） | int |  | YES |  | 0 |
| 10 | like_num | 点赞数量 | int |  | YES |  | 0 |
| 11 | state | 状态，是否可用 | int |  | YES |  | 1 |
| 12 | comment_num | 评论数量 | int |  | YES |  | 0 |
| 13 | browse_num | 浏览量 | int |  | YES |  | 0 |
| 14 | is_deleted | 是否逻辑删除（0 否 1是） | int |  | YES |  | 0 |
