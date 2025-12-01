select * from employees where id = 1;

select array[1,2,3]::integer[];
select name, (select e.core/2 from employees as e where e.id = employees.id) from employees;

create table arr (f1 int[], f2 int[]);
insert into arr values (array[[1,2], [3,4]], array[[5,6], [7,8]]);
select array[f1, f2, '{{10,11}, {12,13}}'::int[]] from arr;
SELECT ROW(1,2.5,'this is a test');

select *, case
            when core < 0 then 'bad'
            when core < 60 then 'no good'
            when core < 80 then 'great'
            when core < 90 then 'good'
            else 'very good'
    end
from employees order by core DESC;

select *,
case
    when core >= 90 then 'high'
    when core >= 80 then 'mid'
    when core >= 70 then 'low'
    else 'low low'
end
from employees order by
case
    when core >= 90 then 1
    when core >= 80 then 2
    when core >= 70 then 3
    else 4
end;

-- with recursive 递归操作 向下查询

with recursive menu_node as (
    -- 非递归查询
    select id, name, parent_id, sort, 1 as depth from menus where parent_id is null
    union all
    -- 递归查询
    select s.id, s.name, s.parent_id, s.sort, m.depth + 1 from menus s join menu_node m on m.id = s.parent_id
)
select * from menu_node order by depth, sort desc;

-- with recursive 递归操作 向上查询

with recursive c_node as (
    select * from menus where id = 8
    union all
    select m.* from menus m
    inner join c_node c on c.parent_id = m.id
)
select string_agg(name, '/') as path from( select name from c_node order by id ) as cnn

-- INSERT INTO menus (name, link) VALUES
--                                    ('首页', '/'),
--                                    ('关于我们', '/about'),
--                                    ('服务', '/services');

-- 假设上面的查询返回的ID为1, 2, 3（对应于“首页”, “关于我们”, “服务”）
-- INSERT INTO menus (name, link, parent_id) VALUES
--   ('首页 - 子项1', '/home-sub1', 1), -- 假设1是首页的ID
--   ('关于我们 - 子项1', '/about-sub1', 2), -- 假设2是“关于我们”的ID
--   ('服务 - 子项1', '/services-sub1', 3); -- 假设3是“服务”的ID

-- 角色权限管理 postgresql 不存在 用户的概念，权限管理都是通过角色进行控制
-- grant(授予) revoke(撤销)
-- 创建登录角色
create role king WITH LOGIN PASSWORD '123456';
-- 授予权限
GRANT all on table test.public.menus TO king;
-- 撤销数据库权限
REVOKE ALL ON DATABASE test FROM king;
-- 赋予全部表的权限
grant all on all tables in schema public to king;
-- 授权schema角色权限
grant all on schema custom to king2;
-- 授权其他schema表权限
grant select on table custom.users to king2;
-- 修改用户密码
alter role King2 with password '123456';

-- 撤销全部表的权限
revoke all on ALL TABLES IN SCHEMA PUBLIC FROM king;
-- 分配数据库 所有者
-- * 数据库拥有者不能被撤销，原因是数据库必须有一个拥有者。
-- * 可以进行转移操作进行权限撤销
-- * 一般情况不推荐使用
alter database test owner to postgres;

-- 分组查询
select item,year,sum(quantity) as num from sales group by grouping sets ((item,year),());
select item,year,sum(quantity) as num, count(1) as total from sales group by grouping sets ((item),(year),());
-- 配置给定字段所有匹配的可能性
select item,year,sum(quantity) as num from sales group by cube(item,year);

-- distinct
SELECT distinct customer_id FROM orders;
SELECT count(DISTINCT customer_id) FROM orders;
select DISTINCT ON (customer_id) * from orders order by customer_id,amount ASC;

-- timestamp date timestamptz interval 时间类型
select '2025-12-01 15:18:02'::timestamptz + interval '1 day';
select '2025-12-01'::date + interval '1 day';

-- enum 枚举类型 可以在创建字段时使用枚举类型
CREATE TYPE mood AS ENUM ('sad', 'ok', 'happy');

-- json jsonb变量类型操作
select '5 '::jsonb;
SELECT '{"bar": "baz", "balance": 7.77, "active": false}'::json;
select '["bar", "foo", "666"]'::jsonb ? '666';
SELECT '["BAR"]'::json;
SELECT '["BAR"] '::jsonb;
select '{"lee": 2}'::json;
select '{"lee": 2.00, "king": 3, "zhang": 2.65}'::jsonb ? 'king';
select '["liu"]'::jsonb || '["lee"]'::jsonb || '["lee"]'::jsonb;
select '{"xx": 1}'::jsonb || '{"xxx": 3}'::jsonb;
-- json 和 jonsb 操作符
-- ->、->>、#>、#>> 单个箭头返回jsonb 两个箭头返回文本 字符串代表索引 数值代表索引
select '{"king": 1, "lee": 2}'::jsonb -> 'king'; --字符串为键 返回键对应的值
select '[1,23,3,45,9]'::jsonb -> 1; -- 数值为下标索引 返回键对应的值
select '[{"a":"foo"},{"b":"bar"},{"c":"baz"}]'::json->2;  -- 返回键对应的值
select '{"a": {"b":{"c": "foo"}}}'::json#>'{a,b}'; -- 查询子级元素 返回对应的json值
select '{"a":[1,2,3],"b":[4,5,6]}'::json#>>'{a,2}'; -- 查询子元素 返回对应text值
-- jsonb 对比
select '{"a":1, "b":2}'::jsonb @> '{"b":2}'::jsonb; -- 右侧顶层元素是否在左侧匹配 返回 boolean
select '{"b":2}'::jsonb <@ '{"a":1, "b":2}'::jsonb; -- 和@>比对相反
select '{"a":1, "b":2}'::jsonb ? 'b'; -- 键是否在左测jsonb顶层存在 返回 boolean
select '{"a":1, "b":2, "c":3}'::jsonb ?| array['c', 'd']; -- 右侧数组中的元素是否在左侧jsonb中存在对应键 任何满足返回true
select '["a", "b"]'::jsonb ?& array['a', 'b']; -- 右侧数组中的元素是否在左侧jsonb中存在对应键 全部满足返回true
-- jsonb 操作
select '["a", "b"]'::jsonb || '["c", "d"]'::jsonb; -- jsonb 拼接 返回拼接后新的 jsonb
select '{"a": "1", "b": "2"}'::jsonb - 'a'; -- 删除左侧的键 返回删除后的jsonb
select '["a", "b"]'::jsonb - 1; -- 删除左侧指定的下标元素 返回删除后的jsonb
select '["a", {"b":1}]'::jsonb #- '{1,b}'; -- 删除对应的子元素 返回删除后的jsonb。{1,b}:表示第一个索引中的b元素

