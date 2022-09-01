package listReleases

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/machado-br/helm-api/adapters/helm/mocks"
	"github.com/machado-br/helm-api/adapters/models"
	"github.com/machado-br/helm-api/services"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("Successful responses", func(t *testing.T) {

		adapterMock := mocks.NewMockAdapter(ctrl)

		adapterMock.EXPECT().ListReleases().Return([]models.Release{}, nil)

		service, err := NewService(adapterMock)
		if err != nil {
			t.Fatal(err)
		}

		releases, err := service.Run()
		if err != nil {
			t.Fatal(err)
		}

		assert.NoError(t, err)
		assert.NotNil(t, releases)
	})

	t.Run("Failure responses", func(t *testing.T) {

		adapterMock := mocks.NewMockAdapter(ctrl)

		errM := errors.New("mock-error")
		adapterMock.EXPECT().ListReleases().Return([]models.Release{}, errM)

		service, err := NewService(adapterMock)
		if err != nil {
			t.Fatal(err)
		}

		_, err = service.Run()
		if err == nil {
			t.Fatalf("Should have failed by '%s', got nothing", services.ErrListReleases)
		}

		if err.Error() != services.ErrListReleases.Error() {
			t.Fatalf("Should have failed by '%s', got '%s'", services.ErrListReleases, err.Error())
		}
	})
}
