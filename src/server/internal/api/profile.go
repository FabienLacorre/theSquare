// File: profile.go
// File Created: 07 Mar 2019 16:10
// By Maxence Moutoussamy <maxence.moutoussamy1@gmail.com>

package api

import (
	"github.com/emicklei/go-restful"
)

type ProfileService struct {
}

func (s *ProfileService) Register(root string) *restful.WebService {
	ws := new(restful.WebService)
	ws.Path(root + "/profile")
	ws.Route(ws.GET("/{id}").To(s.getByID).
		Doc("Get the profile that is corresponding to the given ID").
		Param(ws.PathParameter("id", "ID that correspond to the profile").DataType("int"))).
		Produces(restful.MIME_JSON)
	return ws
}

func (s *ProfileService) getByID(req *restful.Request, resp *restful.Response) {
	resp.Write([]byte("Hello world!"))
}
