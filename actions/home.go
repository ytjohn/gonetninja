package actions

import (
	"github.com/gofrs/uuid"
	"gonetninja/models"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/pkg/errors"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("home/index.plush.html"))
}

func NetActivityHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("home/netactivity.plush.html"))
}

func NetListHandler(c buffalo.Context) error {
	//tx := c.Value("tx").(*pop.Connection)
	nets := models.Nets{}
	if err := models.DB.Order("name").All(&nets); err != nil {
		return errors.WithStack(err)
	}
	//if err := tx.Order("name").All(&nets); err != nil {
	//	return errors.WithStack(err)
	//}
	c.Set("nets", nets)
	return c.Render(http.StatusOK, r.HTML("home/netlist.plush.html"))
}

func NetHandler(c buffalo.Context) error {
	//tx := c.Value("tx").(*pop.Connection)
	net := models.Net{}
	//err := models.DB.Find(&net, c.id)
	//if err := tx.Find(&net, c.Param("id")); err != nil {
	//	return errors.WithStack(err)
	//}
	if err := models.DB.Find(&net, c.Param("id")); err != nil {
		return errors.WithStack(err)
	}

	c.Set("net", net)
	return c.Render(http.StatusOK, r.HTML("home/netview.plush.html"))
}

func NewNetFormHandler(c buffalo.Context) error {
	c.Set("net", models.Net{})
	return c.Render(http.StatusOK, r.HTML("home/netnew.plush.html"))
}

func CreateNetHandler(c buffalo.Context) error {
	net := &models.Net{}
	if err := c.Bind(net); err != nil {
		return err
	}

	newId, err := uuid.NewV1()
	if err != nil {
		return err
	}
	net.ID = newId

	// Validate the data from the html form
	verrs, err := models.DB.ValidateAndCreate(net)
	if err != nil {
		return errors.WithStack(err)
	}
	if verrs.HasAny() {
		c.Set("net", net)
		// Make the errors available inside the html template
		c.Set("errors", verrs)
		return c.Render(422, r.HTML("home/netnew.plush.html"))
	}
	c.Flash().Add("success", "Net was created successfully")
	return c.Redirect(302, "/nets/%s", net.ID)
}
