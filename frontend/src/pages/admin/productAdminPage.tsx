import React, { useState } from 'react';
import { productService } from '../../services/ProductService';
import { 
    CreateProductRequest, 
    GetProductRequest, 
    UpdateProductRequest, 
    DeleteProductRequest, 
    Product 
} from '../../rpc/proto/product/product_pb';

const ProductPage = () => {
    const [product, setProduct] = useState(new Product());
    const [searchId, setSearchId] = useState<string>('');
    const [searchResult, setSearchResult] = useState<Product | null>(null);
    const [status, setStatus] = useState<string>('');

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target;
        const parsedValue = 
        name === "id" || 
        
        
        
        name === "amount" || 
        
        
         false ? parseInt(value, 10) : value;
        setProduct(prevProduct => {
            const updatedProduct = { ...prevProduct } as Product;
            (updatedProduct as any)[name] = parsedValue;
            return updatedProduct;
        });
    };

    const handleCreate = async (e: React.FormEvent) => {
        e.preventDefault();
        const request = new CreateProductRequest();
        request.product = new Product();

        request.product.id = product.id;

        request.product.name = product.name;

        request.product.amount = product.amount;

        request.product.description = product.description;


        console.log('Creating product:', request.product);
        try {
            await productService.createProduct(request, {});
            setStatus('Product created successfully');
        } catch (error) {
            console.error('Error creating product:', error);
            setStatus('Product creation failed');
        }
    };

    const handleGet = async (e: React.FormEvent) => {
        e.preventDefault();
        const request = new GetProductRequest();
        request.id = parseInt(searchId, 10);
        try {
            const response = await productService.getProduct(request, {});
            console.log('Fetched product:', response.product);
            setSearchResult(response.product as Product);
        } catch (error) {
            console.error('Error fetching product:', error);
            setSearchResult(null);
            setStatus('Product fetch failed');
        }
    };

    const handleUpdate = async (e: React.FormEvent) => {
        e.preventDefault();
        const request = new UpdateProductRequest();
        request.product = new Product();
        request.product.id = product.id;

        request.product.id = product.id;

        request.product.name = product.name;

        request.product.amount = product.amount;

        request.product.description = product.description;

        console.log('Updating product:', request.product);
        try {
            await productService.updateProduct(request, {});
            setStatus('Product updated successfully');
        } catch (error) {
            console.error('Error updating product:', error);
            setStatus('Product update failed');
        }
    };

    const handleDelete = async (e: React.FormEvent) => {
        e.preventDefault();
        const request = new DeleteProductRequest();
        request.id = product.id;

        try {
            await productService.deleteProduct(request, {});
            setStatus('Product deleted successfully');
        } catch (error) {
            console.error('Error deleting product:', error);
            setStatus('Product deletion failed');
        }
    };

    return (
        <div className="container mx-auto p-4">
            <h1 className="text-2xl font-bold mb-4">Product Management</h1>
            <form onSubmit={handleCreate}>

                <input 
                    type="text" 
                    name="id" 
                    value={product.id}
                    onChange={handleChange} 
                    placeholder="Id" 
                />

                <input 
                    type="text" 
                    name="name" 
                    value={product.name}
                    onChange={handleChange} 
                    placeholder="Name" 
                />

                <input 
                    type="text" 
                    name="amount" 
                    value={product.amount}
                    onChange={handleChange} 
                    placeholder="Amount" 
                />

                <input 
                    type="text" 
                    name="description" 
                    value={product.description}
                    onChange={handleChange} 
                    placeholder="Description" 
                />

                <button type="submit">Create Product</button>
            </form>
            <form onSubmit={handleGet}>
                <input 
                    type="text" 
                    value={searchId} 
                    onChange={e => setSearchId(e.target.value)} 
                    placeholder="Product ID" 
                />
                <button type="submit">Get Product</button>
            </form>
            {searchResult && (
                <div>
                    <h3>Search Result:</h3>
                    <p>ID: {searchResult.id}</p>

                    <p>Id: {searchResult.id}</p>

                    <p>Name: {searchResult.name}</p>

                    <p>Amount: {searchResult.amount}</p>

                    <p>Description: {searchResult.description}</p>

                </div>
            )}
            <form onSubmit={handleUpdate}>

                <input 
                    type="text" 
                    name="id" 
                    value={product.id}
                    onChange={handleChange} 
                    placeholder="Id" 
                />

                <input 
                    type="text" 
                    name="name" 
                    value={product.name}
                    onChange={handleChange} 
                    placeholder="Name" 
                />

                <input 
                    type="text" 
                    name="amount" 
                    value={product.amount}
                    onChange={handleChange} 
                    placeholder="Amount" 
                />

                <input 
                    type="text" 
                    name="description" 
                    value={product.description}
                    onChange={handleChange} 
                    placeholder="Description" 
                />

                <button type="submit">Update Product</button>
            </form>
            <form onSubmit={handleDelete}>
                <input 
                    type="text" 
                    value={searchId} 
                    onChange={e => setSearchId(e.target.value)} 
                    placeholder="Product ID" 
                />
                <button type="submit">Delete Product</button>
            </form>
            {status && <p>{status}</p>}
        </div>
    );
};

export default ProductPage;