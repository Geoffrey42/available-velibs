import React from 'react';
import './App.css';
import AppHeader from './components/AppHeader';
import ResultsTable from './components/ResultsTable';

function App() {
    return (
        <div className="App">
            <AppHeader />
            <ResultsTable />
        </div>
    );
}

export default App;
