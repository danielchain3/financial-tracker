import React from 'react';
import { Link, Outlet } from "react-router-dom";
import './App.css';
import Navbar from './Navbar';

function App() {
  return (
    <div>
      <Navbar/>
      <Outlet/>
    </div>
  );
}

export default App;
