import React from 'react';
import IDE from './components/IDE'; 

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <h1 style={{ textAlign: 'center' }}>My IDE App</h1>
        <IDE /> 
      </header>
    </div>
  );
}

export default App;
