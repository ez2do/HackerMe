/*
 * forwardservice_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package category_pkg

import "forwardservice_lib/configuration_pkg"

/*
 * Interface for the CATEGORY_IMPL
 */
type CATEGORY interface {
    CategoryInfo (int64) (error)
}

/*
 * Factory for the CATEGORY interaface returning CATEGORY_IMPL
 */
func NewCATEGORY(config configuration_pkg.CONFIGURATION) *CATEGORY_IMPL {
    client := new(CATEGORY_IMPL)
    client.config = config
    return client
}