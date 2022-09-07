package uninstallRelease

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/machado-br/helm-api/adapters/helm/mocks"
	"github.com/machado-br/helm-api/services"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("Successful responses", func(t *testing.T) {

		adapterMock := mocks.NewMockAdapter(ctrl)

		adapterMock.EXPECT().UninstallRelease(gomock.Any(), gomock.Any()).Return(nil)

		service, err := NewService(adapterMock)
		if err != nil {
			t.Fatal(err)
		}

		err = service.Run("", false)
		if err != nil {
			t.Fatal(err)
		}

		assert.NoError(t, err)
	})

	t.Run("Failure responses", func(t *testing.T) {

		adapterMock := mocks.NewMockAdapter(ctrl)

		errM := errors.New("mock-error")
		adapterMock.EXPECT().UninstallRelease(gomock.Any(), gomock.Any()).Return(errM)

		service, err := NewService(adapterMock)
		if err != nil {
			t.Fatal(err)
		}

		err = service.Run("", false)
		if err == nil {
			t.Fatalf("Should have failed by '%s', got nothing", services.ErrUninstallRelease)
		}

		if err.Error() != services.ErrUninstallRelease.Error() {
			t.Fatalf("Should have failed by '%s', got '%s'", services.ErrUninstallRelease, err.Error())
		}
	})
}
