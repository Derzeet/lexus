import React, { useState, useEffect } from 'react';
import axios from 'axios';

function GunList() {
  const [guns, setGuns] = useState([]);

  useEffect(() => {
    // Make a GET request to retrieve all guns data
    axios.get('http://localhost:8000/store')
      .then(response => {
        console.log(response.data)
        // Update the guns state with the retrieved data
        setGuns(response.data.data);
      })
      .catch(error => {
        console.log(error);
      });
  }, []);

  const createorder = (id) => {
    axios.post("http://localhost:8000/order/" + id)
    .then(res => {
        console.log(res.data)
    })
  }

  return (
    <div className='container'>
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
    <button onClick={createorder(gun.ID)}>order</button>
  </tr>
        ))}
  </tbody>
</table>
        </div>
  );
}

export default GunList;
