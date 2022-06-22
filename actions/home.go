package actions

import (
	"github.com/sirupsen/logrus"
	"gonetninja/models"
	"net/http"
	"sort"
	"time"

	"github.com/gofrs/uuid"

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
	net := models.Net{}
	if err := models.DB.Find(&net, c.Param("id")); err != nil {
		return errors.WithStack(err)
	}

	activities := models.Activities{}
	query := models.DB.Where("net = (?)", net.ID)
	if err := query.Order("time_at desc").All(&activities); err != nil {
		return errors.WithStack(err)
	}
	c.Set("opened", GetOpen(net.ID))
	c.Set("closed", GetClose(net.ID))
	c.Set("netcontrols", NetControls(net.ID))
	c.Set("participants", NetParticipants(net.ID))
	c.Set("net", net)
	c.Set("activities", activities)
	return c.Render(http.StatusOK, r.HTML("home/netview.plush.html"))
}

func NetControls(id uuid.UUID) []string {
	//err = models.DB.Select("name").All(&users)
	set := make(map[string]struct{})

	activities := models.Activities{}
	query := models.DB.Where("net = (?) AND action = (?)", id, "netcontrol")
	query.Order("name desc").All(&activities)

	for _, a := range activities {
		logrus.Info("checking participant ?", a.Name)
		_, isPresent := set[a.Name]
		if !isPresent {
			logrus.Info("found new participant ?", a.Name)
			set[a.Name] = struct{}{}
		}
	}
	keys := make([]string, 0, len(set))
	for k := range set {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	logrus.Info(keys)
	return keys
}

func NetParticipants(id uuid.UUID) []string {
	//err = models.DB.Select("name").All(&users)
	set := make(map[string]struct{})

	activities := models.Activities{}
	//query := models.DB.Select("name")
	query := models.DB.Where("net = (?)", id)
	query.Order("name desc").All(&activities)
	//query.Order("time_at asc").All(&activities)

	for _, a := range activities {
		logrus.Info("checking participant ?", a.Name)
		_, isPresent := set[a.Name]
		if !isPresent {
			logrus.Info("found new participant ?", a.Name)
			set[a.Name] = struct{}{}
		}
	}
	keys := make([]string, 0, len(set))
	for k := range set {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	logrus.Info(keys)
	return keys
}

func GetOpen(id uuid.UUID) time.Time {
	activity := models.Activity{}
	query := models.DB.Where("net = (?) AND action = (?)", id, "open")
	query.Order("time_at asc").First(&activity)
	// SELECT activity.action, activity.created_at, activity.description, activity.id, activity.name,
	// 	activity.net, activity.time_at, activity.updated_at FROM activity AS activity
	// WHERE net = (?) AND action = (?)
	// ORDER BY time_at asc LIMIT 1 $1=786829ca-f1c3-11ec-a9d5-ce5390833a0f $2=open
	return activity.TimeAt
}

func GetClose(id uuid.UUID) time.Time {
	activity := models.Activity{}
	query := models.DB.Where("net = (?) AND action = (?)", id, "close")
	query.Order("time_at desc").First(&activity)
	// SELECT activity.action, activity.created_at, activity.description, activity.id, activity.name,
	// 	activity.net, activity.time_at, activity.updated_at FROM activity AS activity
	// WHERE net = (?) AND action = (?)
	// ORDER BY time_at asc LIMIT 1 $1=786829ca-f1c3-11ec-a9d5-ce5390833a0f $2=open
	return activity.TimeAt
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
