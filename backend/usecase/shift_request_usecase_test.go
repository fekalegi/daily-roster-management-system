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

var _ = Describe("ShiftRequestUsecase", func() {
	var (
		mockCtrl         *gomock.Controller
		mockRepo         *mock_repository.MockShiftRequestRepository
		mockShiftRepo    *mock_repository.MockShiftRepository
		mockAssignRepo   *mock_repository.MockAssignmentRepository
		usecaseUnderTest usecase.ShiftRequestUsecase
		req              *model.ShiftRequest
		shift            *model.Shift
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockRepo = mock_repository.NewMockShiftRequestRepository(mockCtrl)
		mockShiftRepo = mock_repository.NewMockShiftRepository(mockCtrl)
		mockAssignRepo = mock_repository.NewMockAssignmentRepository(mockCtrl)

		usecaseUnderTest = usecase.NewShiftRequestUsecase(mockRepo, mockShiftRepo, mockAssignRepo)

		req = &model.ShiftRequest{ID: 1, WorkerID: 2, ShiftID: 3, Status: "pending"}
		shift = &model.Shift{ID: 3, Date: "2025-05-20"}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("CreateRequest", func() {
		It("creates a valid shift request", func() {
			mockShiftRepo.EXPECT().GetByID(req.ShiftID).Return(shift, nil)
			mockAssignRepo.EXPECT().IsShiftAlreadyAssigned(req.ShiftID).Return(false, nil)
			mockAssignRepo.EXPECT().HasAssignmentOnDay(req.WorkerID, shift.Date).Return(false, nil)
			mockAssignRepo.EXPECT().CountAssignmentsInWeek(req.WorkerID, "2025-05-19", "2025-05-25").Return(2, nil)
			mockRepo.EXPECT().Create(req).Return(nil)

			err := usecaseUnderTest.CreateRequest(req)
			Expect(err).ToNot(HaveOccurred())
		})

		It("rejects if shift already assigned", func() {
			mockShiftRepo.EXPECT().GetByID(req.ShiftID).Return(shift, nil)
			mockAssignRepo.EXPECT().IsShiftAlreadyAssigned(req.ShiftID).Return(true, nil)

			err := usecaseUnderTest.CreateRequest(req)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("already assigned"))
		})
	})

	Describe("Approve", func() {
		It("approves a valid shift request", func() {
			mockRepo.EXPECT().GetByID(req.ID).Return(req, nil)
			mockShiftRepo.EXPECT().GetByID(req.ShiftID).Return(shift, nil)
			mockAssignRepo.EXPECT().IsShiftAlreadyAssigned(req.ShiftID).Return(false, nil)
			mockAssignRepo.EXPECT().HasAssignmentOnDay(req.WorkerID, shift.Date).Return(false, nil)
			mockAssignRepo.EXPECT().CountAssignmentsInWeek(req.WorkerID, "2025-05-19", "2025-05-25").Return(3, nil)
			mockAssignRepo.EXPECT().Create(gomock.Any()).Return(nil)
			mockRepo.EXPECT().UpdateStatus(req.ID, "approved").Return(nil)

			err := usecaseUnderTest.Approve(req.ID)
			Expect(err).ToNot(HaveOccurred())
		})

		It("returns error if assignment creation fails", func() {
			mockRepo.EXPECT().GetByID(req.ID).Return(req, nil)
			mockShiftRepo.EXPECT().GetByID(req.ShiftID).Return(shift, nil)
			mockAssignRepo.EXPECT().IsShiftAlreadyAssigned(req.ShiftID).Return(false, nil)
			mockAssignRepo.EXPECT().HasAssignmentOnDay(req.WorkerID, shift.Date).Return(false, nil)
			mockAssignRepo.EXPECT().CountAssignmentsInWeek(req.WorkerID, "2025-05-19", "2025-05-25").Return(2, nil)
			mockAssignRepo.EXPECT().Create(gomock.Any()).Return(errors.New("db error"))

			err := usecaseUnderTest.Approve(req.ID)
			Expect(err).To(HaveOccurred())
		})
	})

	Describe("Reject", func() {
		It("rejects a pending request", func() {
			mockRepo.EXPECT().GetByID(req.ID).Return(req, nil)
			mockRepo.EXPECT().UpdateStatus(req.ID, "rejected").Return(nil)

			err := usecaseUnderTest.Reject(req.ID)
			Expect(err).ToNot(HaveOccurred())
		})

		It("returns error if request is not pending", func() {
			nonPending := &model.ShiftRequest{ID: 99, Status: "approved"}
			mockRepo.EXPECT().GetByID(99).Return(nonPending, nil)

			err := usecaseUnderTest.Reject(99)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("not pending"))
		})
	})

	Describe("GetByStatus", func() {
		It("returns requests filtered by status", func() {
			mockRepo.EXPECT().GetWithDetailsFilterStatus("approved").Return([]model.ShiftRequestDetail{}, nil)
			result, err := usecaseUnderTest.GetByStatus("approved")
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(HaveLen(0))
		})
	})

	Describe("GetAll", func() {
		It("returns all shift requests", func() {
			mockRepo.EXPECT().GetWithDetails().Return([]model.ShiftRequestDetail{}, nil)
			result, err := usecaseUnderTest.GetAll()
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(HaveLen(0))
		})
	})
})
