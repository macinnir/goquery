package testassets

import (
	"database/sql"
	"encoding/json"
	"fmt"

	goquery "github.com/macinnir/query"
	"gopkg.in/guregu/null.v3"
)

const (

	// User_SchemaName is the name of the schema group this model is in
	User_SchemaName = "core"

	// User_TableName is the name of the table
	User_TableName goquery.TableName = "User"

	// Columns

	User_Column_APIKeyID                goquery.Column = "APIKeyID"
	User_Column_AccountID               goquery.Column = "AccountID"
	User_Column_AccountRole             goquery.Column = "AccountRole"
	User_Column_AppTheme                goquery.Column = "AppTheme"
	User_Column_Coins                   goquery.Column = "Coins"
	User_Column_DailyStudiedTermsGoal   goquery.Column = "DailyStudiedTermsGoal"
	User_Column_DailyTestedTermsGoal    goquery.Column = "DailyTestedTermsGoal"
	User_Column_DateActivated           goquery.Column = "DateActivated"
	User_Column_DateCreated             goquery.Column = "DateCreated"
	User_Column_DateEmailValidated      goquery.Column = "DateEmailValidated"
	User_Column_Email                   goquery.Column = "Email"
	User_Column_EnergyPoints            goquery.Column = "EnergyPoints"
	User_Column_FocusedCollectionID     goquery.Column = "FocusedCollectionID"
	User_Column_FollowerCount           goquery.Column = "FollowerCount"
	User_Column_FollowingCount          goquery.Column = "FollowingCount"
	User_Column_FullName                goquery.Column = "FullName"
	User_Column_GoogleSubID             goquery.Column = "GoogleSubID"
	User_Column_ID                      goquery.Column = "ID"
	User_Column_ImageID                 goquery.Column = "ImageID"
	User_Column_IsDeleted               goquery.Column = "IsDeleted"
	User_Column_IsDisabled              goquery.Column = "IsDisabled"
	User_Column_IsLocked                goquery.Column = "IsLocked"
	User_Column_IsPublic                goquery.Column = "IsPublic"
	User_Column_IsService               goquery.Column = "IsService"
	User_Column_IsShortcutsEnabled      goquery.Column = "IsShortcutsEnabled"
	User_Column_Language                goquery.Column = "Language"
	User_Column_LastLogin               goquery.Column = "LastLogin"
	User_Column_LastLoginAttempt        goquery.Column = "LastLoginAttempt"
	User_Column_LastLoginAttemptCounter goquery.Column = "LastLoginAttemptCounter"
	User_Column_LastNotificationDate    goquery.Column = "LastNotificationDate"
	User_Column_LastUpdated             goquery.Column = "LastUpdated"
	User_Column_Level                   goquery.Column = "Level"
	User_Column_Locale                  goquery.Column = "Locale"
	User_Column_ParentUserID            goquery.Column = "ParentUserID"
	User_Column_PermissionCount         goquery.Column = "PermissionCount"
	User_Column_Points                  goquery.Column = "Points"
	User_Column_PolicyCount             goquery.Column = "PolicyCount"
	User_Column_ProfileID               goquery.Column = "ProfileID"
	User_Column_ProfileImageID          goquery.Column = "ProfileImageID"
	User_Column_ProfileImageThumbURL    goquery.Column = "ProfileImageThumbURL"
	User_Column_ProfileImageURL         goquery.Column = "ProfileImageURL"
	User_Column_ReasonForLeaving        goquery.Column = "ReasonForLeaving"
	User_Column_RoleCount               goquery.Column = "RoleCount"
	User_Column_SMSPhoneNumber          goquery.Column = "SMSPhoneNumber"
	User_Column_SendSMSNotifications    goquery.Column = "SendSMSNotifications"
	User_Column_Streak                  goquery.Column = "Streak"
	User_Column_Timezone                goquery.Column = "Timezone"
	User_Column_TwoFactorEnabled        goquery.Column = "TwoFactorEnabled"
	User_Column_TwoFactorLastCompleted  goquery.Column = "TwoFactorLastCompleted"
	User_Column_UserGroupCount          goquery.Column = "UserGroupCount"
	User_Column_UserID                  goquery.Column = "UserID"
	User_Column_UserPrefix              goquery.Column = "UserPrefix"
	User_Column_Username                goquery.Column = "Username"
)

