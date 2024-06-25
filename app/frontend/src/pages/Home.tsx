import React from 'react';

const Home: React.FC = () => {
  return (
    <div className="container mx-auto p-4">
      <h1 className="text-3xl font-bold mb-4">Welcome to Rapid Web Development</h1>
      <p className="mb-4">
        This project leverages a single YAML file to generate the proto file, database initialization code, model, service file, and a React admin page for CRUD operations.
      </p>
      <p className="mb-4">
        By reducing boilerplate code, you can focus on business logic, making development faster and more efficient.
      </p>
      <p>
        Simply define your data schema and let the automation handle the rest. Get started by exploring the features and functionalities provided.
      </p>
    </div>
  );
};

export default Home;