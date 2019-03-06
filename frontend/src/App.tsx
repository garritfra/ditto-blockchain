import React, { useState, useEffect } from "react";
import axios from "axios";

export default function App() {
  useEffect(() => {
    axios.get("https://baconipsum.com/api/?type=meat-and-filler").then(res => {
      console.log(res.data);
    });
  });
  return (
    <div>
      <h1>Hello from App!</h1>
    </div>
  );
}
