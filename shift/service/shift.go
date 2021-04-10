package service

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

// ShiftAdminService ...
type ShiftAdminService struct {
	ShiftDAO model.ShiftDAO
	StaffDAO model.StaffDAO
}

// NewShiftAdminService ...
func NewShiftAdminService(d *model.CommonDAO) model.ShiftAdminService {
	return &ShiftAdminService{
		ShiftDAO: d.Shift,
		StaffDAO: d.Staff,
	}
}

// Create ...
func (d *ShiftAdminService) Create(ctx context.Context, body model.ShiftBody, staff model.StaffRaw) (doc model.ShiftResponse, err error) {

	payload := body.NewShiftRaw(staff.ID)
	err = d.ShiftDAO.InsertOne(ctx, payload)

	if err != nil {
		return doc, err
	}

	// convert
	staffRaw, _ := d.StaffDAO.FindByID(ctx, staff.ID)
	staffInfo := staffRaw.GetStaffInfo()
	doc = payload.GetResponse(staffInfo)
	return doc, nil
}

// ListAll ...
func (d *ShiftAdminService) ListAll(ctx context.Context, q model.CommonQuery) ([]model.ShiftResponse, int64) {
	var (
		wg    sync.WaitGroup
		cond  = bson.M{}
		total int64
		res   = make([]model.ShiftResponse, 0)
	)

	q.AssignKeyword(&cond)
	q.AssignStaff(&cond)
	q.AssignStartAtAndEndAt(&cond)

	total = d.ShiftDAO.CountByCondition(ctx, cond)
	shifts, _ := d.ShiftDAO.FindByCondition(ctx, cond)
	if len(shifts) > 0 {
		wg.Add(len(shifts))
		for index, shift := range shifts {
			res = make([]model.ShiftResponse, len(shifts))
			go func(s model.ShiftRaw, i int) {
				defer wg.Done()
				staffRaw, _ := d.StaffDAO.FindByID(ctx, q.Staff)
				staffInfo := staffRaw.GetStaffInfo()
				doc := s.GetResponse(staffInfo)
				res[i] = doc
			}(shift, index)
		}
		wg.Wait()
	}

	return res, total
}

// Update ....
func (d *ShiftAdminService) Update(ctx context.Context, Shift model.ShiftRaw, body model.ShiftBody, staff model.StaffRaw) (doc model.ShiftResponse, err error) {
	payload := body.NewShiftRaw(staff.ID)

	// assgin
	Shift.Name = payload.Name
	Shift.UpdatedAt = payload.UpdatedAt
	Shift.Date = payload.Date

	err = d.ShiftDAO.UpdateByID(ctx, Shift.ID, bson.M{"$set": Shift})
	if err != nil {
		return doc, errors.New("khong the cap nhat ca lam viec")
	}
	staffRaw, _ := d.StaffDAO.FindByID(ctx, staff.ID)
	staffInfo := staffRaw.GetStaffInfo()
	doc = Shift.GetResponse(staffInfo)
	return doc, nil
}

// FindByID ...
func (d *ShiftAdminService) FindByID(ctx context.Context, id model.AppID) (Shift model.ShiftRaw, err error) {
	return d.ShiftDAO.FindOneByCondition(ctx, bson.M{"_id": id})
}

func (d *ShiftAdminService) AcceptShiftByAdmin(ctx context.Context, raw model.ShiftRaw) (bool, error) {
	check := !raw.IsCheck
	payload := bson.M{
		"isCheck":   check,
		"updatedAt": time.Now(),
	}
	err := d.ShiftDAO.UpdateByID(ctx, raw.ID, bson.M{"$set": payload})
	if err != nil {
		return raw.IsCheck, errors.New("khong the thay doi trang thai ca lam viec")
	}
	return check, nil
}
