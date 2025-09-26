1. 数据库选型 和 设计
    1. 数据库选型： Mysql，因为比较熟悉
    2. 设计：movie 和 rating 两个表，除了 boxoffice 字段不存其他字段全存，不需要外键，不需要索引
2. 后端服务选型 和 设计
    1. 后端服务选型：GIN + GORM
    2. 设计：GIN handler 实现 openapi 接口，GORM 实现对数据库的操作
3. 在完成项目后，思考一下整体项目还有哪些可以优化的内容
    1. GetMovies 接口请求 boxoffice 可以再优化一下，池化或者 mapreduce
    2. 返回码的治理，可以考虑在中间价里去做