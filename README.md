1. 数据库选型 和 设计
    1. 数据库选型： Mysql，因为比较熟悉
    2. 设计：movie 和 rating 两个表，~~除了 boxoffice 字段不存其他字段全存~~，不需要外键，不需要索引
    3. 按要求需要在创建电影时把 boxoffice 票房信息同步写进 db，不是很认可这个设计，票房信息应该实时从接口里拿，如果觉得实时拿会对查询接口性能有影响，可以考虑两种方案
        1. 第一版 GetMovies 接口的实现，电影数不太多时直接并发从接口实时拿，电影数非常多时可以 mapreduce
           优化，最终的效果是能实时拿到票房数据，且接口性能不会有太多影响（多了调下游接口的时延）
        2. 加入缓存层，如 redis，在创建电影是就将 movieTitle 或某唯一标识记录到 redis 中，异步线程定时的调票房接口然后更新数据，GetMovies
           从 redis 中拿票房数据
    4. 为什么不需要外键，百害一利
    5. 为什么不需要索引
        1. 初期只需要必要的索引（主键索引）为后续的性能优化留出灵活空间，后续可以考虑，YEAR(release_date)、budget、genre
           COLLATE utf8mb4_unicode_ci、genre COLLATE utf8mb4_unicode_ci 单列索引
        2. title 模糊匹配、mpa_rating 区分度不高，不考虑创建索引
2. 后端服务选型 和 设计
    1. 后端服务选型：GIN + GORM
    2. 设计：GIN handler 实现 openapi 接口，GORM 实现对数据库的操作
3. 在完成项目后，思考一下整体项目还有哪些可以优化的内容
    1. ~~GetMovies 接口请求 boxoffice 可以再优化一下，池化或者 mapreduce~~
    2. 返回码的治理，可以考虑在中间价里去做
    3. GetMovies 接口提供的查询条件太多，性能会有问题，维护复杂度也高