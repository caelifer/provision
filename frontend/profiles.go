package frontend

import (
	"github.com/VictorLowther/jsonpatch2"
	"github.com/digitalrebar/provision/backend"
	"github.com/digitalrebar/provision/models"
	"github.com/gin-gonic/gin"
)

// ProfileResponse returned on a successful GET, PUT, PATCH, or POST of a single profile
// swagger:response
type ProfileResponse struct {
	// in: body
	Body *models.Profile
}

// ProfilesResponse returned on a successful GET of all the profiles
// swagger:response
type ProfilesResponse struct {
	//in: body
	Body []*models.Profile
}

// ProfileParamsResponse return on a successful GET of all Profile's Params
// swagger:response
type ProfileParamsResponse struct {
	// in: body
	Body map[string]interface{}
}

// ProfileParamResponse return on a successful GET of a single Param for a Profile
// swagger:response
type ProfileParamResponse struct {
	// in: body
	Body interface{}
}

// ProfileBodyParameter used to inject a Profile
// swagger:parameters createProfile putProfile
type ProfileBodyParameter struct {
	// in: body
	// required: true
	Body *models.Profile
}

// ProfilePatchBodyParameter used to patch a Profile
// swagger:parameters patchProfile
type ProfilePatchBodyParameter struct {
	// in: body
	// required: true
	Body jsonpatch2.Patch
}

// ProfilePathParameter used to name a Profile in the path
// swagger:parameters putProfiles getProfile putProfile patchProfile deleteProfile getProfileParams postProfileParams headProfile
type ProfilePathParameter struct {
	// in: path
	// required: true
	Name string `json:"name"`
}

// ProfileParamsPathParameter used to get or set a single Parameter in a Profile
// swagger:parameters getProfileParam postProfileParam
type ProfileParamsPathParameter struct {
	// in: path
	// required: true
	Name string `json:"name"`
	// in: path
	// required: true
	Key string `json:"key"`
}

// ProfileParamsBodyParameter used to set Profile Params
// swagger:parameters postProfileParams
type ProfileParamsBodyParameter struct {
	// in: body
	// required: true
	Body map[string]interface{}
}

// ProfileParamBodyParameter used to set a single Param on a Profile
// swagger:parameters postProfileParam
type ProfileParamBodyParameter struct {
	// in: body
	// required: true
	Body map[string]interface{}
}

// ProfileListPathParameter used to limit lists of Profile by path options
// swagger:parameters listProfiles
type ProfileListPathParameter struct {
	// in: query
	Offest int `json:"offset"`
	// in: query
	Limit int `json:"limit"`
	// in: query
	Available string
	// in: query
	Valid string
	// in: query
	ReadOnly string
	// in: query
	Name string
}

