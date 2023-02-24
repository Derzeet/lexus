import React, { useState, useEffect } from 'react';
import axios from 'axios';

const ProfilePage = () => {
  const [account, setAccount] = useState(null);
  const [guns, setGuns] = useState(null);

  const token = localStorage.getItem('token'); // get token from local storage

const config = {
  headers: { Authorization: `Bearer ${token}` }
};

  useEffect(() => {
    const fetchAccount = async () => {
      try {
        const token = localStorage.getItem('token');
        console.log(config)
        axios.defaults.headers.common['Authorization'] = token;
        const response = await axios.get('http://localhost:8000/profile', config)
        console.log(response.data.data)
        setGuns(response.data.data);
        const res = await axios.get('http://localhost:8000/user/'+response.data.data[0].user_id, config)
        setAccount(res.data.user)
      } catch (error) {
        console.error(error);
      }
    };
    fetchAccount();
  }, []);


  return (
    <div className='container'>
      {account && (
        <>
          <h1>Profile Page</h1>
          <p>Email: {account.email}</p>
          <p>Created At: {account.CreatedAt}</p>
          <p>Seller: {account.Seller ? 'Yes' : 'No'}</p>
        </>
      )}
      
      <h2>MY PRODUCTS</h2>
      {guns && (
        <>
          <table class="table">
  <thead>
    <tr>
      <th scope="col">#</th>
      <th scope="col">Name</th>
      <th scope="col">Caliber</th>
      <th scope="col">Model</th>
      <th scope="col">Price</th>
      <th scope="col">Type</th>
    </tr>
  </thead>
  <tbody>
  {guns.map((gun, index) => (
    <tr>
    <th scope="row">{index + 1}</th>
    <td>{gun.name}</td>
    <td>{gun.caliber}</td>
    <td>{gun.model}</td>
    <td>{gun.price}</td>
    <td>{gun.type}</td>
  </tr>
        ))}
  </tbody>
</table>
        </>
      )}
    </div>
  );
};

export default ProfilePage;
