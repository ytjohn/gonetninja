package actions

import (
	"gonetninja/models"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/sirupsen/logrus"

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
	c.Set("net", net)
	//activities := models.Activities{}
	//query := models.DB.Where("net = (?)", net.ID)
	//net.PlannedStart.S
	//net.PlannedStart.Fo
	//if err := query.Order("time_at desc").All(&activities); err != nil {
	//	return errors.WithStack(err)
	//}
	c.Set("quicknet", models.Quicknet{})
	//c.Set("opened", GetOpen(net.ID))
	//c.Set("closed", GetClose(net.ID))
	//c.Set("netcontrols", NetControls(net.ID))
	//c.Set("participants", NetParticipants(net.ID))
	//c.Set("net", net)
	//c.Set("activities", activities)
	c = _LearnNet(c, net)
	return c.Render(http.StatusOK, r.HTML("home/netedit.plush.html"))
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
	//c.Set("net", models.Net{})
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
	return c.Redirect(302, "/nets/%s?editmode=true", net.ID)
}

func _LearnNet(c buffalo.Context, net models.Net) buffalo.Context {
	activities := models.Activities{}
	query := models.DB.Where("net = (?)", net.ID)
	query.Order("time_at desc").All(&activities)
	c.Set("activities", activities)
	c.Set("opened", GetOpen(net.ID))
	c.Set("closed", GetClose(net.ID))
	c.Set("netcontrols", NetControls(net.ID))
	c.Set("participants", NetParticipants(net.ID))
	return c
}

func QuickNetHandler(c buffalo.Context) error {
	net := models.Net{}
	if err := models.DB.Find(&net, c.Param("id")); err != nil {
		return errors.WithStack(err)
	}
	c.Set("net", net)
	quicknet := &models.Quicknet{}
	c.Set("quicknet", quicknet)
	if err := c.Bind(quicknet); err != nil {
		return err
	}

	// Validate the data from the html form
	verrs, err := quicknet.Validate()
	if err != nil {
		return errors.WithStack(err)
	}
	if verrs.HasAny() {
		c = _LearnNet(c, net)
		logrus.Info("quicknet has verrs")

		// Make the errors available inside the html template
		c.Set("errors", verrs)
		//c.Flash().Add("alert", "verrs")
		return c.Render(422, r.HTML("home/netedit.plush.html"))
	}
	logrus.Info("quicknet  has no verrs, continuing")

	// Assume Net Control
	_ = models.DB.Create(&models.Activity{
		//ID:          ncr_id,
		Net:         net.ID,
		Action:      "netcontrol",
		Name:        quicknet.NetControl,
		TimeAt:      quicknet.Opened.Add(-time.Minute * 5),
		Description: "Assumed Net Control",
	})

	// Open the net
	_ = models.DB.Create(&models.Activity{
		Net:         net.ID,
		Action:      "open",
		Name:        quicknet.NetControl,
		TimeAt:      quicknet.Opened,
		Description: "Opened the net",
	})

	// Close the net if set
	if !quicknet.Closed.IsZero() {
		// In validate, we forced quicknet.Closed to be at least 5
		// minutes after the quicknet.Open
		_ = models.DB.Create(&models.Activity{
			Net:         net.ID,
			Action:      "close",
			Name:        quicknet.NetControl,
			TimeAt:      quicknet.Closed,
			Description: "Closed the net",
		})
	}

	// now do the early checkins
	if quicknet.EarlyCheckins != "" {
		early_names := ParseNames(quicknet.EarlyCheckins)
		early_time := quicknet.Opened.Add(-time.Second * 60)
		for _, n := range early_names {
			_ = models.DB.Create(&models.Activity{
				Net:         net.ID,
				Action:      "checkin",
				Name:        n,
				TimeAt:      early_time,
				Description: "Early checkin",
			})
		}
	}

	if quicknet.RegularCheckins != "" {
		regular_names := ParseNames(quicknet.RegularCheckins)
		regular_time := quicknet.Opened.Add(time.Second * 30)
		for _, n := range regular_names {
			_ = models.DB.Create(&models.Activity{
				Net:         net.ID,
				Action:      "checkin",
				Name:        n,
				TimeAt:      regular_time,
				Description: "Regular checkin",
			})
		}
	}

	return c.Redirect(302, "/nets/%s?editmode=true", net.ID)
}

func ParseNames(raw string) []string {

	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fields := strings.FieldsFunc(raw, f)
	set := make(map[string]struct{})
	for _, f := range fields {
		_, isPresent := set[f]
		if !isPresent {
			// If the name is just a number, skip
			// ex: "1. John" => ["1", "John"]
			_, err := strconv.Atoi(f)
			if err != nil {
				set[f] = struct{}{}
			}
		}
	}
	keys := make([]string, 0, len(set))
	for k := range set {
		keys = append(keys, k)
	}
	return keys
}
