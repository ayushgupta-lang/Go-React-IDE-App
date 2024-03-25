import React, { useState } from 'react';
import axios from 'axios';
import './IDE.css'; 

function IDE() {
  const [code, setCode] = useState('');
  const [language, setLanguage] = useState('');
  const [output, setOutput] = useState('');
  const [error, setError] = useState('');

  const handleSubmit = async () => {
    try {
      const response = await axios.post('http://localhost:8080/api/run', { language, code });
      setOutput(response.data.output);
      setError(response.data.error);
    } catch (error) {
      console.error('Error:', error);
      setError('An error occurred while processing the request.');
    }
  };

  return (
    <div className="container"> 
      <h1 className="heading">IDE</h1>
      <div className="inputContainer">
        <label className="label">Language:</label><br></br><br></br>
        <select className="select" value={language} onChange={(e) => setLanguage(e.target.value)}>
          <option value="python">Python</option>
          <option value="go">Go</option>
          <option value="javascript">JavaScript</option>
        </select>
      </div>
      <div className="inputContainer">
        <label className="label">Code:</label>
        <textarea className="textarea" value={code} onChange={(e) => setCode(e.target.value)} />
      </div>
      <button className="button" onClick={handleSubmit}>Run Code</button>
      <div className='mine'>
        <h2>Output:</h2>
        <pre>{output}</pre>
      </div>
      <div className='mine'>
        <h2>Error:</h2>
        <pre>{error}</pre>
      </div>
    </div>
  );
}

export default IDE;
