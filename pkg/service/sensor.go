package service

import (
	"Go-gRPC-React-starter/pkg/model"
	"Go-gRPC-React-starter/gen/proto/sensor"
	"context"

	"github.com/bufbuild/connect-go"
)

type SensorService struct {
	// Add any fields if needed
}

func (s *SensorService) CreateSensor(ctx context.Context, req *connect.Request[sensor.CreateSensorRequest]) (*connect.Response[sensor.CreateSensorResponse], error) {
	newSensor, err := model.CreateSensor(req.Msg.Sensor)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&sensor.CreateSensorResponse{
		Sensor: newSensor,
	}), nil
}

func (s *SensorService) GetSensor(ctx context.Context, req *connect.Request[sensor.GetSensorRequest]) (*connect.Response[sensor.GetSensorResponse], error) {
	u, err := model.GetSensorFromDB(req.Msg.Id)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&sensor.GetSensorResponse{
		Sensor: u,
	}), nil
}

func (s *SensorService) UpdateSensor(ctx context.Context, req *connect.Request[sensor.UpdateSensorRequest]) (*connect.Response[sensor.UpdateSensorResponse], error) {
	updatedSensor, err := model.UpdateSensorInDB(req.Msg.Sensor)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&sensor.UpdateSensorResponse{
		Sensor: updatedSensor,
	}), nil
}

func (s *SensorService) DeleteSensor(ctx context.Context, req *connect.Request[sensor.DeleteSensorRequest]) (*connect.Response[sensor.DeleteSensorResponse], error) {
	err := model.DeleteSensorFromDB(req.Msg.Id)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&sensor.DeleteSensorResponse{
		Success: true,
	}), nil
}
