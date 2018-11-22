import * as net from "net";
const server = net.createServer(c => {
  //'connection' listener
  console.log("client connected");

  c.on("end", () => {
    console.log("client disconnected");
  });

  c.write("Hi!");
  c.pipe(c);
});

server.listen(42, () => {
  //'listening' listener
  console.log("server bound");
});
