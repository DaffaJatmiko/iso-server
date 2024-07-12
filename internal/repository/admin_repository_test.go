package repository_test

import (
	"github.com/DaffaJatmiko/project-iso/internal/model"
	"github.com/DaffaJatmiko/project-iso/internal/repository"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var _ = Describe("User Repository", func() {
	var (
		db        *gorm.DB
		adminRepo repository.UserRepository
	)

	BeforeEach(func() {
		// Setup in-memory database
		var err error
		db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
		Expect(err).NotTo(HaveOccurred())

		// Migrate the schema
		err = db.AutoMigrate(&model.Admin{})
		Expect(err).NotTo(HaveOccurred())

		// Initialize the repository
		adminRepo = repository.NewUserRepository(db)
	})

	Context("Create Admin", func() {
		It("should create a new Admin", func() {
			admin := &model.Admin{Username: "admin", Email: "admin@gmail.com", Password: "admin"}
			err := adminRepo.CreateUser(admin)
			Expect(err).NotTo(HaveOccurred())
			Expect(admin.ID).NotTo(BeZero())
		})
	})

	Context("Get Admin by Username", func() {
		It("should get Admin by username", func() {
			admin := &model.Admin{Username: "admin", Email: "admin@gmail.com", Password: "admin"}
			adminRepo.CreateUser(admin)

			fetchedAdmin, err := adminRepo.GetUserByUsername(admin.Username)
			Expect(err).NotTo(HaveOccurred())
			Expect(fetchedAdmin).NotTo(BeNil())
			Expect(fetchedAdmin.Username).To(Equal(admin.Username))
		})
	})

	Context("Get Admin by Email", func() {
		It("should get Admin by email", func() {
			admin := &model.Admin{Username: "admin", Email: "admin@gmail.com", Password: "admin"}
			adminRepo.CreateUser(admin)

			fetchedAdmin, err := adminRepo.GetUserByEmail(admin.Email)
			Expect(err).NotTo(HaveOccurred())
			Expect(fetchedAdmin).NotTo(BeNil())
			Expect(fetchedAdmin.Username).To(Equal(admin.Username))
		})
	})

	Context("Update Admin", func() {
		It("should update Admin", func() {
			admin := &model.Admin{Username: "admin", Email: "admin@gmail.com", Password: "admin"}
			adminRepo.CreateUser(admin)

			admin.Username = "new_admin"
			err := adminRepo.UpdateUser(admin)
			Expect(err).NotTo(HaveOccurred())

			fetchedAdmin, err := adminRepo.GetUserByUsername(admin.Username)
			Expect(err).NotTo(HaveOccurred())
			Expect(fetchedAdmin).NotTo(BeNil())
			Expect(fetchedAdmin.Username).To(Equal("new_admin"))
		})
	})

	Context("Delete Admin", func() {
		It("should delete Admin", func() {
			admin := &model.Admin{Username: "admin", Email: "admin@gmail.com", Password: "admin"}
			adminRepo.CreateUser(admin)

			err := adminRepo.DeleteUser(admin)
			Expect(err).NotTo(HaveOccurred())

			fetchedAdmin, err := adminRepo.GetUserByUsername("admin")
			Expect(err).To(HaveOccurred())
			Expect(fetchedAdmin).To(BeNil())
		})
	})
})
