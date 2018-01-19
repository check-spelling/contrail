package api

import (
	"database/sql"
	"net/http"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/db"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"

	log "github.com/sirupsen/logrus"
)

//ServiceObjectRESTAPI
type ServiceObjectRESTAPI struct {
	DB *sql.DB
}

type ServiceObjectCreateRequest struct {
	Data *models.ServiceObject `json:"service-object"`
}

type ServiceObjectUpdateRequest struct {
	Data map[string]interface{} `json:"service-object"`
}

//Path returns api path for collections.
func (api *ServiceObjectRESTAPI) Path() string {
	return "/service-objects"
}

//LongPath returns api path for elements.
func (api *ServiceObjectRESTAPI) LongPath() string {
	return "/service-object/:id"
}

//SetDB sets db object
func (api *ServiceObjectRESTAPI) SetDB(db *sql.DB) {
	api.DB = db
}

//Create handle a Create REST API.
func (api *ServiceObjectRESTAPI) Create(c echo.Context) error {
	requestData := &ServiceObjectCreateRequest{
		Data: models.MakeServiceObject(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_object",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	model := requestData.Data
	if model == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	if model.UUID == "" {
		model.UUID = uuid.NewV4().String()
	}

	if model.FQName == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing FQName")
	}

	auth := common.GetAuthContext(c)
	model.Perms2.Owner = auth.ProjectID()
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			return db.CreateServiceObject(tx, model)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_object",
		}).Debug("db create failed on create")
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusCreated, requestData)
}

//Update handles a REST Update request.
func (api *ServiceObjectRESTAPI) Update(c echo.Context) error {
	id := c.Param("id")
	requestData := &ServiceObjectUpdateRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_object",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	model := requestData.Data
	if model == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	auth := common.GetAuthContext(c)
	ok := common.SetValueByPath(model, "Perms2.Owner", ".", auth.ProjectID())
	if !ok {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			return db.UpdateServiceObject(tx, id, model)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_object",
		}).Debug("db update failed")
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, map[string]map[string]string{
		"service-object": {
			"uuid": id,
			"uri":  "/" + "service-object" + "/" + id},
	})
}

//Delete handles a REST Delete request.
func (api *ServiceObjectRESTAPI) Delete(c echo.Context) error {
	id := c.Param("id")
	auth := common.GetAuthContext(c)
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			return db.DeleteServiceObject(tx, id, auth)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return echo.NewHTTPError(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//Show handles a REST Show request.
func (api *ServiceObjectRESTAPI) Show(c echo.Context) error {
	id := c.Param("id")
	auth := common.GetAuthContext(c)
	var result []*models.ServiceObject
	var err error
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListServiceObject(tx, &common.ListSpec{
				Limit: 1000,
				Auth:  auth,
				Filter: common.Filter{
					"uuid": []string{id},
				},
			})
			return err
		}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"service_object": result,
	})
}

//List handles a List REST API Request.
func (api *ServiceObjectRESTAPI) List(c echo.Context) error {
	var result []*models.ServiceObject
	var err error
	auth := common.GetAuthContext(c)
	listSpec := common.GetListSpec(c)
	listSpec.Auth = auth
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListServiceObject(tx, listSpec)
			return err
		}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"service-objects": result,
	})
}
