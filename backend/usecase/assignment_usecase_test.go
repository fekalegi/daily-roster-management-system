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

var errSomething = errors.New("something error")

var _ = Describe("AssignmentUsecase", func() {
	var (
		mockCtrl              *gomock.Controller
		userUC                usecase.AssignmentUsecase
		repo                  *mock_repository.MockAssignmentRepository
		shiftRepo             *mock_repository.MockShiftRepository
		mockAssignment        *model.Assignment
		mockAssignmentDetail  *model.AssignmentDetail
		mockAssignmentDetails []model.AssignmentDetail
		mockShift             *model.Shift
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockCtrl.Finish()
		repo = mock_repository.NewMockAssignmentRepository(mockCtrl)
		shiftRepo = mock_repository.NewMockShiftRepository(mockCtrl)
		userUC = usecase.NewAssignmentUsecase(repo, shiftRepo)
		mockAssignment = &model.Assignment{
			ID:       1,
			ShiftID:  1,
			WorkerID: 1,
		}

		mockAssignmentDetail = &model.AssignmentDetail{
			ID: 1,
			Shift: &model.Shift{
				ID: 1,
			},
			Worker: &model.Worker{
				ID:   1,
				Name: "John",
			},
		}

		mockAssignmentDetails = []model.AssignmentDetail{*mockAssignmentDetail}

		mockShift = &model.Shift{
			ID: 1,
		}
	})

	checkShiftRequestSuccess := func() {
		repo.EXPECT().IsShiftAlreadyAssigned(gomock.Any()).Return(false, nil)
		repo.EXPECT().HasAssignmentOnDay(gomock.Any(), gomock.Any()).Return(false, nil)
		repo.EXPECT().CountAssignmentsInWeek(gomock.Any(), gomock.Any(), gomock.Any()).Return(3, nil)
	}

	Describe("Assign", func() {
		It("return success", func() {
			shiftRepo.EXPECT().GetByID(gomock.Any()).Return(mockShift, nil)
			checkShiftRequestSuccess()
			repo.EXPECT().Create(gomock.Any()).Return(nil)
			err := userUC.Assign(mockAssignment)
			Expect(err).Should(Succeed())
		})

		It("return error", func() {
			shiftRepo.EXPECT().GetByID(gomock.Any()).Return(mockShift, nil)
			checkShiftRequestSuccess()
			repo.EXPECT().Create(gomock.Any()).Return(errSomething)
			err := userUC.Assign(mockAssignment)
			Expect(err).Should(HaveOccurred())
		})
	})

	Describe("GetByWorker", func() {
		It("return success", func() {
			repo.EXPECT().FindByWorkerID(gomock.Any()).Return(mockAssignmentDetails, nil)
			data, err := userUC.GetByWorker(1)
			Expect(err).Should(Succeed())
			Expect(data).Should(Equal(mockAssignmentDetails))
		})

		It("return error", func() {
			repo.EXPECT().FindByWorkerID(gomock.Any()).Return(nil, errSomething)
			_, err := userUC.GetByWorker(1)
			Expect(err).Should(HaveOccurred())
		})
	})

	Describe("GetByDate", func() {
		It("return success", func() {
			repo.EXPECT().FindByDate(gomock.Any()).Return(mockAssignmentDetails, nil)
			data, err := userUC.GetByDate("")
			Expect(err).Should(Succeed())
			Expect(data).Should(Equal(mockAssignmentDetails))
		})

		It("return error", func() {
			repo.EXPECT().FindByDate(gomock.Any()).Return(nil, errSomething)
			_, err := userUC.GetByDate("")
			Expect(err).Should(HaveOccurred())
		})
	})

	Describe("FindAll", func() {
		It("return success", func() {
			repo.EXPECT().FindAll().Return(mockAssignmentDetails, nil)
			data, err := userUC.FindAll()
			Expect(err).Should(Succeed())
			Expect(data).Should(Equal(mockAssignmentDetails))
		})

		It("return error", func() {
			repo.EXPECT().FindAll().Return(nil, errSomething)
			_, err := userUC.FindAll()
			Expect(err).Should(HaveOccurred())
		})
	})
})