var (
	// User_Columns is a list of all the columns
	User_Columns = []goquery.Column{
		User_Column_APIKeyID,
		User_Column_AccountID,
		User_Column_AccountRole,
		User_Column_AppTheme,
		User_Column_Coins,
		User_Column_DailyStudiedTermsGoal,
		User_Column_DailyTestedTermsGoal,
		User_Column_DateActivated,
		User_Column_DateCreated,
		User_Column_DateEmailValidated,
		User_Column_Email,
		User_Column_EnergyPoints,
		User_Column_FocusedCollectionID,
		User_Column_FollowerCount,
		User_Column_FollowingCount,
		User_Column_FullName,
		User_Column_GoogleSubID,
		User_Column_ID,
		User_Column_ImageID,
		User_Column_IsDeleted,
		User_Column_IsDisabled,
		User_Column_IsLocked,
		User_Column_IsPublic,
		User_Column_IsService,
		User_Column_IsShortcutsEnabled,
		User_Column_Language,
		User_Column_LastLogin,
		User_Column_LastLoginAttempt,
		User_Column_LastLoginAttemptCounter,
		User_Column_LastNotificationDate,
		User_Column_LastUpdated,
		User_Column_Level,
		User_Column_Locale,
		User_Column_ParentUserID,
		User_Column_PermissionCount,
		User_Column_Points,
		User_Column_PolicyCount,
		User_Column_ProfileID,
		User_Column_ProfileImageID,
		User_Column_ProfileImageThumbURL,
		User_Column_ProfileImageURL,
		User_Column_ReasonForLeaving,
		User_Column_RoleCount,
		User_Column_SMSPhoneNumber,
		User_Column_SendSMSNotifications,
		User_Column_Streak,
		User_Column_Timezone,
		User_Column_TwoFactorEnabled,
		User_Column_TwoFactorLastCompleted,
		User_Column_UserGroupCount,
		User_Column_UserID,
		User_Column_UserPrefix,
		User_Column_Username,
	}

	// User_Column_Types maps columns to their string types
	User_Column_Types = map[goquery.Column]string{
		User_Column_APIKeyID:                "%d",
		User_Column_AccountID:               "%d",
		User_Column_AccountRole:             "%d",
		User_Column_AppTheme:                "%s",
		User_Column_Coins:                   "%d",
		User_Column_DailyStudiedTermsGoal:   "%d",
		User_Column_DailyTestedTermsGoal:    "%d",
		User_Column_DateActivated:           "%d",
		User_Column_DateCreated:             "%d",
		User_Column_DateEmailValidated:      "%d",
		User_Column_Email:                   "%s",
		User_Column_EnergyPoints:            "%d",
		User_Column_FocusedCollectionID:     "%d",
		User_Column_FollowerCount:           "%d",
		User_Column_FollowingCount:          "%d",
		User_Column_FullName:                "%s",
		User_Column_GoogleSubID:             "%s",
		User_Column_ID:                      "%s",
		User_Column_ImageID:                 "%d",
		User_Column_IsDeleted:               "%d",
		User_Column_IsDisabled:              "%d",
		User_Column_IsLocked:                "%d",
		User_Column_IsPublic:                "%d",
		User_Column_IsService:               "%d",
		User_Column_IsShortcutsEnabled:      "%d",
		User_Column_Language:                "%s",
		User_Column_LastLogin:               "%d",
		User_Column_LastLoginAttempt:        "%d",
		User_Column_LastLoginAttemptCounter: "%d",
		User_Column_LastNotificationDate:    "%d",
		User_Column_LastUpdated:             "%d",
		User_Column_Level:                   "%d",
		User_Column_Locale:                  "%s",
		User_Column_ParentUserID:            "%d",
		User_Column_PermissionCount:         "%d",
		User_Column_Points:                  "%d",
		User_Column_PolicyCount:             "%d",
		User_Column_ProfileID:               "%d",
		User_Column_ProfileImageID:          "%d",
		User_Column_ProfileImageThumbURL:    "%s",
		User_Column_ProfileImageURL:         "%s",
		User_Column_ReasonForLeaving:        "%s",
		User_Column_RoleCount:               "%d",
		User_Column_SMSPhoneNumber:          "%s",
		User_Column_SendSMSNotifications:    "%d",
		User_Column_Streak:                  "%d",
		User_Column_Timezone:                "%s",
		User_Column_TwoFactorEnabled:        "%d",
		User_Column_TwoFactorLastCompleted:  "%d",
		User_Column_UserGroupCount:          "%d",
		User_Column_UserID:                  "%d",
		User_Column_UserPrefix:              "%s",
		User_Column_Username:                "%s",
	}

	// Update columns
	// User_UpdateColumns is a list of all update columns for this model
	User_UpdateColumns = []goquery.Column{
		User_Column_APIKeyID,
		User_Column_AccountID,
		User_Column_AccountRole,
		User_Column_AppTheme,
		User_Column_Coins,
		User_Column_DailyStudiedTermsGoal,
		User_Column_DailyTestedTermsGoal,
		User_Column_DateActivated,
		User_Column_DateEmailValidated,
		User_Column_Email,
		User_Column_EnergyPoints,
		User_Column_FocusedCollectionID,
		User_Column_FollowerCount,
		User_Column_FollowingCount,
		User_Column_FullName,
		User_Column_GoogleSubID,
		User_Column_ID,
		User_Column_ImageID,
		User_Column_IsDisabled,
		User_Column_IsLocked,
		User_Column_IsPublic,
		User_Column_IsService,
		User_Column_IsShortcutsEnabled,
		User_Column_Language,
		User_Column_LastLogin,
		User_Column_LastLoginAttempt,
		User_Column_LastLoginAttemptCounter,
		User_Column_LastNotificationDate,
		User_Column_LastUpdated,
		User_Column_Level,
		User_Column_Locale,
		User_Column_ParentUserID,
		User_Column_PermissionCount,
		User_Column_Points,
		User_Column_PolicyCount,
		User_Column_ProfileID,
		User_Column_ProfileImageID,
		User_Column_ProfileImageThumbURL,
		User_Column_ProfileImageURL,
		User_Column_ReasonForLeaving,
		User_Column_RoleCount,
		User_Column_SMSPhoneNumber,
		User_Column_SendSMSNotifications,
		User_Column_Streak,
		User_Column_Timezone,
		User_Column_TwoFactorEnabled,
		User_Column_TwoFactorLastCompleted,
		User_Column_UserGroupCount,
		User_Column_UserPrefix,
		User_Column_Username,
	}

	// Insert columns
	// User_InsertColumns is a list of all insert columns for this model
	User_InsertColumns = []goquery.Column{
		User_Column_APIKeyID,
		User_Column_AccountID,
		User_Column_AccountRole,
		User_Column_AppTheme,
		User_Column_Coins,
		User_Column_DailyStudiedTermsGoal,
		User_Column_DailyTestedTermsGoal,
		User_Column_DateActivated,
		User_Column_DateCreated,
		User_Column_DateEmailValidated,
		User_Column_Email,
		User_Column_EnergyPoints,
		User_Column_FocusedCollectionID,
		User_Column_FollowerCount,
		User_Column_FollowingCount,
		User_Column_FullName,
		User_Column_GoogleSubID,
		User_Column_ID,
		User_Column_ImageID,
		User_Column_IsDisabled,
		User_Column_IsLocked,
		User_Column_IsPublic,
		User_Column_IsService,
		User_Column_IsShortcutsEnabled,
		User_Column_Language,
		User_Column_LastLogin,
		User_Column_LastLoginAttempt,
		User_Column_LastLoginAttemptCounter,
		User_Column_LastNotificationDate,
		User_Column_LastUpdated,
		User_Column_Level,
		User_Column_Locale,
		User_Column_ParentUserID,
		User_Column_PermissionCount,
		User_Column_Points,
		User_Column_PolicyCount,
		User_Column_ProfileID,
		User_Column_ProfileImageID,
		User_Column_ProfileImageThumbURL,
		User_Column_ProfileImageURL,
		User_Column_ReasonForLeaving,
		User_Column_RoleCount,
		User_Column_SMSPhoneNumber,
		User_Column_SendSMSNotifications,
		User_Column_Streak,
		User_Column_Timezone,
		User_Column_TwoFactorEnabled,
		User_Column_TwoFactorLastCompleted,
		User_Column_UserGroupCount,
		User_Column_UserPrefix,
		User_Column_Username,
	}

	// Primary Key
	// User_PrimaryKey is the name of the table's primary key
	User_PrimaryKey goquery.Column = "UserID"
)

