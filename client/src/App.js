import React from 'react';
import './App.css';
import AppFooter from './components/AppFooter';
import AppHeader from './components/AppHeader';
import ResultsTable from './components/ResultsTable';

function App() {
    return (
        <div className="App">
            <AppHeader />
            <ResultsTable />
            <AppFooter />
        </div>
    );
}

export default App;
