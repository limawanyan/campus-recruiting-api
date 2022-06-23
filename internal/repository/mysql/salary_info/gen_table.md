#### campus-recruiting.salary_info 
薪资爆料信息

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| ---- | ---- | ---- | ---- | ---- | ---- | ---- | ---- |
| 1 | id | 数据唯一标识 | int unsigned | PRI | NO | auto_increment |  |
| 2 | created_at | 创建时间 | datetime |  | NO |  |  |
| 3 | updated_at | 更新时间 | datetime |  | NO |  |  |
| 4 | description | 薪资描述（薪资详情） | varchar(255) |  | YES |  |  |
| 5 | state | 状态，是否上下架 | int |  | YES |  |  |
| 6 | company_name | 公司名称 | varchar(60) |  | YES |  |  |
| 7 | post_name | 岗位名称 | varchar(60) |  | YES |  |  |
| 8 | city_name | 城市名称 | varchar(60) |  | YES |  |  |
| 9 | remarks | 备注（offer相关信息） | varchar(255) |  | YES |  |  |
| 10 | education | 学历 | int |  | YES |  |  |
| 11 | industry | 行业 | int |  | YES |  |  |
| 12 | reliability | 可信度 | int |  | YES |  |  |
| 13 | from_user_id | 爆料人 | int |  | YES |  |  |
| 14 | fuzzy_query | 模糊查询（公司+职位+地点） | varchar(60) |  | YES |  |  |
