import React, { useState, useEffect } from "react";
import axios from "axios";

export default function BlockList() {
  let [blocks, setBlocks] = useState([]);

  let backendUrl: string =
    process.env.BACKEND_ADDRESS || "http://localhost:42000/";

  if (backendUrl === undefined) {
    throw Error("backend address not specified in ENV variable");
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
