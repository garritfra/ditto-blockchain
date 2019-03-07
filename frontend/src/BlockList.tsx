import React, { useState, useEffect } from "react";
import axios from "axios";

export default function BlockList() {
  let [blocks, setBlocks] = useState([]);

  const backendUrl = process.env.BACKEND_ADDRESS!;

  if (backendUrl === undefined) {
    console.error(
      "backend address missing or wrong in ENV variable: ",
      backendUrl
    );
  }

  const blocksComponent = blocks.map(block => {
    return <li key={block.hash}>{block.hash}</li>;
  });

  async function fetchBlocks() {
    let result = await axios(backendUrl);
    return result.data.Blocks;
  }

  useEffect(() => {
    fetchBlocks().then(blocks => {
      setBlocks(blocks);
    });
  }, []);

  return (
    <div>
      <h1>{blocksComponent}</h1>
    </div>
  );
}