// User is a data model
type User struct {
	APIKeyID                int64       `db:"APIKeyID" json:"APIKeyID"`
	AccountID               int64       `db:"AccountID" json:"AccountID"`
	AccountRole             int         `db:"AccountRole" json:"AccountRole"`
	AppTheme                string      `db:"AppTheme" json:"AppTheme"`
	Coins                   int64       `db:"Coins" json:"Coins"`
	DailyStudiedTermsGoal   int64       `db:"DailyStudiedTermsGoal" json:"DailyStudiedTermsGoal"`
	DailyTestedTermsGoal    int64       `db:"DailyTestedTermsGoal" json:"DailyTestedTermsGoal"`
	DateActivated           int64       `db:"DateActivated" json:"DateActivated"`
	DateCreated             int64       `db:"DateCreated" json:"DateCreated"`
	DateEmailValidated      int64       `db:"DateEmailValidated" json:"DateEmailValidated"`
	Email                   string      `db:"Email" json:"Email"`
	EnergyPoints            int64       `db:"EnergyPoints" json:"EnergyPoints"`
	FocusedCollectionID     int64       `db:"FocusedCollectionID" json:"FocusedCollectionID"`
	FollowerCount           int64       `db:"FollowerCount" json:"FollowerCount"`
	FollowingCount          int64       `db:"FollowingCount" json:"FollowingCount"`
	FullName                string      `db:"FullName" json:"FullName"`
	GoogleSubID             string      `db:"GoogleSubID" json:"GoogleSubID"`
	ID                      string      `db:"ID" json:"ID"`
	ImageID                 int64       `db:"ImageID" json:"ImageID"`
	IsDeleted               int         `db:"IsDeleted" json:"IsDeleted"`
	IsDisabled              int         `db:"IsDisabled" json:"IsDisabled"`
	IsLocked                int         `db:"IsLocked" json:"IsLocked"`
	IsPublic                int         `db:"IsPublic" json:"IsPublic"`
	IsService               int         `db:"IsService" json:"IsService"`
	IsShortcutsEnabled      int         `db:"IsShortcutsEnabled" json:"IsShortcutsEnabled"`
	Language                string      `db:"Language" json:"Language"`
	LastLogin               int64       `db:"LastLogin" json:"LastLogin"`
	LastLoginAttempt        int64       `db:"LastLoginAttempt" json:"LastLoginAttempt"`
	LastLoginAttemptCounter int         `db:"LastLoginAttemptCounter" json:"LastLoginAttemptCounter"`
	LastNotificationDate    int64       `db:"LastNotificationDate" json:"LastNotificationDate"`
	LastUpdated             int64       `db:"LastUpdated" json:"LastUpdated"`
	Level                   int64       `db:"Level" json:"Level"`
	Locale                  string      `db:"Locale" json:"Locale"`
	ParentUserID            int64       `db:"ParentUserID" json:"ParentUserID"`
	PermissionCount         int64       `db:"PermissionCount" json:"PermissionCount"`
	Points                  int64       `db:"Points" json:"Points"`
	PolicyCount             int64       `db:"PolicyCount" json:"PolicyCount"`
	ProfileID               int64       `db:"ProfileID" json:"ProfileID"`
	ProfileImageID          int64       `db:"ProfileImageID" json:"ProfileImageID"`
	ProfileImageThumbURL    string      `db:"ProfileImageThumbURL" json:"ProfileImageThumbURL"`
	ProfileImageURL         string      `db:"ProfileImageURL" json:"ProfileImageURL"`
	ReasonForLeaving        null.String `db:"ReasonForLeaving" json:"ReasonForLeaving"`
	RoleCount               int64       `db:"RoleCount" json:"RoleCount"`
	SMSPhoneNumber          string      `db:"SMSPhoneNumber" json:"SMSPhoneNumber"`
	SendSMSNotifications    int         `db:"SendSMSNotifications" json:"SendSMSNotifications"`
	Streak                  int64       `db:"Streak" json:"Streak"`
	Timezone                string      `db:"Timezone" json:"Timezone"`
	TwoFactorEnabled        int         `db:"TwoFactorEnabled" json:"TwoFactorEnabled"`
	TwoFactorLastCompleted  int64       `db:"TwoFactorLastCompleted" json:"TwoFactorLastCompleted"`
	UserGroupCount          int64       `db:"UserGroupCount" json:"UserGroupCount"`
	UserID                  int64       `db:"UserID" json:"UserID"`
	UserPrefix              string      `db:"UserPrefix" json:"UserPrefix"`
	Username                string      `db:"Username" json:"Username"`
}

// Account satifies the IAccountable interface
func (c *User) Account() int64 {
	return c.AccountID
} // 63

// User satisifies the IUserable interface
func (c *User) User() int64 {
	return c.UserID
} // 68

// User_TableName is the name of the table
func (c *User) Table_Name() goquery.TableName {
	return User_TableName
}

func (c *User) Table_Columns() []goquery.Column {
	return User_Columns
}

// Table_ColumnTypes returns a map of tableColumn names with their fmt string types
func (c *User) Table_Column_Types() map[goquery.Column]string {
	return User_Column_Types
}

