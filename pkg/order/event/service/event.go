package service

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"net/smtp"
	"sync"
	"text/template"

	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

// EventAdminService ...
type EventAdminService struct {
	EventDAO model.EventDAO
	User     model.UserDAO
}

// NewEventAdminService ...
func NewEventAdminService(d *model.CommonDAO) model.EventAdminService {
	return &EventAdminService{
		EventDAO: d.Event,
		User:     d.User,
	}
}

// Create ...
func (d *EventAdminService) Create(ctx context.Context, body model.EventBody) (doc model.EventAdminResponse, err error) {
	payload := body.NewEventRaw()
	err = d.EventDAO.InsertOne(ctx, payload)
	res := payload.EventGetAdminResponse()
	return res, err
}

func (d *EventAdminService) GetDetail(ctx context.Context, event model.EventRaw) model.EventAdminResponse {
	return event.EventGetAdminResponse()
}

// ListAll ...
func (d *EventAdminService) ListAll(ctx context.Context, q model.CommonQuery) ([]model.EventAdminResponse, int64) {
	var (
		wg    sync.WaitGroup
		cond  = bson.M{}
		total int64
		res   = make([]model.EventAdminResponse, 0)
	)

	q.AssignActive(&cond)

	wg.Add(2)
	go func() {
		defer wg.Done()
		total = d.EventDAO.CountByCondition(ctx, cond)
	}()

	go func() {
		defer wg.Done()
		events, _ := d.EventDAO.FindByCondition(ctx, cond, q.GetFindOptsUsingPage())
		for _, value := range events {
			temp := value.EventGetAdminResponse()
			res = append(res, temp)
		}
	}()

	wg.Wait()
	return res, total
}

// Update ....
func (d *EventAdminService) Update(ctx context.Context, Event model.EventRaw, body model.EventBody) (doc model.EventAdminResponse, err error) {
	payload := body.NewEventRaw()

	// assgin
	Event.Name = payload.Name
	Event.Desc = payload.Desc
	Event.UpdatedAt = payload.UpdatedAt

	err = d.EventDAO.UpdateByID(ctx, Event.ID, bson.M{"$set": Event})
	if err != nil {
		return doc, errors.New(locale.EventKeyCanNotUpdate)
	}

	event, _ := d.EventDAO.FindOneByCondition(ctx, bson.M{"_id": Event.ID})
	res := event.EventGetAdminResponse()
	return res, err
}

// FindByID ...
func (d *EventAdminService) FindByID(ctx context.Context, id model.primitive.ObjectID) (event model.EventRaw, err error) {
	return d.EventDAO.FindOneByCondition(ctx, bson.M{"_id": id})
}

func (d *EventAdminService) ChangeStatus(ctx context.Context, event model.EventRaw) (err error) {

	// //check trang thai hien tai cua event
	// if event.Active {
	// 	return fmt.Errorf("Trang thai da active")
	// }

	payload := bson.M{
		"$set": bson.M{
			"active": !event.Active,
		},
	}

	err = d.EventDAO.UpdateByID(ctx, event.ID, payload)
	if err != nil {
		return errors.New(locale.EventKeyCanNotUpdate)
	}
	// cond := bson.M{}
	// var to = make([]string, 0)
	// // get email
	// users, err := d.User.FindByCondition(ctx, cond)
	// if err != nil {
	// 	return fmt.Errorf("Da xay ra loi")
	// }
	// for _, users := range users {
	// 	to = append(to, users.Email)
	// }
	// fmt.Println("to", to)

	// // send email cho ng dung
	// d.sendEmailForUser(event, to)

	return nil

}

func (d *EventAdminService) SendEmail(ctx context.Context, event model.EventRaw) (err error) {
	cond := bson.M{}
	var to = make([]string, 0)
	// get email
	users, err := d.User.FindByCondition(ctx, cond)
	if err != nil {
		return fmt.Errorf("Da xay ra loi")
	}
	for _, users := range users {
		to = append(to, users.Email)
	}
	// send email cho ng dung
	d.sendEmailForUser(event, to)
	return nil
}

func (d *EventAdminService) sendEmailForUser(event model.EventRaw, to []string) {
	from := "nopromise1999@gmail.com"
	pass := "nitranhngao@81299@"

	// to := []string{
	// 	args.Email,
	// }

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Authentication.
	auth := smtp.PlainAuth("", from, pass, smtpHost)

	t, err := template.ParseFiles("template_event.html")

	if err != nil {
		log.Printf(" error: %s", err)
	}
	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: [CoffeeShop] Event \n%s\n\n", mimeHeaders)))
	t.Execute(&body, struct {
		Name string
		Desc string
	}{
		Name: event.Name,
		Desc: event.Desc,
	})

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	fmt.Println(err)
	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
}

func (d *EventAdminService) DeleteEvent(ctx context.Context, c model.EventRaw) error {
	err := d.EventDAO.DeleteByID(ctx, c.ID)
	return err
}
