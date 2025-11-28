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
select item,year,sum(quantity) as num from sales group by grouping sets ((item,year),())
select item,year,sum(quantity) as num, count(1) as total from sales group by grouping sets ((item),(year),())
-- 配置给定字段所有匹配的可能性
select item,year,sum(quantity) as num from sales group by cube(item,year)