// Table_PrimaryKey returns the name of this table's primary key
func (c *User) Table_PrimaryKey() goquery.Column {
	return User_PrimaryKey
}

// Table_PrimaryKey_Value returns the value of this table's primary key
func (c *User) Table_PrimaryKey_Value() int64 {
	return c.UserID
}

// Table_InsertColumns is a list of all insert columns for this model
func (c *User) Table_InsertColumns() []goquery.Column {
	return User_InsertColumns
}

// Table_UpdateColumns is a list of all update columns for this model
func (c *User) Table_UpdateColumns() []goquery.Column { // 100
	return User_UpdateColumns
}

// User_SchemaName is the name of this table's schema
func (c *User) Table_SchemaName() string {
	return User_SchemaName
}

// FromID returns a FromID query statement
func (c *User) FromID(db goquery.IDB, id int64) (goquery.IModel, error) {

	sel := goquery.Select(c)
	sel.Fields(
		goquery.NewField(goquery.FieldTypeBasic, User_Column_APIKeyID),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_AccountID),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_AccountRole),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_AppTheme),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_Coins),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_DailyStudiedTermsGoal),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_DailyTestedTermsGoal),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_DateActivated),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_DateCreated),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_DateEmailValidated),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_Email),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_EnergyPoints),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_FocusedCollectionID),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_FollowerCount),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_FollowingCount),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_FullName),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_GoogleSubID),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_ID),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_ImageID),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_IsDeleted),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_IsDisabled),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_IsLocked),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_IsPublic),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_IsService),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_IsShortcutsEnabled),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_Language),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_LastLogin),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_LastLoginAttempt),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_LastLoginAttemptCounter),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_LastNotificationDate),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_LastUpdated),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_Level),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_Locale),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_ParentUserID),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_PermissionCount),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_Points),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_PolicyCount),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_ProfileID),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_ProfileImageID),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_ProfileImageThumbURL),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_ProfileImageURL),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_ReasonForLeaving),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_RoleCount),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_SMSPhoneNumber),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_SendSMSNotifications),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_Streak),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_Timezone),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_TwoFactorEnabled),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_TwoFactorLastCompleted),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_UserGroupCount),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_UserID),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_UserPrefix),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_Username),
	)
	q, e := sel.String()
	if e != nil {
		return nil, fmt.Errorf("User.FromID.Query.String(): %w", e)
	}

	row := db.QueryRow(q)

	switch e = row.Scan(
		&c.APIKeyID,
		&c.AccountID,
		&c.AccountRole,
		&c.AppTheme,
		&c.Coins,
		&c.DailyStudiedTermsGoal,
		&c.DailyTestedTermsGoal,
		&c.DateActivated,
		&c.DateCreated,
		&c.DateEmailValidated,
		&c.Email,
		&c.EnergyPoints,
		&c.FocusedCollectionID,
		&c.FollowerCount,
		&c.FollowingCount,
		&c.FullName,
		&c.GoogleSubID,
		&c.ID,
		&c.ImageID,
		&c.IsDeleted,
		&c.IsDisabled,
		&c.IsLocked,
		&c.IsPublic,
		&c.IsService,
		&c.IsShortcutsEnabled,
		&c.Language,
		&c.LastLogin,
		&c.LastLoginAttempt,
		&c.LastLoginAttemptCounter,
		&c.LastNotificationDate,
		&c.LastUpdated,
		&c.Level,
		&c.Locale,
		&c.ParentUserID,
		&c.PermissionCount,
		&c.Points,
		&c.PolicyCount,
		&c.ProfileID,
		&c.ProfileImageID,
		&c.ProfileImageThumbURL,
		&c.ProfileImageURL,
		&c.ReasonForLeaving,
		&c.RoleCount,
		&c.SMSPhoneNumber,
		&c.SendSMSNotifications,
		&c.Streak,
		&c.Timezone,
		&c.TwoFactorEnabled,
		&c.TwoFactorLastCompleted,
		&c.UserGroupCount,
		&c.UserID,
		&c.UserPrefix,
		&c.Username,
	); e {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		// fmt.Printf("UserDALGetter.Get(%s).Run()\n", q)
		return c, nil
	default:
		return nil, fmt.Errorf("UserDALGetter(%s).Run(): %w", q, e)
	}
}

// String returns a json marshalled string of the object
func (c *User) String() string {
	bytes, _ := json.Marshal(c)
	return string(bytes)
}

