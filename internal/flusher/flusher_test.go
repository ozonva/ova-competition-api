package flusher

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"ozonva/ova-competition-api/internal/models"
	"ozonva/ova-competition-api/internal/repo"
	"time"
)

var _ = Describe("Flusher", func() {
	var (
		batchSize    = 3
		mockCtrl     *gomock.Controller
		mockRepo     *repo.MockRepo
		f            Flusher
		competitions = []models.Competition{
			models.NewCompetition(1, "n1", time.Now()),
			models.NewCompetition(2, "n2", time.Now()),
			models.NewCompetition(3, "n3", time.Now()),
			models.NewCompetition(4, "n4", time.Now()),
		}
		ctx = context.TODO()
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockRepo = repo.NewMockRepo(mockCtrl)
		f = NewFlusher(batchSize, mockRepo)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("Flush competitions", func() {
		When("can flush all competitions", func() {
			Context("competitions count less than batch size", func() {
				BeforeEach(func() {
					mockRepo.EXPECT().AddEntities(ctx, competitions[:batchSize-1]).Return(nil).Times(1)
				})
				It("all competitions should be flushed", func() {
					Expect(f.Flush(competitions[:batchSize-1])).To(BeNil())
				})
			})
			Context("competitions count more than batch size", func() {
				BeforeEach(func() {
					gomock.InOrder(
						mockRepo.EXPECT().AddEntities(ctx, competitions[:batchSize]).Return(nil).Times(1),
						mockRepo.EXPECT().AddEntities(ctx, competitions[batchSize:]).Return(nil).Times(1),
					)
				})
				It("there are no unflushed competitions", func() {
					Expect(f.Flush(competitions)).To(BeNil())
				})
			})
		})
		When("unable to flush", func() {
			err := errors.New("failed to save competitions")
			Context("some of the competitions", func() {
				BeforeEach(func() {
					gomock.InOrder(
						mockRepo.EXPECT().AddEntities(ctx, competitions[:batchSize]).Return(err).Times(1),
						mockRepo.EXPECT().AddEntities(ctx, competitions[batchSize:]).Return(nil).Times(1),
					)
				})
				It("only first batch of competitions is unflushed", func() {
					Expect(f.Flush(competitions)).To(Equal(competitions[:batchSize]))
				})
			})
			Context("all competitions", func() {
				BeforeEach(func() {
					gomock.InOrder(
						mockRepo.EXPECT().AddEntities(ctx, competitions[:batchSize]).Return(err).Times(1),
						mockRepo.EXPECT().AddEntities(ctx, competitions[batchSize:]).Return(err).Times(1),
					)
				})
				It("all competitions are unflushed", func() {
					Expect(f.Flush(competitions)).To(Equal(competitions))
				})
			})
		})
	})
})
