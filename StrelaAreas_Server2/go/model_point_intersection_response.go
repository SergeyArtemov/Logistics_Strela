/*
 * Area API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// PointIntersectionResponse - Ответ на запрос пересечений с точкой.
type PointIntersectionResponse struct {

	// Массив пересечений с зонами маршрутизатора.
	Areas []AreaIntersection `json:"areas,omitempty"`
}