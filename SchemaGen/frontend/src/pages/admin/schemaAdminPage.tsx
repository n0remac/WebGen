import React, { useState } from 'react';
import { schemaService } from '../../services/SchemaService';
import { 
    CreateSchemaRequest, 
    GetSchemaRequest, 
    UpdateSchemaRequest, 
    DeleteSchemaRequest, 
    Schema 
} from '../../rpc/proto/schema/schema_pb';

const SchemaPage = () => {
    const [schema, setSchema] = useState(new Schema());
    const [searchId, setSearchId] = useState<string>('');
    const [searchResult, setSearchResult] = useState<Schema | null>(null);
    const [status, setStatus] = useState<string>('');

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target;
        const parsedValue = 
        name === "id" || 
        
        
        
        
         false ? parseInt(value, 10) : value;
        setSchema(prevSchema => {
            const updatedSchema = { ...prevSchema } as Schema;
            (updatedSchema as any)[name] = parsedValue;
            return updatedSchema;
        });
    };

    const handleCreate = async (e: React.FormEvent) => {
        e.preventDefault();
        const request = new CreateSchemaRequest();
        request.schema = new Schema();

        request.schema.id = schema.id;

        request.schema.name = schema.name;

        request.schema.content = schema.content;


        console.log('Creating schema:', request.schema);
        try {
            await schemaService.createSchema(request, {});
            setStatus('Schema created successfully');
        } catch (error) {
            console.error('Error creating schema:', error);
            setStatus('Schema creation failed');
        }
    };

    const handleGet = async (e: React.FormEvent) => {
        e.preventDefault();
        const request = new GetSchemaRequest();
        request.id = parseInt(searchId, 10);
        try {
            const response = await schemaService.getSchema(request, {});
            console.log('Fetched schema:', response.schema);
            setSearchResult(response.schema as Schema);
        } catch (error) {
            console.error('Error fetching schema:', error);
            setSearchResult(null);
            setStatus('Schema fetch failed');
        }
    };

    const handleUpdate = async (e: React.FormEvent) => {
        e.preventDefault();
        const request = new UpdateSchemaRequest();
        request.schema = new Schema();
        request.schema.id = schema.id;

        request.schema.id = schema.id;

        request.schema.name = schema.name;

        request.schema.content = schema.content;

        console.log('Updating schema:', request.schema);
        try {
            await schemaService.updateSchema(request, {});
            setStatus('Schema updated successfully');
        } catch (error) {
            console.error('Error updating schema:', error);
            setStatus('Schema update failed');
        }
    };

    const handleDelete = async (e: React.FormEvent) => {
        e.preventDefault();
        const request = new DeleteSchemaRequest();
        request.id = schema.id;

        try {
            await schemaService.deleteSchema(request, {});
            setStatus('Schema deleted successfully');
        } catch (error) {
            console.error('Error deleting schema:', error);
            setStatus('Schema deletion failed');
        }
    };

    return (
        <div className="container mx-auto p-4">
            <h1 className="text-2xl font-bold mb-4">Schema Management</h1>
            <form onSubmit={handleCreate}>

                <input 
                    type="text" 
                    name="id" 
                    value={schema.id}
                    onChange={handleChange} 
                    placeholder="Id" 
                />

                <input 
                    type="text" 
                    name="name" 
                    value={schema.name}
                    onChange={handleChange} 
                    placeholder="Name" 
                />

                <input 
                    type="text" 
                    name="content" 
                    value={schema.content}
                    onChange={handleChange} 
                    placeholder="Content" 
                />

                <button type="submit">Create Schema</button>
            </form>
            <form onSubmit={handleGet}>
                <input 
                    type="text" 
                    value={searchId} 
                    onChange={e => setSearchId(e.target.value)} 
                    placeholder="Schema ID" 
                />
                <button type="submit">Get Schema</button>
            </form>
            {searchResult && (
                <div>
                    <h3>Search Result:</h3>
                    <p>ID: {searchResult.id}</p>

                    <p>Id: {searchResult.id}</p>

                    <p>Name: {searchResult.name}</p>

                    <p>Content: {searchResult.content}</p>

                </div>
            )}
            <form onSubmit={handleUpdate}>

                <input 
                    type="text" 
                    name="id" 
                    value={schema.id}
                    onChange={handleChange} 
                    placeholder="Id" 
                />

                <input 
                    type="text" 
                    name="name" 
                    value={schema.name}
                    onChange={handleChange} 
                    placeholder="Name" 
                />

                <input 
                    type="text" 
                    name="content" 
                    value={schema.content}
                    onChange={handleChange} 
                    placeholder="Content" 
                />

                <button type="submit">Update Schema</button>
            </form>
            <form onSubmit={handleDelete}>
                <input 
                    type="text" 
                    value={searchId} 
                    onChange={e => setSearchId(e.target.value)} 
                    placeholder="Schema ID" 
                />
                <button type="submit">Delete Schema</button>
            </form>
            {status && <p>{status}</p>}
        </div>
    );
};

export default SchemaPage;