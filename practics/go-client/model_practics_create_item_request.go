/*
 * api/purchase-api.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type PracticsCreateItemRequest struct {
	Title string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Cost string `json:"cost,omitempty"`
}
