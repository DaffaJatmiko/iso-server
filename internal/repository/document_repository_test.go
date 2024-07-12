package repository_test

import (
	"github.com/DaffaJatmiko/project-iso/internal/model"
	"github.com/DaffaJatmiko/project-iso/internal/repository"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var _ = Describe("Document Repository", func() {
	var (
		db           *gorm.DB
		documentRepo repository.DocumentRepository
	)

	BeforeEach(func() {
		// Setup in-memory database
		var err error
		db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
		Expect(err).NotTo(HaveOccurred())

		// Migrate the schema
		err = db.AutoMigrate(&model.Document{})
		Expect(err).NotTo(HaveOccurred())

		documentRepo = repository.NewDocumentRepository(db)
	})

	Context("Create Document", func() {
		It("should create a new document", func() {
			doc := &model.Document{FileName: "Test Document", Link: "Test Document Link"}
			err := documentRepo.CreateDocument(doc)
			Expect(err).NotTo(HaveOccurred())
			Expect(doc.ID).NotTo(BeZero())
		})
	})

	Context("Get Document", func() {
		It("should find a document", func() {
			doc1 := &model.Document{FileName: "Test Document 1", Link: "Test Document Link 1"}
			doc2 := &model.Document{FileName: "Test Document 2", Link: "Test Document Link 2"}
			documentRepo.CreateDocument(doc1)
			documentRepo.CreateDocument(doc2)

			docs, err := documentRepo.GetDocuments()
			Expect(err).NotTo(HaveOccurred())
			Expect(docs).To(HaveLen(2))
		})
	})

	Context("Update Document", func() {
		It("should update document", func() {
			doc := &model.Document{FileName: "Test Document", Link: "Test Document Link"}
			documentRepo.CreateDocument(doc)

			doc.FileName = "Updated Test Document"
			err := documentRepo.UpdateDocument(doc)
			Expect(err).NotTo(HaveOccurred())

			fetchedDoc, err := documentRepo.GetDocumentById(doc.ID)
			Expect(err).NotTo(HaveOccurred())
			Expect(fetchedDoc.FileName).To(Equal("Updated Test Document"))

		})
	})

	Context("Delete Document", func() {
		It("should delete document", func() {
			doc := &model.Document{FileName: "Test Document", Link: "Test Document Link"}
			documentRepo.CreateDocument(doc)

			err := documentRepo.DeleteDocument(doc.ID)
			Expect(err).NotTo(HaveOccurred())

			fetchedDoc, err := documentRepo.GetDocumentById(doc.ID)
			Expect(err).To(HaveOccurred())
			Expect(fetchedDoc).To(BeNil())
		})
	})

})
