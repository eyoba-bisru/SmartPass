package db

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Role struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Role      string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"role"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

type PassStatus struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Status    string    `gorm:"type:varchar(20);uniqueIndex;not null" json:"status"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

type AccessResult struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Result    string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"result"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

type User struct {
	ID           uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Email        string     `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	PasswordHash string     `gorm:"type:varchar(255);not null" json:"-"`
	FullName     string     `gorm:"type:varchar(255);not null" json:"fullName"`
	RoleID       uuid.UUID  `gorm:"type:uuid;not null" json:"roleId"`
	Role         Role       `gorm:"foreignKey:RoleID" json:"role,omitempty"`
	CreatedAt    time.Time  `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime" json:"updatedAt"`
	CreatedBy    *uuid.UUID `gorm:"type:uuid" json:"createdBy,omitempty"`
}

type Pass struct {
	ID           uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	PassCode     string         `gorm:"type:varchar(512);uniqueIndex;not null" json:"passCode"`
	UserID       uuid.UUID      `gorm:"type:uuid;not null" json:"userId"`
	User         User           `gorm:"foreignKey:UserID" json:"user,omitempty"`
	StatusID     uuid.UUID      `gorm:"type:uuid;not null" json:"statusId"`
	Status       PassStatus     `gorm:"foreignKey:StatusID" json:"status,omitempty"`
	ValidFrom    time.Time      `gorm:"not null" json:"validFrom"`
	ValidUntil   time.Time      `gorm:"not null" json:"validUntil"`
	AllowedZones pq.StringArray `gorm:"type:text[];not null" json:"allowedZones"`
	CreatedAt    time.Time      `gorm:"autoCreateTime" json:"createdAt"`
}

type Gate struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	ZoneCode  string    `gorm:"type:varchar(100);not null" json:"zoneCode"`
	IsOnline  bool      `gorm:"default:true;not null" json:"isOnline"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
}

type AccessLog struct {
	ID             uuid.UUID    `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	PassID         *uuid.UUID   `gorm:"type:uuid" json:"passId,omitempty"`
	Pass           *Pass        `gorm:"foreignKey:PassID" json:"pass,omitempty"`
	GateID         uuid.UUID    `gorm:"type:uuid;not null;index:idx_access_logs_gate_time" json:"gateId"`
	Gate           Gate         `gorm:"foreignKey:GateID" json:"gate,omitempty"`
	AccessResultID uuid.UUID    `gorm:"type:uuid;not null" json:"accessResultId"`
	AccessResult   AccessResult `gorm:"foreignKey:AccessResultID" json:"accessResult,omitempty"`
	LatencyMs      int          `gorm:"not null" json:"latencyMs"`
	Timestamp      time.Time    `gorm:"autoCreateTime;index:idx_access_logs_gate_time,sort:desc" json:"timestamp"`
}
