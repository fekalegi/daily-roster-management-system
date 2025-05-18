package usecase_test

import (
	mock_repository "daily-worker-roster-management-system/mocks/repository"
	model "daily-worker-roster-management-system/models"
	"daily-worker-roster-management-system/usecase"
	"errors"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UserUsecase", func() {
	var (
		mockCtrl   *gomock.Controller
		workerRepo *mock_repository.MockWorkerRepository
		userUC     usecase.UserUsecase
		mockWorker *model.Worker
		errExample = errors.New("worker not found")
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		workerRepo = mock_repository.NewMockWorkerRepository(mockCtrl)
		userUC = usecase.NewUserUsecase(workerRepo)

		mockWorker = &model.Worker{
			ID:   123,
			Name: "johndoe",
		}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("ValidateCredentials", func() {
		Context("when admin credentials are valid", func() {
			It("returns admin user", func() {
				user, err := userUC.ValidateCredentials("admin", "password")
				Expect(err).ToNot(HaveOccurred())
				Expect(user).To(Equal(&model.User{
					RoleId:   1,
					UserId:   0,
					Username: "Admin",
				}))
			})
		})

		Context("when worker credentials are valid", func() {
			It("returns the worker user", func() {
				workerRepo.EXPECT().GetByName("johndoe").Return(mockWorker, nil)

				user, err := userUC.ValidateCredentials("johndoe", "password")
				Expect(err).ToNot(HaveOccurred())
				Expect(user).To(Equal(&model.User{
					RoleId:   2,
					UserId:   123,
					Username: "johndoe",
				}))
			})
		})

		Context("when worker is not found", func() {
			It("returns error", func() {
				workerRepo.EXPECT().GetByName("janedoe").Return(nil, errExample)

				user, err := userUC.ValidateCredentials("janedoe", "password")
				Expect(err).To(HaveOccurred())
				Expect(user).To(BeNil())
			})
		})

		Context("when password is wrong", func() {
			It("returns error", func() {
				user, err := userUC.ValidateCredentials("johndoe", "wrongpass")
				Expect(err).To(HaveOccurred())
				Expect(user).To(BeNil())
			})
		})
	})
})
