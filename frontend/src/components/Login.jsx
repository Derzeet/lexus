import React, { useState } from 'react';

import axios from 'axios';

function Login() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleSubmit = async (event) => {
    event.preventDefault();
    try {
      const response = await axios.post('http://localhost:8000/login', { email, password });
      if (response.data.account && response.data.account.token) {
        localStorage.setItem('token', response.data.account.token);

        console.log(localStorage.getItem('token'))
        axios.defaults.headers.common['Authorization'] = response.data.account.token;

        window.location.reload();

        
      }
    } catch (error) {
      console.error(error); // Handle the error
    }
  };

  return (
    <div className='container'>
         <div class="row justify-content-center">

         <div class="col-md-5">
   <div class="card">
   <h2 class="card-title text-center">Login </h2>
    <form onSubmit={handleSubmit}>
        <div className='mb-3'>

      <label>
        Email:
        <input type="email" value={email} onChange={(e) => setEmail(e.target.value)} />
      </label>
        </div>
        <div className='mb-3'>

      <label>
        Password:
        <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} />
      </label>
        </div>
      <button type="submit">Login</button>
    </form>
         </div>
    </div>
         </div>
    </div>
  );
}

export default Login;