// Update updates a User record
func (c *User) Update(db goquery.IDB) error {
	var e error
	var ql string
	ql, _ = goquery.Update(c).
		Set(User_Column_APIKeyID, c.APIKeyID).
		Set(User_Column_AccountID, c.AccountID).
		Set(User_Column_AccountRole, c.AccountRole).
		Set(User_Column_AppTheme, c.AppTheme).
		Set(User_Column_Coins, c.Coins).
		Set(User_Column_DailyStudiedTermsGoal, c.DailyStudiedTermsGoal).
		Set(User_Column_DailyTestedTermsGoal, c.DailyTestedTermsGoal).
		Set(User_Column_DateActivated, c.DateActivated).
		Set(User_Column_DateEmailValidated, c.DateEmailValidated).
		Set(User_Column_Email, c.Email).
		Set(User_Column_EnergyPoints, c.EnergyPoints).
		Set(User_Column_FocusedCollectionID, c.FocusedCollectionID).
		Set(User_Column_FollowerCount, c.FollowerCount).
		Set(User_Column_FollowingCount, c.FollowingCount).
		Set(User_Column_FullName, c.FullName).
		Set(User_Column_GoogleSubID, c.GoogleSubID).
		Set(User_Column_ID, c.ID).
		Set(User_Column_ImageID, c.ImageID).
		Set(User_Column_IsDisabled, c.IsDisabled).
		Set(User_Column_IsLocked, c.IsLocked).
		Set(User_Column_IsPublic, c.IsPublic).
		Set(User_Column_IsService, c.IsService).
		Set(User_Column_IsShortcutsEnabled, c.IsShortcutsEnabled).
		Set(User_Column_Language, c.Language).
		Set(User_Column_LastLogin, c.LastLogin).
		Set(User_Column_LastLoginAttempt, c.LastLoginAttempt).
		Set(User_Column_LastLoginAttemptCounter, c.LastLoginAttemptCounter).
		Set(User_Column_LastNotificationDate, c.LastNotificationDate).
		Set(User_Column_LastUpdated, c.LastUpdated).
		Set(User_Column_Level, c.Level).
		Set(User_Column_Locale, c.Locale).
		Set(User_Column_ParentUserID, c.ParentUserID).
		Set(User_Column_PermissionCount, c.PermissionCount).
		Set(User_Column_Points, c.Points).
		Set(User_Column_PolicyCount, c.PolicyCount).
		Set(User_Column_ProfileID, c.ProfileID).
		Set(User_Column_ProfileImageID, c.ProfileImageID).
		Set(User_Column_ProfileImageThumbURL, c.ProfileImageThumbURL).
		Set(User_Column_ProfileImageURL, c.ProfileImageURL).
		Set(User_Column_ReasonForLeaving, c.ReasonForLeaving.String).
		Set(User_Column_RoleCount, c.RoleCount).
		Set(User_Column_SMSPhoneNumber, c.SMSPhoneNumber).
		Set(User_Column_SendSMSNotifications, c.SendSMSNotifications).
		Set(User_Column_Streak, c.Streak).
		Set(User_Column_Timezone, c.Timezone).
		Set(User_Column_TwoFactorEnabled, c.TwoFactorEnabled).
		Set(User_Column_TwoFactorLastCompleted, c.TwoFactorLastCompleted).
		Set(User_Column_UserGroupCount, c.UserGroupCount).
		Set(User_Column_UserPrefix, c.UserPrefix).
		Set(User_Column_Username, c.Username).
		Where(goquery.EQ(User_Column_UserID, c.UserID)).
		String()
	_, e = db.Exec(ql)
	if e != nil {
		return fmt.Errorf("User.Update(): %w", e)
	}

	return e
}

// Create inserts a User record
func (c *User) Create(db goquery.IDB) error {

	var e error

	q := goquery.Insert(c)

	if c.UserID > 0 {
		q.Set(User_Column_UserID, c.UserID)
	}

	q.Set(User_Column_APIKeyID, c.APIKeyID)
	q.Set(User_Column_AccountID, c.AccountID)
	q.Set(User_Column_AccountRole, c.AccountRole)
	q.Set(User_Column_AppTheme, c.AppTheme)
	q.Set(User_Column_Coins, c.Coins)
	q.Set(User_Column_DailyStudiedTermsGoal, c.DailyStudiedTermsGoal)
	q.Set(User_Column_DailyTestedTermsGoal, c.DailyTestedTermsGoal)
	q.Set(User_Column_DateActivated, c.DateActivated)
	q.Set(User_Column_DateCreated, c.DateCreated)
	q.Set(User_Column_DateEmailValidated, c.DateEmailValidated)
	q.Set(User_Column_Email, c.Email)
	q.Set(User_Column_EnergyPoints, c.EnergyPoints)
	q.Set(User_Column_FocusedCollectionID, c.FocusedCollectionID)
	q.Set(User_Column_FollowerCount, c.FollowerCount)
	q.Set(User_Column_FollowingCount, c.FollowingCount)
	q.Set(User_Column_FullName, c.FullName)
	q.Set(User_Column_GoogleSubID, c.GoogleSubID)
	q.Set(User_Column_ID, c.ID)
	q.Set(User_Column_ImageID, c.ImageID)
	q.Set(User_Column_IsDisabled, c.IsDisabled)
	q.Set(User_Column_IsLocked, c.IsLocked)
	q.Set(User_Column_IsPublic, c.IsPublic)
	q.Set(User_Column_IsService, c.IsService)
	q.Set(User_Column_IsShortcutsEnabled, c.IsShortcutsEnabled)
	q.Set(User_Column_Language, c.Language)
	q.Set(User_Column_LastLogin, c.LastLogin)
	q.Set(User_Column_LastLoginAttempt, c.LastLoginAttempt)
	q.Set(User_Column_LastLoginAttemptCounter, c.LastLoginAttemptCounter)
	q.Set(User_Column_LastNotificationDate, c.LastNotificationDate)
	q.Set(User_Column_LastUpdated, c.LastUpdated)
	q.Set(User_Column_Level, c.Level)
	q.Set(User_Column_Locale, c.Locale)
	q.Set(User_Column_ParentUserID, c.ParentUserID)
	q.Set(User_Column_PermissionCount, c.PermissionCount)
	q.Set(User_Column_Points, c.Points)
	q.Set(User_Column_PolicyCount, c.PolicyCount)
	q.Set(User_Column_ProfileID, c.ProfileID)
	q.Set(User_Column_ProfileImageID, c.ProfileImageID)
	q.Set(User_Column_ProfileImageThumbURL, c.ProfileImageThumbURL)
	q.Set(User_Column_ProfileImageURL, c.ProfileImageURL)
	q.Set(User_Column_ReasonForLeaving, c.ReasonForLeaving.String)
	q.Set(User_Column_RoleCount, c.RoleCount)
	q.Set(User_Column_SMSPhoneNumber, c.SMSPhoneNumber)
	q.Set(User_Column_SendSMSNotifications, c.SendSMSNotifications)
	q.Set(User_Column_Streak, c.Streak)
	q.Set(User_Column_Timezone, c.Timezone)
	q.Set(User_Column_TwoFactorEnabled, c.TwoFactorEnabled)
	q.Set(User_Column_TwoFactorLastCompleted, c.TwoFactorLastCompleted)
	q.Set(User_Column_UserGroupCount, c.UserGroupCount)
	q.Set(User_Column_UserPrefix, c.UserPrefix)
	q.Set(User_Column_Username, c.Username)

	ql, _ := q.String()
	var result sql.Result
	result, e = db.Exec(ql)

	if e != nil {
		return fmt.Errorf("User.Create(): %w", e) // 177
	}

	// Assumes auto-increment
	if c.UserID == 0 {
		c.UserID, e = result.LastInsertId()
	}

	return e
}

