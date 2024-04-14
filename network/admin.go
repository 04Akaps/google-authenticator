package network

import (
	"net/http"

	"demo-scrapping/types"

	"github.com/gin-gonic/gin"
)

type admin struct {
	network *Network
}

func newAdmin(network *Network) {
	a := &admin{network: network}

	basePath := "/admin"

	network.register(basePath+"/add", POST, network.verifyAuth(), a.add)
	network.register(basePath+"/update", PUT, a.update)

	network.register(basePath+"/delete", DELETE, a.delete)
	network.register(basePath+"/view", POST, a.view)
	network.register(basePath+"/view-all", POST, a.viewAll)
}

func (a *admin) add(c *gin.Context) {
	// var req types.AddReq

	// if err := c.ShouldBindJSON(&req); err != nil {
	// 	res(c, http.StatusUnprocessableEntity, nil, err.Error())
	// } else if err := a.network.service.Add(req.URL, req.CardSelector, req.InnerSelector, req.Tag); err != nil {
	// 	res(c, http.StatusInternalServerError, nil, err.Error())
	// } else {
	// 	res(c, http.StatusOK, "Success,", "Success")
	// }

	res(c, http.StatusOK, "Success,", "Success")
}

func (a *admin) update(c *gin.Context) {
	var req types.UpdateReq

	if err := c.ShouldBindJSON(&req); err != nil {
		res(c, http.StatusUnprocessableEntity, nil, err.Error())
	} else if err := a.network.service.Update(req.URL, req.CardSelector, req.InnerSelector, req.Tag); err != nil {
		res(c, http.StatusInternalServerError, nil, err.Error())
	} else {
		res(c, http.StatusOK, "Success,", "Success")
	}
}

func (a *admin) view(c *gin.Context) {
	var req types.ViewReq

	if err := c.ShouldBindJSON(&req); err != nil {
		res(c, http.StatusUnprocessableEntity, nil, err.Error())
	} else if response, err := a.network.service.View(req.URL); err != nil {
		res(c, http.StatusInternalServerError, nil, err.Error())
	} else {
		res(c, http.StatusOK, response, "Success")
	}
}

func (a *admin) viewAll(c *gin.Context) {
	if response, err := a.network.service.ViewAll(); err != nil {
		res(c, http.StatusInternalServerError, nil, err.Error())
	} else {
		res(c, http.StatusOK, response, "Success")
	}
}

func (a *admin) delete(c *gin.Context) {
	var req types.DeleteReq

	if err := c.ShouldBindJSON(&req); err != nil {
		res(c, http.StatusUnprocessableEntity, nil, err.Error())
	} else if err := a.network.service.Delete(req.URL); err != nil {
		res(c, http.StatusInternalServerError, nil, err.Error())
	} else {
		res(c, http.StatusOK, "Success,", "Success")
	}
}
