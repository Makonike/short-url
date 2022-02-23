# short-url

## 说明

a demo base beego of short url

通过自增唯一主键作为短链的短链接系统<br/>

sql如下，主键的offset和inc可以设大一点
```sql
-- auto-generated definition
create schema short_url collate utf8mb4_0900_ai_ci;


-- auto-generated definition
create table t_short
(
    id         bigint unsigned auto_increment comment '短链 id'
        primary key,
    lurl       char(255) null comment '原始长链',
    md5        varchar(32)     null comment 'md5长链',
    gmt_create int          null comment '创建时间'
)
    charset = utf8;


-- auto-generated definition
create table t_ticket
(
    id   bigint unsigned auto_increment
        primary key,
    stub char(4) default '' not null,
    constraint stub
        unique (stub)
);


```

## 参考<br/>
- [短连接系统](https://soulmachine.gitbooks.io/system-design/content/cn/tinyurl.html) 
- [分布式id生成器](https://soulmachine.gitbooks.io/system-design/content/cn/distributed-id-generator.html) 

