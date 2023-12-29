/*
 * PUN street Universal Access - OpenAPI 3.0
 *
 * pua
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type SearchInfo struct {
	CategoryArray []Category `json:"category_array"`

	SearchString string `json:"search_string"`

	PriceHigh float32 `json:"price_high"`

	PriceLow float32 `json:"price_low"`
}
