import React, { useState } from 'react';
import axios from 'axios';

function Register() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [seller, setSeller] = useState(true);
  const [confirmPassword, setConfirmPassword] = useState('');

  const handleSubmit = async (event) => {
    event.preventDefault();
    if (password !== confirmPassword) {
      alert('Passwords do not match');
      return;
    }
    try {
      const response = await axios.post('http://localhost:8000/register', { email, password, seller});
      console.log(response.data); // Handle the API response
    } catch (error) {
      console.error(error); // Handle the error
    }
  };

  return (

    
    <div className='container'>
<div class="row justify-content-center">
  <div class="col-md-5">
   <div class="card">
    <h2>Register</h2>
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
        <div className='mb-3'>

      <label>
        Seller:
        <input type="checkbox" checked={seller} onChange={(e) => setSeller(e.target.checked)} />
      </label>
        </div>
        <div className='mb-3'>

      <label>
        Confirm Password:
        <input type="password" value={confirmPassword} onChange={(e) => setConfirmPassword(e.target.value)} />
      </label>
        </div>
      <button type="submit">Register</button>
    </form>
    </div>
    </div>
    </div>
    </div>
  );
}

export default Register;
