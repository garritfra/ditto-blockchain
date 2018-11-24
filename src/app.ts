import * as net from "net";
const server = net.createServer(c => {
  //'connection' listener
  console.log("client connected");

  c.on("end", () => {
    console.log("client disconnected");
  });

  c.on("data", data => {
    try {
      console.log(JSON.parse(data.toString()).message);
    } catch (e) {
      console.log("Request was not a JSON Object!");
    }
  });

  c.write("Hi!");
  c.pipe(c);
});

server.listen(42000, () => {
  //'listening' listener
  console.log("server bound");
});
