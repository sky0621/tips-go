# performance

## sample1

```
select i.order_id as order_id,
       i.id as item_id,
       i.product_id as product_id,
       p.name as product_name,
       p.price as price
from order_items i
         inner join orders o on i.order_id = o.id
         inner join products p on i.product_id = p.id
where i.quantity > 2
and product_id > 1000 and product_id < 6000
group by order_id, item_id, product_id, product_name;
```

22秒