// Destroy deletes a User record
func (c *User) Delete(db goquery.IDB) error {
	var e error
	ql, _ := goquery.Delete(c).
		Where(
			goquery.EQ(User_Column_UserID, c.UserID),
		).String()

	_, e = db.Exec(ql)
	if e != nil {
		return fmt.Errorf("User.Delete(): %w", e)
	}

	return e
}

func (r *User) Raw(db goquery.IDB, queryRaw string) ([]*User, error) {

	var e error
	model := []*User{}
	sel := goquery.Select(r)
	sel.Fields(
		goquery.NewField(goquery.FieldTypeBasic, User_Column_APIKeyID),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_AccountID),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_AccountRole),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_AppTheme),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_Coins),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_DailyStudiedTermsGoal),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_DailyTestedTermsGoal),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_DateActivated),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_DateCreated),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_DateEmailValidated),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_Email),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_EnergyPoints),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_FocusedCollectionID),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_FollowerCount),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_FollowingCount),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_FullName),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_GoogleSubID),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_ID),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_ImageID),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_IsDeleted),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_IsDisabled),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_IsLocked),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_IsPublic),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_IsService),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_IsShortcutsEnabled),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_Language),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_LastLogin),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_LastLoginAttempt),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_LastLoginAttemptCounter),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_LastNotificationDate),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_LastUpdated),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_Level),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_Locale),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_ParentUserID),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_PermissionCount),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_Points),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_PolicyCount),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_ProfileID),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_ProfileImageID),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_ProfileImageThumbURL),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_ProfileImageURL),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_ReasonForLeaving),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_RoleCount),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_SMSPhoneNumber),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_SendSMSNotifications),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_Streak),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_Timezone),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_TwoFactorEnabled),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_TwoFactorLastCompleted),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_UserGroupCount),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_UserID),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_UserPrefix),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_Username),
	)

	q, e := sel.String()
	if e != nil {
		return nil, fmt.Errorf("UserDAL.Raw.String(): %w", e)
	}

	var rows *sql.Rows
	rows, e = db.Query(q)

	if e != nil {
		if e == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("UserDAL.Raw.Run(%s): %w", q, e)
	}

	defer rows.Close()
	for rows.Next() {
		m := &User{}
		if e = rows.Scan(
			&m.APIKeyID,
			&m.AccountID,
			&m.AccountRole,
			&m.AppTheme,
			&m.Coins,
			&m.DailyStudiedTermsGoal,
			&m.DailyTestedTermsGoal,
			&m.DateActivated,
			&m.DateCreated,
			&m.DateEmailValidated,
			&m.Email,
			&m.EnergyPoints,
			&m.FocusedCollectionID,
			&m.FollowerCount,
			&m.FollowingCount,
			&m.FullName,
			&m.GoogleSubID,
			&m.ID,
			&m.ImageID,
			&m.IsDeleted,
			&m.IsDisabled,
			&m.IsLocked,
			&m.IsPublic,
			&m.IsService,
			&m.IsShortcutsEnabled,
			&m.Language,
			&m.LastLogin,
			&m.LastLoginAttempt,
			&m.LastLoginAttemptCounter,
			&m.LastNotificationDate,
			&m.LastUpdated,
			&m.Level,
			&m.Locale,
			&m.ParentUserID,
			&m.PermissionCount,
			&m.Points,
			&m.PolicyCount,
			&m.ProfileID,
			&m.ProfileImageID,
			&m.ProfileImageThumbURL,
			&m.ProfileImageURL,
			&m.ReasonForLeaving,
			&m.RoleCount,
			&m.SMSPhoneNumber,
			&m.SendSMSNotifications,
			&m.Streak,
			&m.Timezone,
			&m.TwoFactorEnabled,
			&m.TwoFactorLastCompleted,
			&m.UserGroupCount,
			&m.UserID,
			&m.UserPrefix,
			&m.Username,
		); e != nil {
			return nil, fmt.Errorf("UserDALRaw(%s).Run(): %w", q, e)
		}
		model = append(model, m)
	}

	// fmt.Printf("UserDAL.Raw(%s).Run()\n", q)

	return model, nil
}

type IUserDALSelector interface {
	Select(db goquery.IDB) IUserDALSelector
}

type UserDALSelector struct {
	db       goquery.IDB
	q        *goquery.Q
	isSingle bool
}

func (r *User) Select(db goquery.IDB) *UserDALSelector {
	return &UserDALSelector{
		db: db,
		q:  goquery.Select(r),
	}
}

func (r *UserDALSelector) Alias(alias string) *UserDALSelector {
	r.q.Alias(alias)
	return r
}

func (r *UserDALSelector) Where(whereParts ...*goquery.WherePart) *UserDALSelector {
	r.q.Where(whereParts...)
	return r
}

func (r *UserDALSelector) Limit(limit, offset int64) *UserDALSelector {
	r.q = r.q.Limit(limit, offset)
	return r
}

func (r *UserDALSelector) OrderBy(col goquery.Column, dir goquery.OrderDir) *UserDALSelector {
	r.q = r.q.OrderBy(col, dir)
	return r
}

