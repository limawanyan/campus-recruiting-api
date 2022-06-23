#### campus-recruiting.user 
平台用户

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| ---- | ---- | ---- | ---- | ---- | ---- | ---- | ---- |
| 1 | id | 数据唯一标识 | int unsigned | PRI | NO | auto_increment |  |
| 2 | created_at | 创建时间 | datetime |  | NO |  |  |
| 3 | updated_at | 更新时间 | datetime |  | NO |  |  |
| 4 | wx_openid | 微信小程序授权openid | varchar(255) |  | YES |  |  |
| 5 | nickname | 用户昵称 | varchar(60) |  | YES |  |  |
| 6 | phone | 手机号/账号 | varchar(11) |  | YES |  |  |
| 7 | password | 密码 | varchar(20) |  | YES |  |  |
| 8 | sex | 用户性别 | int |  | YES |  |  |
| 9 | head_portrait | 用户头像 | varchar(255) |  | YES |  |  |
| 10 | graduation_year | 毕业年份 | datetime |  | YES |  |  |
| 11 | intentional_work | 意向工作 | varchar(60) |  | YES |  |  |
| 12 | autograph | 签名/简介 | varchar(60) |  | YES |  |  |
| 13 | email | 邮箱 | varchar(30) |  | YES |  |  |
| 14 | state | 状态，是否可用 | int |  | YES |  |  |
| 15 | is_deleted | 是否逻辑删除（0 否 1是） | int |  | YES |  | 0 |
