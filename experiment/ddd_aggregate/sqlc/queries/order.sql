-- name: UpsertOrder :exec
INSERT INTO orders (id, customer_id, order_status)
VALUES (?, ?, ?)
ON DUPLICATE KEY UPDATE
  customer_id = VALUES(customer_id),
  order_status = VALUES(order_status);

-- name: GetOrder :one
SELECT id, customer_id, order_status
FROM orders
WHERE id = ?;

-- name: DeleteOrderItemsByOrderID :exec
DELETE FROM order_items
WHERE order_id = ?;

-- name: CreateOrderItem :exec
INSERT INTO order_items (order_id, product_id, quantity, unit_price)
VALUES (?, ?, ?, ?);

-- name: ListOrderItemsByOrderID :many
SELECT order_id, product_id, quantity, unit_price
FROM order_items
WHERE order_id = ?;