func (r *UserDALSelector) String() (string, error) {

	r.q.Fields(
		goquery.NewField(goquery.FieldTypeBasic, User_Column_APIKeyID),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_AccountID),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_AccountRole),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_AppTheme),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_Coins),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_DailyStudiedTermsGoal),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_DailyTestedTermsGoal),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_DateActivated),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_DateCreated),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_DateEmailValidated),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_Email),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_EnergyPoints),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_FocusedCollectionID),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_FollowerCount),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_FollowingCount),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_FullName),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_GoogleSubID),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_ID),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_ImageID),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_IsDeleted),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_IsDisabled),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_IsLocked),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_IsPublic),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_IsService),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_IsShortcutsEnabled),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_Language),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_LastLogin),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_LastLoginAttempt),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_LastLoginAttemptCounter),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_LastNotificationDate),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_LastUpdated),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_Level),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_Locale),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_ParentUserID),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_PermissionCount),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_Points),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_PolicyCount),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_ProfileID),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_ProfileImageID),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_ProfileImageThumbURL),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_ProfileImageURL),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_ReasonForLeaving),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_RoleCount),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_SMSPhoneNumber),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_SendSMSNotifications),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_Streak),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_Timezone),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_TwoFactorEnabled),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_TwoFactorLastCompleted),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_UserGroupCount),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_UserID),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_UserPrefix),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_Username),
	)

	q, e := r.q.String()
	if e != nil {
		return "", fmt.Errorf("UserDAL.Query.String(): %w", e)
	}

	return q, nil
}

func (r *UserDALSelector) Run() ([]*User, error) {

	var e error
	var q = ""
	var model = []*User{}
	q, e = r.String()
	if e != nil {
		return nil, fmt.Errorf("UserDALSelector.Query.String(): %w", e)
	}

	var rows *sql.Rows
	rows, e = r.db.Query(q)

	if e != nil {
		if e == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("UserDALSelector.Run(%s): %w", q, e)
	}

	defer rows.Close()
	for rows.Next() {
		m := &User{}
		if e = rows.Scan(
			&m.APIKeyID,
			&m.AccountID,
			&m.AccountRole,
			&m.AppTheme,
			&m.Coins,
			&m.DailyStudiedTermsGoal,
			&m.DailyTestedTermsGoal,
			&m.DateActivated,
			&m.DateCreated,
			&m.DateEmailValidated,
			&m.Email,
			&m.EnergyPoints,
			&m.FocusedCollectionID,
			&m.FollowerCount,
			&m.FollowingCount,
			&m.FullName,
			&m.GoogleSubID,
			&m.ID,
			&m.ImageID,
			&m.IsDeleted,
			&m.IsDisabled,
			&m.IsLocked,
			&m.IsPublic,
			&m.IsService,
			&m.IsShortcutsEnabled,
			&m.Language,
			&m.LastLogin,
			&m.LastLoginAttempt,
			&m.LastLoginAttemptCounter,
			&m.LastNotificationDate,
			&m.LastUpdated,
			&m.Level,
			&m.Locale,
			&m.ParentUserID,
			&m.PermissionCount,
			&m.Points,
			&m.PolicyCount,
			&m.ProfileID,
			&m.ProfileImageID,
			&m.ProfileImageThumbURL,
			&m.ProfileImageURL,
			&m.ReasonForLeaving,
			&m.RoleCount,
			&m.SMSPhoneNumber,
			&m.SendSMSNotifications,
			&m.Streak,
			&m.Timezone,
			&m.TwoFactorEnabled,
			&m.TwoFactorLastCompleted,
			&m.UserGroupCount,
			&m.UserID,
			&m.UserPrefix,
			&m.Username,
		); e != nil {
			return nil, fmt.Errorf("UserDALSelector(%s).Run(): %w", q, e)
		}

		model = append(model, m)
	}

	// fmt.Printf("UserDALSelector(%s).Run()\n", q)

	return model, nil
}

// Counter
type UserDALCounter struct {
	db goquery.IDB
	q  *goquery.Q
}

func (r *User) Count(db goquery.IDB) *UserDALCounter {
	return &UserDALCounter{
		db: db,
		q:  goquery.Select(r).Count(r.Table_PrimaryKey(), "c"),
	}
}

func (r *UserDALCounter) Alias(alias string) *UserDALCounter {
	r.q.Alias(alias)
	return r
}

func (ds *UserDALCounter) Where(whereParts ...*goquery.WherePart) *UserDALCounter {
	ds.q.Where(whereParts...)
	return ds
}

func (ds *UserDALCounter) Run() (int64, error) {

	count := int64(0)
	q, e := ds.q.String()
	if e != nil {
		return 0, fmt.Errorf("UserDALCounter.Query.String(): %w", e)
	}

	row := ds.db.QueryRow(q)

	switch e = row.Scan(&count); e {
	case sql.ErrNoRows:
		return 0, nil
	case nil:
		// fmt.Printf("UserDALCounter.QueryRow(%s).Run()\n", q)
		return count, nil
	default:
		return 0, fmt.Errorf("UserDALCounter.QueryRow(%s).Run(): %w", q, e)
	}
}

// Summer
type UserDALSummer struct {
	db goquery.IDB
	q  *goquery.Q
}

func (r *User) Sum(db goquery.IDB, col goquery.Column) *UserDALSummer {
	return &UserDALSummer{
		db: db,
		q:  goquery.Select(r).Sum(col, "c"),
	}
}

func (ds *UserDALSummer) Where(whereParts ...*goquery.WherePart) *UserDALSummer {
	ds.q.Where(whereParts...)
	return ds
}

func (ds *UserDALSummer) Run() (float64, error) {

	sum := float64(0)
	q, e := ds.q.String()
	if e != nil {
		return 0, fmt.Errorf("UserDALSummer.Query.String(): %w", e)
	}

	row := ds.db.QueryRow(q)

	switch e = row.Scan(&sum); e {
	case sql.ErrNoRows:
		return 0, nil
	case nil:
		// fmt.Printf("UserDALSummer.QueryRow(%s).Run()\n", q)
		return sum, nil
	default:
		return 0, fmt.Errorf("UserDALSummer.QueryRow(%s).Run(): %w", q, e)
	}
}

// Minner
type UserDALMinner struct {
	db goquery.IDB
	q  *goquery.Q
}

func (r *User) Min(db goquery.IDB, col goquery.Column) *UserDALMinner {
	return &UserDALMinner{
		db: db,
		q:  goquery.Select(r).Min(col, "c"),
	}
}

func (ds *UserDALMinner) Where(whereParts ...*goquery.WherePart) *UserDALMinner {
	ds.q.Where(whereParts...)
	return ds
}

