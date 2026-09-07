package db

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Seed() error {

	DB.AutoMigrate(Role{})
	DB.AutoMigrate(PassStatus{})
	DB.AutoMigrate(AccessResult{})
	DB.AutoMigrate(User{})
	DB.AutoMigrate(Pass{})
	DB.AutoMigrate(Gate{})

	log.Println("🌱 Seeding lookup tables...")

	// 1. Seed Roles
	roles := []string{"ADMIN", "OPERATOR", "USER"}
	for _, roleName := range roles {
		var role Role
		if err := DB.Where("role = ?", roleName).FirstOrCreate(&role, Role{Role: roleName}).Error; err != nil {
			return fmt.Errorf("failed to seed role %s: %w", roleName, err)
		}
	}

	// 2. Seed Pass Statuses
	statuses := []string{"ACTIVE", "REVOKED", "EXPIRED"}
	for _, statusName := range statuses {
		var status PassStatus
		if err := DB.Where("status = ?", statusName).FirstOrCreate(&status, PassStatus{Status: statusName}).Error; err != nil {
			return fmt.Errorf("failed to seed pass status %s: %w", statusName, err)
		}
	}

	// 3. Seed Access Results
	results := []string{
		"GRANTED",
		"DENIED_EXPIRED",
		"DENIED_REVOKED",
		"DENIED_INVALID_ZONE",
		"DENIED_RATE_LIMIT",
	}
	for _, resultName := range results {
		var accessResult AccessResult
		if err := DB.Where("result = ?", resultName).FirstOrCreate(&accessResult, AccessResult{Result: resultName}).Error; err != nil {
			return fmt.Errorf("failed to seed access result %s: %w", resultName, err)
		}
	}

	// 4. Seed Default Admin User
	var adminRole Role
	if err := DB.Where("role = ?", "ADMIN").First(&adminRole).Error; err != nil {
		return fmt.Errorf("failed to fetch ADMIN role for seeding user: %w", err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("AdminPass123!"), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash admin password: %w", err)
	}

	adminUser := User{
		Email:        "admin@smartpass.alps",
		FullName:     "SmartPass Administrator",
		PasswordHash: string(hashedPassword),
		RoleID:       adminRole.ID,
	}

	if err := DB.Where("email = ?", adminUser.Email).FirstOrCreate(&adminUser, adminUser).Error; err != nil {
		return fmt.Errorf("failed to seed admin user: %w", err)
	}

	log.Println("✅ Database seeding complete!")
	log.Println("🔑 Default Admin: admin@smartpass.alps / AdminPass123!")
	return nil
}
