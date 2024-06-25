import { BrowserRouter as Router, Link } from 'react-router-dom';
import AppRouter from './AppRouter';

export default function App() {
    return (
        <Router>
            <div>
                <AppRouter />
            </div>
        </Router>
    );
}
