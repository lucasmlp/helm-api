package createKubeConfig

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/machado-br/helm-api/adapters/k8s/mocks"
	"github.com/machado-br/helm-api/services"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("Successful responses", func(t *testing.T) {

		adapterMock := mocks.NewMockAdapter(ctrl)

		adapterMock.EXPECT().RetrieveSecret().Return([]byte{}, nil)

		adapterMock.EXPECT().WriteToFile(gomock.Any()).Return(nil)

		service, err := NewService(adapterMock)
		if err != nil {
			t.Fatal(err)
		}

		err = service.Run()
		if err != nil {
			t.Fatal(err)
		}

		assert.NoError(t, err)
	})

	t.Run("Failure responses", func(t *testing.T) {
		ctx := context.Background()

		adapterErrorMock := errors.New("error-mock")

		tt := []struct {
			name        string
			errMock     error
			adapterMock func(context.Context, *gomock.Controller) *mocks.MockAdapter
		}{
			{
				name:    "Failed while retrieving k8s secret",
				errMock: services.ErrK8sSecrets,
				adapterMock: func(context.Context, *gomock.Controller) *mocks.MockAdapter {
					adapterMock := mocks.NewMockAdapter(ctrl)

					adapterMock.EXPECT().RetrieveSecret().Return(nil, adapterErrorMock)

					return adapterMock
				},
			},
			{
				name:    "Failed while writing kubeconfig file",
				errMock: services.ErrWriteKubeConfig,
				adapterMock: func(context.Context, *gomock.Controller) *mocks.MockAdapter {
					adapterMock := mocks.NewMockAdapter(ctrl)

					adapterMock.EXPECT().RetrieveSecret().Return([]byte{}, nil)
					adapterMock.EXPECT().WriteToFile(gomock.Any()).Return(adapterErrorMock)

					return adapterMock
				},
			},
		}

		for _, tc := range tt {
			t.Run(tc.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)

				service, err := NewService(tc.adapterMock(ctx, ctrl))
				if err != nil {
					t.Fatal(err)
				}

				err = service.Run()
				if err == nil {
					t.Fatalf("Should have failed by '%s', got nothing", tc.errMock.Error())
				}

				if err.Error() != tc.errMock.Error() {
					t.Fatalf("Should have failed by '%s', got '%s'", tc.errMock.Error(), err.Error())
				}
			})
		}
	})
}
