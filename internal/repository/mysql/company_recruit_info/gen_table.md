#### campus-recruiting.company_recruit_info 
公司校招信息

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| ---- | ---- | ---- | ---- | ---- | ---- | ---- | ---- |
| 1 | id | 数据唯一标识 | int unsigned | PRI | NO | auto_increment |  |
| 2 | created_at | 创建时间 | datetime |  | NO |  |  |
| 3 | updated_at | 更新时间 | datetime |  | NO |  |  |
| 4 | description | 描述信息 | varchar(255) |  | YES |  |  |
| 5 | state | 状态，是否可用 | int |  | YES |  |  |
| 6 | company_name | 公司名称 | varchar(60) |  | YES |  |  |
| 7 | company_attribute | 公司属性 | int |  | YES |  |  |
| 8 | recommend_time | 内推时间 | varchar(255) |  | NO |  |  |
| 9 | apply_online_time | 网申时间 | varchar(255) |  | YES |  |  |
| 10 | written_exam_time | 笔试时间（可能存在多个时间） | varchar(255) |  | YES |  |  |
| 11 | interview_time | 面试时间 | varchar(255) |  | YES |  |  |
| 12 | offer_time | offer发放时间 | varchar(255) |  | YES |  |  |
| 13 | industrys | 所属行业名称(多个行业顿号拼接) | varchar(255) |  | YES |  |  |
| 14 | company_logo | 公司logo | varchar(255) |  | YES |  |  |
| 15 | address | 地址(多个地址顿号拼接) | varchar(255) |  | YES |  |  |
| 16 | financing_type | 融资类型 | int |  | YES |  |  |
| 17 | is_deleted | 是否逻辑删除（0 否 1是） | int |  | YES |  | 0 |
