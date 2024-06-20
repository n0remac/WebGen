import React from 'react';
import { Route, Routes } from 'react-router-dom';
import { routes } from './pages';
import Home from './pages/Home';

interface RouteConfig {
  path: string;
  component: React.ComponentType<any>;
}

const AppRouter: React.FC = () => (
  <Routes>
    {routes.map(({ path, component: Component }: RouteConfig, index: number) => (
      <Route key={index} path={path} element={<Component />} />
    ))}
    <Route path="/" element={<Home />} />
  </Routes>
);

export default AppRouter;
