import React, { useState } from 'react';
import { sensorService } from '../../services/SensorService';
import { 
    CreateSensorRequest, 
    GetSensorRequest, 
    UpdateSensorRequest, 
    DeleteSensorRequest, 
    Sensor 
} from '../../rpc/proto/sensor/sensor_pb';

const SensorPage = () => {
    const [sensor, setSensor] = useState(new Sensor());
    const [searchId, setSearchId] = useState<string>('');
    const [searchResult, setSearchResult] = useState<Sensor | null>(null);
    const [status, setStatus] = useState<string>('');

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target;
        const parsedValue = 
        name === "id" || 
        
        name === "temperature" || 
        
        name === "humidity" || 
        
        name === "distance" || 
        
        name === "light" || 
        
        name === "sound" || 
         false ? parseInt(value, 10) : value;
        setSensor(prevSensor => {
            const updatedSensor = { ...prevSensor } as Sensor;
            (updatedSensor as any)[name] = parsedValue;
            return updatedSensor;
        });
    };

    const handleCreate = async (e: React.FormEvent) => {
        e.preventDefault();
        const request = new CreateSensorRequest();
        request.sensor = new Sensor();

        request.sensor.id = sensor.id;

        request.sensor.temperature = sensor.temperature;

        request.sensor.humidity = sensor.humidity;

        request.sensor.distance = sensor.distance;

        request.sensor.light = sensor.light;

        request.sensor.sound = sensor.sound;


        console.log('Creating sensor:', request.sensor);
        try {
            await sensorService.createSensor(request, {});
            setStatus('Sensor created successfully');
        } catch (error) {
            console.error('Error creating sensor:', error);
            setStatus('Sensor creation failed');
        }
    };

    const handleGet = async (e: React.FormEvent) => {
        e.preventDefault();
        const request = new GetSensorRequest();
        request.id = parseInt(searchId, 10);
        try {
            const response = await sensorService.getSensor(request, {});
            console.log('Fetched sensor:', response.sensor);
            setSearchResult(response.sensor as Sensor);
        } catch (error) {
            console.error('Error fetching sensor:', error);
            setSearchResult(null);
            setStatus('Sensor fetch failed');
        }
    };

    const handleUpdate = async (e: React.FormEvent) => {
        e.preventDefault();
        const request = new UpdateSensorRequest();
        request.sensor = new Sensor();
        request.sensor.id = sensor.id;

        request.sensor.id = sensor.id;

        request.sensor.temperature = sensor.temperature;

        request.sensor.humidity = sensor.humidity;

        request.sensor.distance = sensor.distance;

        request.sensor.light = sensor.light;

        request.sensor.sound = sensor.sound;

        console.log('Updating sensor:', request.sensor);
        try {
            await sensorService.updateSensor(request, {});
            setStatus('Sensor updated successfully');
        } catch (error) {
            console.error('Error updating sensor:', error);
            setStatus('Sensor update failed');
        }
    };

    const handleDelete = async (e: React.FormEvent) => {
        e.preventDefault();
        const request = new DeleteSensorRequest();
        request.id = sensor.id;

        try {
            await sensorService.deleteSensor(request, {});
            setStatus('Sensor deleted successfully');
        } catch (error) {
            console.error('Error deleting sensor:', error);
            setStatus('Sensor deletion failed');
        }
    };

    return (
        <div className="container mx-auto p-4">
            <h1 className="text-2xl font-bold mb-4">Sensor Management</h1>
            <form onSubmit={handleCreate}>

                <input 
                    type="text" 
                    name="id" 
                    value={sensor.id}
                    onChange={handleChange} 
                    placeholder="Id" 
                />

                <input 
                    type="text" 
                    name="temperature" 
                    value={sensor.temperature}
                    onChange={handleChange} 
                    placeholder="Temperature" 
                />

                <input 
                    type="text" 
                    name="humidity" 
                    value={sensor.humidity}
                    onChange={handleChange} 
                    placeholder="Humidity" 
                />

                <input 
                    type="text" 
                    name="distance" 
                    value={sensor.distance}
                    onChange={handleChange} 
                    placeholder="Distance" 
                />

                <input 
                    type="text" 
                    name="light" 
                    value={sensor.light}
                    onChange={handleChange} 
                    placeholder="Light" 
                />

                <input 
                    type="text" 
                    name="sound" 
                    value={sensor.sound}
                    onChange={handleChange} 
                    placeholder="Sound" 
                />

                <button type="submit">Create Sensor</button>
            </form>
            <form onSubmit={handleGet}>
                <input 
                    type="text" 
                    value={searchId} 
                    onChange={e => setSearchId(e.target.value)} 
                    placeholder="Sensor ID" 
                />
                <button type="submit">Get Sensor</button>
            </form>
            {searchResult && (
                <div>
                    <h3>Search Result:</h3>
                    <p>ID: {searchResult.id}</p>

                    <p>Id: {searchResult.id}</p>

                    <p>Temperature: {searchResult.temperature}</p>

                    <p>Humidity: {searchResult.humidity}</p>

                    <p>Distance: {searchResult.distance}</p>

                    <p>Light: {searchResult.light}</p>

                    <p>Sound: {searchResult.sound}</p>

                </div>
            )}
            <form onSubmit={handleUpdate}>

                <input 
                    type="text" 
                    name="id" 
                    value={sensor.id}
                    onChange={handleChange} 
                    placeholder="Id" 
                />

                <input 
                    type="text" 
                    name="temperature" 
                    value={sensor.temperature}
                    onChange={handleChange} 
                    placeholder="Temperature" 
                />

                <input 
                    type="text" 
                    name="humidity" 
                    value={sensor.humidity}
                    onChange={handleChange} 
                    placeholder="Humidity" 
                />

                <input 
                    type="text" 
                    name="distance" 
                    value={sensor.distance}
                    onChange={handleChange} 
                    placeholder="Distance" 
                />

                <input 
                    type="text" 
                    name="light" 
                    value={sensor.light}
                    onChange={handleChange} 
                    placeholder="Light" 
                />

                <input 
                    type="text" 
                    name="sound" 
                    value={sensor.sound}
                    onChange={handleChange} 
                    placeholder="Sound" 
                />

                <button type="submit">Update Sensor</button>
            </form>
            <form onSubmit={handleDelete}>
                <input 
                    type="text" 
                    value={searchId} 
                    onChange={e => setSearchId(e.target.value)} 
                    placeholder="Sensor ID" 
                />
                <button type="submit">Delete Sensor</button>
            </form>
            {status && <p>{status}</p>}
        </div>
    );
};

export default SensorPage;