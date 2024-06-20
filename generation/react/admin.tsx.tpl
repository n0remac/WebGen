import React, { useState } from 'react';
import { {{.ModelName | lower}}Service } from '../../services/{{.ModelName}}Service';
import { 
    Create{{.ModelName}}Request, 
    Get{{.ModelName}}Request, 
    Update{{.ModelName}}Request, 
    Delete{{.ModelName}}Request, 
    {{.ModelName}} 
} from '../../rpc/proto/{{.ModelName | lower}}/{{.ModelName | lower}}_pb';

const {{.ModelName}}Page = () => {
    const [{{.ModelName | lower}}, set{{.ModelName}}] = useState(new {{.ModelName}}());
    const [searchId, setSearchId] = useState<string>('');
    const [searchResult, setSearchResult] = useState<{{.ModelName}} | null>(null);
    const [status, setStatus] = useState<string>('');

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target;
        const parsedValue = {{ range .Fields }}
        {{ if isIntegerType .Type }}name === "{{.Name}}" || {{end}}
        {{ end }} false ? parseInt(value, 10) : value;
        set{{.ModelName}}(prev{{.ModelName}} => {
            const updated{{.ModelName}} = { ...prev{{.ModelName}} } as {{.ModelName}};
            (updated{{.ModelName}} as any)[name] = parsedValue;
            return updated{{.ModelName}};
        });
    };

    const handleCreate = async (e: React.FormEvent) => {
        e.preventDefault();
        const request = new Create{{.ModelName}}Request();
        request.{{.ModelName | lower}} = new {{.ModelName}}();
{{ range .Fields }}
        request.{{$.ModelName | lower}}.{{.Name}} = {{$.ModelName | lower}}.{{.Name}};
{{ end }}

        console.log('Creating {{.ModelName | lower}}:', request.{{.ModelName | lower}});
        try {
            await {{.ModelName | lower}}Service.create{{.ModelName}}(request, {});
            setStatus('{{.ModelName}} created successfully');
        } catch (error) {
            console.error('Error creating {{.ModelName | lower}}:', error);
            setStatus('{{.ModelName}} creation failed');
        }
    };

    const handleGet = async (e: React.FormEvent) => {
        e.preventDefault();
        const request = new Get{{.ModelName}}Request();
        request.id = parseInt(searchId, 10);
        try {
            const response = await {{.ModelName | lower}}Service.get{{.ModelName}}(request, {});
            console.log('Fetched {{.ModelName | lower}}:', response.{{.ModelName | lower}});
            setSearchResult(response.{{.ModelName | lower}} as {{.ModelName}});
        } catch (error) {
            console.error('Error fetching {{.ModelName | lower}}:', error);
            setSearchResult(null);
            setStatus('{{.ModelName}} fetch failed');
        }
    };

    const handleUpdate = async (e: React.FormEvent) => {
        e.preventDefault();
        const request = new Update{{.ModelName}}Request();
        request.{{.ModelName | lower}} = new {{.ModelName}}();
        request.{{.ModelName | lower}}.id = {{.ModelName | lower}}.id;
{{ range .Fields }}
        request.{{$.ModelName | lower}}.{{.Name}} = {{$.ModelName | lower}}.{{.Name}};
{{ end }}
        console.log('Updating {{.ModelName | lower}}:', request.{{.ModelName | lower}});
        try {
            await {{.ModelName | lower}}Service.update{{.ModelName}}(request, {});
            setStatus('{{.ModelName}} updated successfully');
        } catch (error) {
            console.error('Error updating {{.ModelName | lower}}:', error);
            setStatus('{{.ModelName}} update failed');
        }
    };

    const handleDelete = async (e: React.FormEvent) => {
        e.preventDefault();
        const request = new Delete{{.ModelName}}Request();
        request.id = {{.ModelName | lower}}.id;

        try {
            await {{.ModelName | lower}}Service.delete{{.ModelName}}(request, {});
            setStatus('{{.ModelName}} deleted successfully');
        } catch (error) {
            console.error('Error deleting {{.ModelName | lower}}:', error);
            setStatus('{{.ModelName}} deletion failed');
        }
    };

    return (
        <div className="container mx-auto p-4">
            <h1 className="text-2xl font-bold mb-4">{{.ModelName}} Management</h1>
            <form onSubmit={handleCreate}>
{{ range .Fields }}
                <input 
                    type="text" 
                    name="{{.Name}}" 
                    value={{wrapInBraces (printf "%s.%s" ($.ModelName | lower) .Name)}}
                    onChange={handleChange} 
                    placeholder="{{.Name | title}}" 
                />
{{ end }}
                <button type="submit">Create {{.ModelName}}</button>
            </form>
            <form onSubmit={handleGet}>
                <input 
                    type="text" 
                    value={searchId} 
                    onChange={e => setSearchId(e.target.value)} 
                    placeholder="{{.ModelName}} ID" 
                />
                <button type="submit">Get {{.ModelName}}</button>
            </form>
            {searchResult && (
                <div>
                    <h3>Search Result:</h3>
                    <p>ID: {searchResult.id}</p>
{{ range .Fields }}
                    <p>{{.Name | title}}: {searchResult.{{.Name}}}</p>
{{ end }}
                </div>
            )}
            <form onSubmit={handleUpdate}>
{{ range .Fields }}
                <input 
                    type="text" 
                    name="{{.Name}}" 
                    value={{wrapInBraces (printf "%s.%s" ($.ModelName | lower) .Name)}}
                    onChange={handleChange} 
                    placeholder="{{.Name | title}}" 
                />
{{ end }}
                <button type="submit">Update {{.ModelName}}</button>
            </form>
            <form onSubmit={handleDelete}>
                <input 
                    type="text" 
                    value={searchId} 
                    onChange={e => setSearchId(e.target.value)} 
                    placeholder="{{.ModelName}} ID" 
                />
                <button type="submit">Delete {{.ModelName}}</button>
            </form>
            {status && <p>{status}</p>}
        </div>
    );
};

export default {{.ModelName}}Page;