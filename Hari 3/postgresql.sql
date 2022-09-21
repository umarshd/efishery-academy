create table if not exists customers (
	id serial not null primary key,
	customer_name char(50) not null
);

create table if not exists products (
	id serial not null primary key,
	name char(50) not null
);

create table if not exists orders (
	order_id serial not null primary key,
	customer_id int not null references customers (id),
	product_id int not null references products (id),
	order_date date not null,
	total float not null
);

insert into customers (customer_name) values ('bambang'), ('budi'), ('jajang');
insert into products (name) values ('Tas'), ('Sepatu'), ('Tumbler');
insert into orders (customer_id, product_id, order_date,total)
values 
(1, 3, '2022-09-14', 50000),
(3, 2, '2022-09-18', 150000),
(2, 1, '2022-09-19', 250000),
(3, 3, '2022-09-21', 50000),
(2, 3, '2022-09-21', 50000);

update customers set customer_name = 'Bunga' where id = 1;
update products  set name = 'Jacket' where id = 2;
update orders set product_id = 1, total = 250000 where order_id = 5;

delete from orders  where customer_id = 2;
delete from orders where product_id = 1; 
delete from customers where id = 2;
delete from products  where id = 1;

select orders, customers.customer_name , products.name 
from orders 
inner join customers on orders.customer_id = customers.id 
inner join products on products.id = orders.product_id;

