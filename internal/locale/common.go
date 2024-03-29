package locale

import (
	"log"

	"github.com/magiconair/properties"
)

// Keys
const (
	CommonKeySuccess            = "success"
	CommonKeyBadRequest         = "badRequest"
	CommonKeyUnauthorized       = "unauthorized"
	CommonKeyNotFound           = "notFound"
	CommonKeyInvalidChecksum    = "invalidChecksum"
	CommonKeyErrorWhenHandle    = "errorWhenHandle"
	CommonKeyServerIsProcessing = "serverIsProcessing"
	CommonKeyInvalidPhoto       = "invalidPhoto"
	CommonKeyNameIsRequired     = "nameIsRequired"
	CommonKeyPhoneExisted       = "phoneExisted"
	CommonKeyEmailExisted       = "emailExisted"

	CommonKeyContactNameIsRequired    = "contactNameIsRequired"
	CommonKeyContactAddressIsRequired = "contactAddressIsRequired"
	CommonKeyReadFileFail             = "readFileFail"

	CommonKeyPhoneIsRequired      = "phoneIsRequired"
	CommonKeyUIDIsRequired        = "uIDIsRequired"
	CommonKeyPhoneNotMatch        = "phoneNotMatch"
	CommonKeyPhoneNoExisted       = "phoneNoExisted"
	CommonKeyUserUnActive         = "userUnActive"
	CommonKeyPermissionIsRequired = "permissionIsRequired"

	CommonKeyUsernameIsRequired    = "UsernameIsRequired"
	CommonKeyPasswordRequired      = "PasswordRequired"
	AuthKeyUserIsRequired          = "UserIsRequired"
	AuthKeyUserInvalid             = "UserInvalid"
	AuthKeyAccountUnActive         = "AccountUnActive"
	AuthKeyUserUnActive            = "UserUnActive"
	CommonKeyCanNotCreateUser      = "CanNotCreateUser"
	UserKeyCanNotChangeStatus      = "UserCanNotChangeStatus"
	OrderKeyDrinkCanNotFind        = "OrderDrinkCanNotFind"
	OrderKeyCanNotCreateOrder      = "OrderCanNotCreateOrder"
	OrderKeyCanNotFindUserByUserID = "OrderCanNotFindUserByUserID"
	CategoryKeyCanNotUpdate        = "CategoryCanNotUpdate"

	CommonKeyIDMongoInvalid = "CommonIDMongoInvalid"

	CommonNoPermission            = "CommonNoPermission"
	CommonKeyStaffDeactive        = "CommonStaffDeactive"
	CommonyKeyUserNameIsExisted   = "UserNameIsExisted"
	CommonKeyStaffIsDeleted       = "StaffIsDeleted"
	StaffIsNotExisted             = "StaffIsNotExisted"
	PasswordIsIncorrect           = "PasswordIsIncorrect"
	UserNameOrPasswordIsIncorrect = "UserNameOrPasswordIsIncorrect"
	UserIsNotExisted              = "userIsNotExisted"
	CurrentPointIsNotEnough       = "current point is not enough"
	UpdatePointFailed             = "update point failed"
	OrderIsEmpty                  = "orderIsEmpty"
	PointIsRequired               = "pointIsRequired"
	OrderCanNotCancel             = "orderNotCancel"
	OrderStatusIsInvalid          = "OrderStatusIsInvalid"
	OrderStatusCanNotUpdate       = "order status can not update"
	CommonKeyEmailInvalid         = "emailInvalid"
	CommonKeyEmailIsRequired      = "emailIsRequired"
	CommonyKeyEmailIsExisted      = "EmailIsExisted"
	CodeIsRequired                = "codeIsRequired"
)

