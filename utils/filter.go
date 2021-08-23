//=============================================================================
// developer: boxlesslabsng@gmail.com
// Pagination library
//=============================================================================
 
/**
 **
 * @struct FilterUtil
 **
 * @SetPaging() set pagination from echo context
**/

package utils

import (
	"strconv"

	"github.com/labstack/echo"
)

type FilterUtil struct {
	echo.Context
}

func (filter *FilterUtil) SetPaging(ctx echo.Context) map[string]int64 {
	queries := ctx.QueryParams()
	pageSize := int64(50)
	page := int64(1)
	if p, err := strconv.ParseInt(queries.Get("pageSize"), 10, 64); err == nil {
		pageSize = p
	}

	if pp, err := strconv.ParseInt(queries.Get("pageNum"), 10, 64); err == nil {
		page = pp
	}
	skip := (page - 1) * pageSize

	return map[string]int64{
		"page": page, "pageSize": pageSize, "skip": skip,
	}
}