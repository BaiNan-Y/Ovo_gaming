package model

import "time"

type BaseModel struct {
	ID        uint64    `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	BaseModel
	OpenID      string `gorm:"uniqueIndex;not null"`
	UnionID     string
	Username    string
	Nickname    string
	AvatarURL   string
	Phone       string
	Status      string
}

type UserRole struct {
	BaseModel
	UserID             uint64
	RoleType           string
	RoleStatus         string
	IsDefault          bool
	ApplyStatus        string
	ApplyRemark        string
	ReviewedBy         *uint64
	ReviewedAt         *time.Time
	GameMain           string
	RankInfo           string
	ServiceTags        string `gorm:"type:text"`
	ServiceDesc        string
	AcceptOrderStatus   bool
	OnlineStatus       bool
	ServiceScore       float64
	CompletedCount     int
	BalanceAvailable   float64
	BalanceFrozen      float64
	BalanceTotal       float64
}

type Banner struct {
	BaseModel
	Title      string
	Subtitle   string
	ImageURL   string
	JumpType   string
	JumpTarget string
	Status     string
	Sort       int
}

type Notice struct {
	BaseModel
	Title     string
	Content   string
	NoticeType string
	Status    string
	Sort      int
}

type Package struct {
	BaseModel
	GameCode        string
	Title           string
	Price           float64
	OriginalPrice    float64
	DurationMinutes int
	Description     string
	CoverImage      string
	IsHot           bool
	Status          string
	Sort            int
}

type Discount struct {
	BaseModel
	Title               string
	DiscountType        string
	DiscountMode        string
	DiscountValue       float64
	ConditionMinAmount  float64
	ScopeType           string
	ScopeIDs            string `gorm:"type:text"`
	Status              string
}

type Order struct {
	BaseModel
	OrderNo            string `gorm:"uniqueIndex;not null"`
	BossUserID         uint64
	PackageID          uint64
	PackageSnapshot    string `gorm:"type:jsonb"`
	DiscountID         *uint64
	DiscountSnapshot   string `gorm:"type:jsonb"`
	SpecifiedPlayerID  *uint64
	AssignedPlayerID   *uint64
	OrderAmount        float64
	DiscountAmount     float64
	PayAmount          float64
	PayStatus          string
	OrderStatus        string
	DispatchType       string
}

type OrderAssignment struct {
	BaseModel
	OrderID      uint64
	PlayerID     uint64
	AssignType   string
	AssignStatus string
}

type OrderLog struct {
	BaseModel
	OrderID          uint64
	ActionType       string
	OperatorUserID   *uint64
	OperatorRoleType string
	BeforeStatus     string
	AfterStatus      string
	Note             string
}

type Withdrawal struct {
	BaseModel
	PlayerUserID  uint64
	Amount        float64
	WithdrawMethod string
	AccountInfo   string `gorm:"type:jsonb"`
	WithdrawStatus string
}

type Review struct {
	BaseModel
	OrderID       uint64
	BossUserID    uint64
	PlayerUserID  uint64
	Score         int
	Content       string
	Status        string
}

type AuditLog struct {
	BaseModel
	OperatorUserID   *uint64
	OperatorRoleType string
	ActionType       string
	TargetType       string
	TargetID         *uint64
	Note             string
}