type (
	commonLang struct {
		Success            string `properties:"success"`
		BadRequest         string `properties:"badRequest"`
		Unauthorized       string `properties:"unauthorized"`
		NotFound           string `properties:"notFound"`
		InvalidChecksum    string `properties:"invalidChecksum"`
		ErrorWhenHandle    string `properties:"errorWhenHandle"`
		ServerIsProcessing string `properties:"serverIsProcessing"`
		InvalidPhoto       string `properties:"invalidPhoto"`
		NameIsRequired     string `properties:"nameIsRequired"`
		PhoneExisted       string `properties:"phoneExisted"`
		EmailExisted       string `properties:"emailExisted"`

		ContactNameIsRequired    string `properties:"contactNameIsRequired"`
		ContactPhoneIsRequired   string `properties:"contactPhoneIsRequired"`
		ContactAddressIsRequired string `properties:"contactAddressIsRequired"`
		ReadFileFail             string `properties:"readFileFail"`

		PhoneIsRequired             string `properties:"phoneIsRequired"`
		UIDIsRequired               string `properties:"uIDIsRequired"`
		PhoneNotMatch               string `properties:"phoneNotMatch"`
		PhoneNoExisted              string `properties:"phoneNoExisted"`
		UserUnActive                string `properties:"userUnActive"`
		UsernameIsRequired          string `properties:"usernameIsRequired"`
		PasswordRequired            string `properties:"passwordRequired"`
		PermissionIsRequired        string `properties:"permissionIsRequired"`
		UserIsRequired              string `properties:"userIsRequired"`
		UserInvalid                 string `properties:"userInvalid"`
		AccountUnActive             string `properties:"accountUnActive"`
		CanNotCreateUser            string `properties:"canNotCreateUser"`
		UserCanNotChangeStatus      string `properties:"userCanNotChangeStatus"`
		OrderDrinkCanNotFind        string `properties:"orderDrinkCanNotFind"`
		OrderCanNotCreateOrder      string `properties:"orderCanNotCreateOrder"`
		OrderCanNotFindUserByUserID string `properties:"orderCanNotFindUserByUserID"`
		CategoryCanNotUpdate        string `properties:"categoryCanNotUpdate"`

		CommonIDMongoInvalid          string `properties:"commonIDMongoInvalid"`
		CommonNoPermission            string `properties:"commonNoPermission"`
		CommonStaffDeactive           string `properties:"commonStaffDeactive"`
		UserNameIsExisted             string `properties:"userNameIsExisted"`
		CommonKeyStaffIsDeleted       string `properties:"commonKeyStaffIsDeleted"`
		StaffIsNotExisted             string `properties:"StaffIsNotExisted"`
		PasswordIsIncorrect           string `properties:"passwordIsIncorrect"`
		UserNameOrPasswordIsIncorrect string `properties:"UserNameOrPasswordIsIncorrect"`
		UserIsNotExisted              string `properties:"UserIsNotExisted"`
		CurrentPointIsNotEnough       string `properties:"CurrentPointIsNotEnough"`
		UpdatePointFailed             string `properties:"UpdatePointFailed"`
		OrderIsEmpty                  string `properties:"OrderIsEmpty"`
		OrderCanNotCancel             string `properties:"OrderCanNotCancel"`
		OrderStatusIsInvalid          string `properties:"OrderStatusIsInvalid"`
		OrderStatusCanNotUpdate       string `properties:"OrderStatusCanNotUpdate"`
		CodeIsRequired                string `properties:"CodeIsRequired"`
	}
)

var (
	commonVi commonLang
)

func init() {
	// Load propertie
	p2 := properties.MustLoadFile(getLocalePath()+"/properties/common.properties", properties.UTF8)
	if err := p2.Decode(&commonVi); err != nil {
		log.Fatal(err)
	}
}

