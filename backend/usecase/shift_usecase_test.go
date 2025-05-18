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

var _ = Describe("ShiftUsecase", func() {
	var (
		mockCtrl  *gomock.Controller
		shiftRepo *mock_repository.MockShiftRepository
		uc        usecase.ShiftUsecase
		mockShift *model.Shift
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		shiftRepo = mock_repository.NewMockShiftRepository(mockCtrl)
		uc = usecase.NewShiftUsecase(shiftRepo)

		mockShift = &model.Shift{
			ID:   1,
			Role: "Morning Shift",
		}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("GetShifts", func() {
		It("returns list of shifts", func() {
			shifts := []model.Shift{*mockShift}
			shiftRepo.EXPECT().GetAll().Return(shifts, nil)

			result, err := uc.GetShifts()
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(Equal(shifts))
		})

		It("returns error if repository fails", func() {
			shiftRepo.EXPECT().GetAll().Return(nil, errors.New("get error"))

			result, err := uc.GetShifts()
			Expect(err).To(HaveOccurred())
			Expect(result).To(BeNil())
		})
	})

	Describe("CreateShift", func() {
		It("creates a shift successfully", func() {
			shiftRepo.EXPECT().Create(mockShift).Return(nil)

			err := uc.CreateShift(mockShift)
			Expect(err).ToNot(HaveOccurred())
		})

		It("returns error if repository fails", func() {
			shiftRepo.EXPECT().Create(mockShift).Return(errors.New("create error"))

			err := uc.CreateShift(mockShift)
			Expect(err).To(HaveOccurred())
		})
	})
})