func (ds *UserDALMinner) Run() (float64, error) {

	min := float64(0)
	q, e := ds.q.String()
	if e != nil {
		return 0, fmt.Errorf("UserDALMinner.Query.String(): %w", e)
	}

	row := ds.db.QueryRow(q)

	switch e = row.Scan(&min); e {
	case sql.ErrNoRows:
		return 0, nil
	case nil:
		// fmt.Printf("UserDALMinner.QueryRow(%s).Run()\n", q)
		return min, nil
	default:
		return 0, fmt.Errorf("UserDALMinner.QueryRow(%s).Run(): %w", q, e)
	}
}

// Maxer
type UserDALMaxer struct {
	db goquery.IDB
	q  *goquery.Q
}

func (r *User) Max(db goquery.IDB, col goquery.Column) *UserDALMaxer {
	return &UserDALMaxer{
		db: db,
		q:  goquery.Select(r).Max(col, "c"),
	}
}

func (ds *UserDALMaxer) Where(whereParts ...*goquery.WherePart) *UserDALMaxer {
	ds.q.Where(whereParts...)
	return ds
}

func (ds *UserDALMaxer) Run() (float64, error) {

	max := float64(0)
	q, e := ds.q.String()
	if e != nil {
		return 0, fmt.Errorf("UserDALMaxer.Query.String(): %w", e)
	}

	row := ds.db.QueryRow(q)

	switch e = row.Scan(&max); e {
	case sql.ErrNoRows:
		return 0, nil
	case nil:
		// fmt.Printf("UserDALMaxer.QueryRow(%s).Run()\n", q)
		return max, nil
	default:
		return 0, fmt.Errorf("UserDALMaxer.QueryRow(%s).Run(): %w", q, e)
	}
}

type UserDALGetter struct {
	db goquery.IDB
	q  *goquery.Q
}

func (r *User) Get(db goquery.IDB) *UserDALGetter {
	return &UserDALGetter{
		db: db,
		q:  goquery.Select(r),
	}
}

func (r *UserDALGetter) Alias(alias string) *UserDALGetter {
	r.q.Alias(alias)
	return r
}

func (ds *UserDALGetter) Where(whereParts ...*goquery.WherePart) *UserDALGetter {
	ds.q.Where(whereParts...)
	return ds
}

func (ds *UserDALGetter) OrderBy(col goquery.Column, dir goquery.OrderDir) *UserDALGetter {
	ds.q = ds.q.OrderBy(col, dir)
	return ds
}

func (ds *UserDALGetter) Run() (*User, error) {

	model := &User{}

	ds.q.Fields(
		goquery.NewField(goquery.FieldTypeBasic, User_Column_APIKeyID),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_AccountID),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_AccountRole),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_AppTheme),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_Coins),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_DailyStudiedTermsGoal),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_DailyTestedTermsGoal),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_DateActivated),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_DateCreated),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_DateEmailValidated),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_Email),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_EnergyPoints),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_FocusedCollectionID),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_FollowerCount),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_FollowingCount),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_FullName),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_GoogleSubID),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_ID),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_ImageID),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_IsDeleted),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_IsDisabled),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_IsLocked),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_IsPublic),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_IsService),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_IsShortcutsEnabled),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_Language),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_LastLogin),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_LastLoginAttempt),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_LastLoginAttemptCounter),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_LastNotificationDate),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_LastUpdated),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_Level),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_Locale),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_ParentUserID),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_PermissionCount),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_Points),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_PolicyCount),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_ProfileID),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_ProfileImageID),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_ProfileImageThumbURL),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_ProfileImageURL),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_ReasonForLeaving),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_RoleCount),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_SMSPhoneNumber),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_SendSMSNotifications),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_Streak),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_Timezone),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_TwoFactorEnabled),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_TwoFactorLastCompleted),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_UserGroupCount),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_UserID),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_UserPrefix),
		goquery.NewField(goquery.FieldTypeBasic, User_Column_Username),
	)
	q, e := ds.q.String()
	if e != nil {
		return nil, fmt.Errorf("UserDALGetter.Query.String(): %w", e)
	}

	row := ds.db.QueryRow(q)

	switch e = row.Scan(
		&model.APIKeyID,
		&model.AccountID,
		&model.AccountRole,
		&model.AppTheme,
		&model.Coins,
		&model.DailyStudiedTermsGoal,
		&model.DailyTestedTermsGoal,
		&model.DateActivated,
		&model.DateCreated,
		&model.DateEmailValidated,
		&model.Email,
		&model.EnergyPoints,
		&model.FocusedCollectionID,
		&model.FollowerCount,
		&model.FollowingCount,
		&model.FullName,
		&model.GoogleSubID,
		&model.ID,
		&model.ImageID,
		&model.IsDeleted,
		&model.IsDisabled,
		&model.IsLocked,
		&model.IsPublic,
		&model.IsService,
		&model.IsShortcutsEnabled,
		&model.Language,
		&model.LastLogin,
		&model.LastLoginAttempt,
		&model.LastLoginAttemptCounter,
		&model.LastNotificationDate,
		&model.LastUpdated,
		&model.Level,
		&model.Locale,
		&model.ParentUserID,
		&model.PermissionCount,
		&model.Points,
		&model.PolicyCount,
		&model.ProfileID,
		&model.ProfileImageID,
		&model.ProfileImageThumbURL,
		&model.ProfileImageURL,
		&model.ReasonForLeaving,
		&model.RoleCount,
		&model.SMSPhoneNumber,
		&model.SendSMSNotifications,
		&model.Streak,
		&model.Timezone,
		&model.TwoFactorEnabled,
		&model.TwoFactorLastCompleted,
		&model.UserGroupCount,
		&model.UserID,
		&model.UserPrefix,
		&model.Username,
	); e {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		// fmt.Printf("UserDALGetter.Get(%s).Run()\n", q)
		return model, nil
	default:
		return nil, fmt.Errorf("UserDALGetter(%s).Run(): %w", q, e)
	}
}