func commonLoadLocales() (response []Locale) {
	// 1-99
	response = []Locale{
		{
			Key: CommonKeySuccess,
			Message: &Message{

				Vi: commonVi.Success,
			},
			Code: 1,
		},
		{
			Key: CommonKeyBadRequest,
			Message: &Message{

				Vi: commonVi.BadRequest,
			},
			Code: 2,
		},
		{
			Key: CommonKeyUnauthorized,
			Message: &Message{

				Vi: commonVi.Unauthorized,
			},
			Code: 3,
		},
		{
			Key: CommonKeyNotFound,
			Message: &Message{

				Vi: commonVi.NotFound,
			},
			Code: 4,
		},
		{
			Key: CommonKeyInvalidChecksum,
			Message: &Message{

				Vi: commonVi.InvalidChecksum,
			},
			Code: 5,
		},
		{
			Key: CommonKeyErrorWhenHandle,
			Message: &Message{

				Vi: commonVi.ErrorWhenHandle,
			},
			Code: 6,
		},
		{
			Key: CommonKeyServerIsProcessing,
			Message: &Message{

				Vi: commonVi.ServerIsProcessing,
			},
			Code: 7,
		},
		{
			Key: CommonKeyInvalidPhoto,
			Message: &Message{

				Vi: commonVi.InvalidPhoto,
			},
			Code: 8,
		},
		{
			Key: CommonKeyNameIsRequired,
			Message: &Message{

				Vi: commonVi.NameIsRequired,
			},
			Code: 9,
		},
		{
			Key: CommonKeyPhoneExisted,
			Message: &Message{

				Vi: commonVi.PhoneExisted,
			},
			Code: 10,
		},
		{
			Key: CommonKeyEmailExisted,
			Message: &Message{

				Vi: commonVi.EmailExisted,
			},
			Code: 11,
		},

		{
			Key: CommonKeyContactNameIsRequired,
			Message: &Message{

				Vi: commonVi.ContactNameIsRequired,
			},
			Code: 16,
		},

		{
			Key: CommonKeyContactAddressIsRequired,
			Message: &Message{

				Vi: commonVi.ContactAddressIsRequired,
			},
			Code: 18,
		},
		{
			Key: CommonKeyReadFileFail,
			Message: &Message{

				Vi: commonVi.ReadFileFail,
			},
			Code: 19,
		},

		{
			Key: CommonKeyPhoneNotMatch,
			Message: &Message{

				Vi: commonVi.PhoneNotMatch,
			},
			Code: 28,
		},
		{
			Key: CommonKeyPhoneNoExisted,
			Message: &Message{

				Vi: commonVi.PhoneNoExisted,
			},
			Code: 29,
		},
		{
			Key: CommonKeyUserUnActive,
			Message: &Message{

				Vi: commonVi.UserUnActive,
			},
			Code: 30,
		},
		{
			Key: CommonKeyUsernameIsRequired,
			Message: &Message{

				Vi: commonVi.UsernameIsRequired,
			},
			Code: 31,
		},
		{
			Key: CommonKeyPasswordRequired,
			Message: &Message{

				Vi: commonVi.PasswordRequired,
			},
			Code: 32,
		},
		{
			Key: CommonKeyPermissionIsRequired,
			Message: &Message{

				Vi: commonVi.PermissionIsRequired,
			},
			Code: 33,
		},
		{
			Key: AuthKeyUserIsRequired,
			Message: &Message{

				Vi: commonVi.UserIsRequired,
			},
			Code: 34,
		},
		{
			Key: AuthKeyUserInvalid,
			Message: &Message{

				Vi: commonVi.UserInvalid,
			},
			Code: 35,
		},
		{
			Key: AuthKeyAccountUnActive,
			Message: &Message{

				Vi: commonVi.AccountUnActive,
			},
			Code: 36,
		},
		{
			Key: CommonKeyCanNotCreateUser,
			Message: &Message{

				Vi: commonVi.CanNotCreateUser,
			},
			Code: 37,
		},
		{
			Key: UserKeyCanNotChangeStatus,
			Message: &Message{

				Vi: commonVi.UserCanNotChangeStatus,
			},
			Code: 38,
		},

		{
			Key: OrderKeyDrinkCanNotFind,
			Message: &Message{

				Vi: commonVi.OrderDrinkCanNotFind,
			},
			Code: 44,
		},
		{
			Key: OrderKeyCanNotCreateOrder,
			Message: &Message{

				Vi: commonVi.OrderCanNotCreateOrder,
			},
			Code: 45,
		},
		{
			Key: OrderKeyCanNotFindUserByUserID,
			Message: &Message{

				Vi: commonVi.OrderCanNotFindUserByUserID,
			},
			Code: 46,
		},
		{
			Key: CategoryKeyCanNotUpdate,
			Message: &Message{

				Vi: commonVi.CategoryCanNotUpdate,
			},
			Code: 47,
		},

		{
			Key: CommonKeyIDMongoInvalid,
			Message: &Message{

				Vi: commonVi.CommonIDMongoInvalid,
			},
			Code: 48,
		},

		{
			Key: CommonNoPermission,
			Message: &Message{

				Vi: commonVi.CommonNoPermission,
			},
			Code: 49,
		},

		{
			Key: CommonKeyStaffDeactive,
			Message: &Message{

				Vi: commonVi.CommonStaffDeactive,
			},
			Code: 50,
		},

		{
			Key: CommonyKeyUserNameIsExisted,
			Message: &Message{
				Vi: commonVi.UserNameIsExisted,
			},
			Code: 51,
		},

		{
			Key: CommonKeyStaffIsDeleted,
			Message: &Message{
				Vi: commonVi.CommonKeyStaffIsDeleted,
			},
			Code: 52,
		},

		{
			Key: StaffIsNotExisted,
			Message: &Message{
				Vi: commonVi.StaffIsNotExisted,
			},
			Code: 53,
		},

		{
			Key: PasswordIsIncorrect,
			Message: &Message{
				Vi: commonVi.PasswordIsIncorrect,
			},
			Code: 54,
		},

		{
			Key: UserNameOrPasswordIsIncorrect,
			Message: &Message{
				Vi: commonVi.UserNameOrPasswordIsIncorrect,
			},
			Code: 55,
		},

		{
			Key: UserIsNotExisted,
			Message: &Message{
				Vi: commonVi.UserIsNotExisted,
			},
			Code: 56,
		},
		{
			Key: CurrentPointIsNotEnough,
			Message: &Message{
				Vi: commonVi.CurrentPointIsNotEnough,
			},
			Code: 57,
		},

		{
			Key: UpdatePointFailed,
			Message: &Message{
				Vi: commonVi.UpdatePointFailed,
			},
			Code: 58,
		},

		{
			Key: OrderIsEmpty,
			Message: &Message{
				Vi: commonVi.OrderIsEmpty,
			},
			Code: 59,
		},

		{
			Key: OrderCanNotCancel,
			Message: &Message{
				Vi: commonVi.OrderCanNotCancel,
			},
			Code: 60,
		},

		{
			Key: OrderStatusIsInvalid,
			Message: &Message{
				Vi: commonVi.OrderStatusIsInvalid,
			},
			Code: 61,
		},
		{
			Key: OrderStatusCanNotUpdate,
			Message: &Message{
				Vi: commonVi.OrderStatusCanNotUpdate,
			},
			Code: 62,
		},

		{
			Key: CodeIsRequired,
			Message: &Message{
				Vi: commonVi.CodeIsRequired,
			},
			Code: 63,
		},
	}
	return response
}
