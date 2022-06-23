#### campus-recruiting.dictionary 
数据字典;管理网站类型数据 行业类型 公司属性 讨论帖类型 面经类型

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| ---- | ---- | ---- | ---- | ---- | ---- | ---- | ---- |
| 1 | id | 记录ID | int unsigned | PRI | NO | auto_increment |  |
| 2 | created_at | 创建时间 | datetime |  | NO |  |  |
| 3 | updated_at | 更新时间 | datetime |  | NO |  |  |
| 4 | name | 字典名称 | varchar(50) |  | YES |  |  |
| 5 | type_code | 字典类型标识 | varchar(50) |  | YES |  |  |
| 6 | label | 子项展示值 | varchar(50) |  | YES |  |  |
| 7 | value | 子项真实值 | int |  | YES |  |  |
| 8 | state | 状态，是否可用 | int |  | YES |  |  |
| 9 | description | 描述信息 | varchar(255) |  | YES |  |  |
| 10 | sort | 排序 | int |  | YES |  |  |
| 11 | parent_id | 父级ID | int |  | YES |  |  |
| 12 | readonly | 是否只读，系统初始化固定值，不支持修改（默认0不可修改,1可修改） | int |  | NO |  | 0 |
| 13 | is_deleted | 是否逻辑删除（0 否 1是） | int |  | YES |  | 0 |
