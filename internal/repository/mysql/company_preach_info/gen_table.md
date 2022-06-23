#### campus-recruiting.company_preach_info 
公司宣讲信息

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| ---- | ---- | ---- | ---- | ---- | ---- | ---- | ---- |
| 1 | id | 数据唯一标识 | int unsigned | PRI | NO | auto_increment |  |
| 2 | created_at | 创建时间 | datetime |  | NO |  |  |
| 3 | updated_at | 更新时间 | datetime |  | NO |  |  |
| 4 | state | 状态，是否可用 | int |  | YES |  |  |
| 5 | company_recruit_id | 所属校招公司 | int |  | YES |  |  |
| 6 | preach_time | 宣讲时间 | datetime |  | YES |  |  |
| 7 | type | 宣讲类型(在线/线下) | int |  | YES |  |  |
| 8 | address | 宣讲详细地址 | varchar(255) |  | YES |  |  |
| 9 | school | 宣讲学校 | varchar(60) |  | YES |  |  |
| 10 | is_deleted | 是否逻辑删除（0 否 1是） | int |  | YES |  | 0 |
