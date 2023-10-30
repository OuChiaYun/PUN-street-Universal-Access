/*
 * PUN street Universal Access - OpenAPI 3.0
 *
 * pua
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type OrderInfo struct {

	CustomerId int64 `json:"customer_id"`

	DiscountId int64 `json:"discount_id"`

	CartId int64 `json:"cart_id"`

	StoreId int64 `json:"store_id"`

	OrderStatus int64 `json:"order_status"`

	OrderDate string `json:"order_date"`

	TakingAddress string `json:"taking_address"`

	TakingMethod int64 `json:"taking_method"`

	TotalPrice int64 `json:"total_price"`
}
