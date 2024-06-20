import { SensorService } from "../rpc/proto/sensor/sensor_connect";
import { transport } from "./service";
import { createPromiseClient } from "@bufbuild/connect";

export const sensorService = createPromiseClient(SensorService, transport);
