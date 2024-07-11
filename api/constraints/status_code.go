// Package `constraints` provides constraints that will replace "meaningless" values with
// names, providing a better understanding of what that value does.
package constraints

type StatusCode int16

const HTTP_STATUS_OK = 200
const HTTP_STATUS_INTERNAL_SERVER_ERROR = 500
const HTTP_STATUS_BAD_REQUEST = 400
