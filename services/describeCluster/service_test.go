package describeCluster

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/machado-br/helm-api/adapters/aws/mocks"
	"github.com/machado-br/helm-api/adapters/models"
	"github.com/machado-br/helm-api/services"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("Successful responses", func(t *testing.T) {

		adapterMock := mocks.NewMockAdapter(ctrl)

		adapterMock.EXPECT().DescribeCluster().Return(models.Cluster{}, nil)

		service, err := NewService(adapterMock)
		if err != nil {
			t.Fatal(err)
		}

		cluster, err := service.Run()
		if err != nil {
			t.Fatal(err)
		}

		assert.NoError(t, err)
		assert.NotNil(t, cluster)
	})

	t.Run("Failure responses", func(t *testing.T) {

		adapterMock := mocks.NewMockAdapter(ctrl)

		errM := errors.New("mock-error")
		adapterMock.EXPECT().DescribeCluster().Return(models.Cluster{}, errM)

		service, err := NewService(adapterMock)
		if err != nil {
			t.Fatal(err)
		}

		_, err = service.Run()
		if err == nil {
			t.Fatalf("Should have failed by '%s', got nothing", services.ErrGetClusterInfo)
		}

		if err.Error() != services.ErrGetClusterInfo.Error() {
			t.Fatalf("Should have failed by '%s', got '%s'", services.ErrGetClusterInfo, err.Error())
		}
	})
}
