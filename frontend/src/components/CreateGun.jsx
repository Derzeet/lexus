import React, { useState } from 'react';
import axios from 'axios';

const CreateGun = () => {
  const [name, setName] = useState('');
  const [model, setModel] = useState('');
  const [caliber, setCaliber] = useState('');
  const [price, setPrice] = useState(0);
  const [availability, setAvailability] = useState(false);
  const [type, setType] = useState('');
  const token = localStorage.getItem('token'); // get token from local storage

const config = {
  headers: { Authorization: `Bearer ${token}` }
};


  const handleSubmit = async (event) => {
    event.preventDefault();
    let body = {
        "name": name,
        "caliber": caliber,
        "model": model,
        "price": price,
        "availability": availability,
        "type": type
      }
      console.log(body)
    try {
      const response = await axios.post(
        'http://localhost:8000/gun',{
            name: name,
            model: model,
            caliber: caliber,
            price: parseInt(price),
            availability: availability,
            type: type,
          }, config
      );
      console.log(response.data);
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <label>
        Name:
        <input type="text" value={name} onChange={(e) => setName(e.target.value)} />
      </label>
      <label>
        Model:
        <input type="text" value={model} onChange={(e) => setModel(e.target.value)} />
      </label>
      <label>
        Caliber:
        <input type="text" value={caliber} onChange={(e) => setCaliber(e.target.value)} />
      </label>
      <label>
        Price:
        <input type="number" value={price} onChange={(e) => setPrice(e.target.value)} />
      </label>
      <label>
        Availability:
        <input type="checkbox" checked={availability} onChange={(e) => setAvailability(e.target.checked)} />
      </label>
      <label>
        Type:
        <input type="text" value={type} onChange={(e) => setType(e.target.value)} />
      </label>
      <button type="submit">Create Gun</button>
    </form>
  );
};

export default CreateGun;