func (f *Frontend) InitProfileApi() {
	// swagger:route GET /profiles Profiles listProfiles
	//
	// Lists Profiles filtered by some parameters.
	//
	// This will show all Profiles by default.
	//
	// You may specify:
	//    Offset = integer, 0-based inclusive starting point in filter data.
	//    Limit = integer, number of items to return
	//
	// Functional Indexs:
	//    Name = string
	//    Available = boolean
	//    Valid = boolean
	//    ReadOnly = boolean
	//
	// Functions:
	//    Eq(value) = Return items that are equal to value
	//    Lt(value) = Return items that are less than value
	//    Lte(value) = Return items that less than or equal to value
	//    Gt(value) = Return items that are greater than value
	//    Gte(value) = Return items that greater than or equal to value
	//    Between(lower,upper) = Return items that are inclusively between lower and upper
	//    Except(lower,upper) = Return items that are not inclusively between lower and upper
	//
	// Example:
	//    Name=fred - returns items named fred
	//    Name=Lt(fred) - returns items that alphabetically less than fred.
	//
	// Responses:
	//    200: ProfilesResponse
	//    401: NoContentResponse
	//    403: NoContentResponse
	//    406: ErrorResponse
	f.ApiGroup.GET("/profiles",
		func(c *gin.Context) {
			f.List(c, &backend.Profile{})
		})

	// swagger:route POST /profiles Profiles createProfile
	//
	// Create a Profile
	//
	// Create a Profile from the provided object
	//
	//     Responses:
	//       201: ProfileResponse
	//       400: ErrorResponse
	//       401: NoContentResponse
	//       403: NoContentResponse
	//       422: ErrorResponse
	f.ApiGroup.POST("/profiles",
		func(c *gin.Context) {
			b := &backend.Profile{}
			f.Create(c, b)
		})
	// swagger:route GET /profiles/{name} Profiles getProfile
	//
	// Get a Profile
	//
	// Get the Profile specified by {name} or return NotFound.
	//
	//     Responses:
	//       200: ProfileResponse
	//       401: NoContentResponse
	//       403: NoContentResponse
	//       404: ErrorResponse
	f.ApiGroup.GET("/profiles/:name",
		func(c *gin.Context) {
			f.Fetch(c, &backend.Profile{}, c.Param(`name`))
		})

	// swagger:route HEAD /profiles/{name} Profiles headProfile
	//
	// See if a Profile exists
	//
	// Return 200 if the Profile specifiec by {name} exists, or return NotFound.
	//
	//     Responses:
	//       200: NoContentResponse
	//       401: NoContentResponse
	//       403: NoContentResponse
	//       404: NoContentResponse
	f.ApiGroup.HEAD("/profiles/:name",
		func(c *gin.Context) {
			f.Exists(c, &backend.Profile{}, c.Param(`name`))
		})

	// swagger:route PATCH /profiles/{name} Profiles patchProfile
	//
	// Patch a Profile
	//
	// Update a Profile specified by {name} using a RFC6902 Patch structure
	//
	//     Responses:
	//       200: ProfileResponse
	//       400: ErrorResponse
	//       401: NoContentResponse
	//       403: NoContentResponse
	//       404: ErrorResponse
	//       406: ErrorResponse
	//       422: ErrorResponse
	f.ApiGroup.PATCH("/profiles/:name",
		func(c *gin.Context) {
			f.Patch(c, &backend.Profile{}, c.Param(`name`))
		})

	// swagger:route PUT /profiles/{name} Profiles putProfile
	//
	// Put a Profile
	//
	// Update a Profile specified by {name} using a JSON Profile
	//
	//     Responses:
	//       200: ProfileResponse
	//       400: ErrorResponse
	//       401: NoContentResponse
	//       403: NoContentResponse
	//       404: ErrorResponse
	//       422: ErrorResponse
	f.ApiGroup.PUT("/profiles/:name",
		func(c *gin.Context) {
			f.Update(c, &backend.Profile{}, c.Param(`name`))
		})

	// swagger:route DELETE /profiles/{name} Profiles deleteProfile
	//
	// Delete a Profile
	//
	// Delete a Profile specified by {name}
	//
	//     Responses:
	//       200: ProfileResponse
	//       401: NoContentResponse
	//       403: NoContentResponse
	//       404: ErrorResponse
	//       422: ErrorResponse
	f.ApiGroup.DELETE("/profiles/:name",
		func(c *gin.Context) {
			f.Remove(c, &backend.Profile{}, c.Param(`name`))
		})

	pGetAll, pSetAll, pGetOne, pSetOne := f.makeParamEndpoints(&backend.Profile{}, "name")

	// swagger:route GET /profiles/{name}/params Profiles getProfileParams
	//
	// List profile params Profile
	//
	// List Profile parms for a Profile specified by {name}
	//
	//     Responses:
	//       200: ProfileParamsResponse
	//       401: NoContentResponse
	//       403: NoContentResponse
	//       404: ErrorResponse
	f.ApiGroup.GET("/profiles/:name/params", pGetAll)

	// swagger:route POST /profiles/{name}/params Profiles postProfileParams
	//
	// Set/Replace all the Parameters for a profile specified by {name}
	//
	//     Responses:
	//       200: ProfileParamsResponse
	//       401: NoContentResponse
	//       403: NoContentResponse
	//       404: ErrorResponse
	//       409: ErrorResponse
	f.ApiGroup.POST("/profiles/:name/params", pSetAll)

	// swagger:route GET /profiles/{name}/params/{key} Profiles getProfileParam
	//
	// Get a single profile parameter
	//
	// Get a single parameter {key} for a Profile specified by {name}
	//
	//     Responses:
	//       200: ProfileParamResponse
	//       401: NoContentResponse
	//       403: NoContentResponse
	//       404: ErrorResponse
	f.ApiGroup.GET("/profiles/:name/params/*key", pGetOne)

	// swagger:route POST /profiles/{name}/params/{key} Profiles postProfileParam
	//
	// Set as single Parameter {key} for a profile specified by {name}
	//
	//     Responses:
	//       200: ProfileParamResponse
	//       401: NoContentResponse
	//       403: NoContentResponse
	//       404: ErrorResponse
	//       409: ErrorResponse
	f.ApiGroup.POST("/profiles/:name/params/*key", pSetOne)
}
